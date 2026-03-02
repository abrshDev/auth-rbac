package auth

import (
	"github.com/abrshDev/auth-rbac/internal/user"

	"github.com/abrshDev/auth-rbac/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	repo *user.Repository
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"` // optional
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthHandler(repo *user.Repository) *AuthHandler {
	return &AuthHandler{repo: repo}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var body RegisterRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Use repository to create user with hashed password
	user, err := h.repo.CreateUserWithPassword(body.Username, body.Email, body.Password, body.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Return safe response (without password)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"role":     user.Role,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var body LoginRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.repo.GetUserByEmail(body.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}

	if !utils.CheckPassword(user.PasswordHash, body.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}

	// Generate JWT
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "could not generate token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
