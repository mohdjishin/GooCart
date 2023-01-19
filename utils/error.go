package utils

import "github.com/gofiber/fiber/v2"

func InternalServerError(msg string, c *fiber.Ctx) error {
	return c.Status(500).JSON(fiber.Map{
		"message": msg,
	})
}
