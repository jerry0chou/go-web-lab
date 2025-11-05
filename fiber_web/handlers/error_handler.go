package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func TriggerError(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusInternalServerError, "Something went wrong")
}

func CustomError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error":   "Custom error occurred",
		"code":    400,
		"message": "This is a custom error response",
	})
}
