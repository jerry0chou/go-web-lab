package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetCookie(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "session",
		Value:    "abc123",
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
	})
	return c.JSON(fiber.Map{"message": "Cookie set"})
}

func GetCookie(c *fiber.Ctx) error {
	cookie := c.Cookies("session")
	if cookie == "" {
		return c.JSON(fiber.Map{"cookie": "not found"})
	}
	return c.JSON(fiber.Map{"cookie": cookie})
}

func DeleteCookie(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "session",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{"message": "Cookie deleted"})
}

func MultipleCookies(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "user",
		Value:    "alice",
		MaxAge:   3600,
		HTTPOnly: true,
	})
	c.Cookie(&fiber.Cookie{
		Name:     "theme",
		Value:    "dark",
		MaxAge:   3600,
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{"message": "Multiple cookies set"})
}
