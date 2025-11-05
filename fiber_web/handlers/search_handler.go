package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Search(c *fiber.Ctx) error {
	query := c.Query("q")
	return c.JSON(fiber.Map{
		"query":   query,
		"results": []string{"result1", "result2", "result3"},
	})
}
