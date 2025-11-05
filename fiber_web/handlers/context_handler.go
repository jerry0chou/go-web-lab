package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func SetContext(c *fiber.Ctx) error {
	c.Locals("userID", 123)
	c.Locals("username", "alice")
	return c.JSON(fiber.Map{"message": "Context values set"})
}

func GetContext(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	if userID == nil {
		return c.JSON(fiber.Map{"userID": "not set"})
	}
	username := c.Locals("username")
	return c.JSON(fiber.Map{
		"userID":   userID,
		"username": username,
	})
}

func ContextChain(c *fiber.Ctx) error {
	c.Locals("step1", "completed")
	c.Locals("step2", "completed")
	return c.JSON(fiber.Map{
		"step1": c.Locals("step1"),
		"step2": c.Locals("step2"),
	})
}
