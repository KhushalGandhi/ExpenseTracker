package utils

import "github.com/gofiber/fiber/v2"

func ErrorResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": message,
	})
}

func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(data)
}
