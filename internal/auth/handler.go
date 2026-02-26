package auth

import "github.com/abrshDev/auth-rbac/internal/user"

type AuthHandler struct {
	repo user.Repository
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
