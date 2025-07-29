package middleware

import (
	"strings"

	"github.com/iskhakmuhamad/todo-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	jwtSecret string
}

func NewAuthMiddleware(jwtSecret string) *AuthMiddleware {
	return &AuthMiddleware{
		jwtSecret: jwtSecret,
	}
}

func (m *AuthMiddleware) ValidateJWT(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header required",
		})
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	claims, err := utils.ValidateJWT(tokenString, m.jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Store user info in context
	c.Locals("userID", claims.UserID)
	c.Locals("email", claims.Email)

	return c.Next()
}
