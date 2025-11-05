package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to Fiber Demo",
		"version": "1.0",
		"endpoints": []string{
			"GET /users",
			"POST /users",
			"GET /api/products",
			"GET /search?q=query",
		},
	})
}

func NoRoute(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Route not found",
	})
}
