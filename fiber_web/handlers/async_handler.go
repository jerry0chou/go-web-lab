package handlers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AsyncHandler(c *fiber.Ctx) error {
	path := c.Path()
	go func() {
		time.Sleep(5 * time.Second)
		log.Printf("Async task completed for path: %s", path)
	}()
	return c.JSON(fiber.Map{"message": "Async task started"})
}

func DelayedResponse(c *fiber.Ctx) error {
	go func() {
		time.Sleep(2 * time.Second)
		log.Printf("Delayed response after 2 seconds")
	}()
	return c.JSON(fiber.Map{
		"message": "Delayed response started",
	})
}
