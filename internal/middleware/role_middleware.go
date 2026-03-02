package middleware

import "github.com/gofiber/fiber/v2"

func Authorize(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		userRole := c.Locals("role")

		if userRole == nil {
			return c.Status(fiber.StatusForbidden).
				JSON(fiber.Map{"error": "role not found"})
		}

		for _, role := range allowedRoles {
			if role == userRole {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).
			JSON(fiber.Map{"error": "access denied"})
	}
}
