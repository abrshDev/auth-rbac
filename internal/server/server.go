package server

import (
	"github.com/abrshDev/auth-rbac/internal/auth"
	"github.com/abrshDev/auth-rbac/internal/middleware"
	"github.com/abrshDev/auth-rbac/internal/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

func registeroutes(app *fiber.App, db *gorm.DB) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Auth RBAC Fiber API is running!")
	})
	userRepo := user.NewRepository(DB)

	authHandler := auth.NewAuthHandler(userRepo)

	// Group routes
	api := app.Group("/api")
	authGroup := api.Group("/auth")

	// Register endpoint
	authGroup.Post("/register", authHandler.Register)
	authGroup.Post("/login", authHandler.Login)
	api.Get("/admin/dashboard",
		middleware.Protected(),
		middleware.Authorize("admin"),
		func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"message": "Welcome Admin",
			})
		},
	)
	api.Get("/profile",
		middleware.Protected(),
		func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"user_id": c.Locals("user_id"),
				"role":    c.Locals("role"),
			})
		},
	)

}

func NewApp(db *gorm.DB) *fiber.App {
	app := fiber.New()
	registeroutes(app, db)
	return app
}
