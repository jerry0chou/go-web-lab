package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func XMLResponse(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/xml")
	return c.SendString(`<?xml version="1.0"?><response><message>Hello XML</message></response>`)
}

func YAMLResponse(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/x-yaml")
	return c.SendString("message: Hello YAML\nversion: 1.0")
}

func StringResponse(c *fiber.Ctx) error {
	return c.SendString("Hello from Fiber")
}

func DataResponse(c *fiber.Ctx) error {
	data := []byte("Raw binary data")
	return c.Send(data)
}

func HeaderOnlyResponse(c *fiber.Ctx) error {
	c.Set("X-Custom-Header", "custom-value")
	return c.SendStatus(fiber.StatusNoContent)
}
