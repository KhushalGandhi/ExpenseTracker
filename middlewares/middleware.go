package middlewares

import (
	"expensetracker/utils"
	"github.com/gofiber/fiber/v2"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
		}

		claims, err := utils.ParseJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		c.Locals("user_id", claims["user_id"])
		return c.Next()
	}
}
