package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func ProtectedData(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "This is protected data",
		"data":    "sensitive information",
	})
}

func ProtectedWithUser(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	return c.JSON(fiber.Map{
		"message": "Protected user data",
		"token":   token,
	})
}
