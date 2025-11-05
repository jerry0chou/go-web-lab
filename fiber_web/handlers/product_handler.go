package handlers

import (
	"go-web-lab/fiber_web/models"
	"go-web-lab/fiber_web/store"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"products": store.Products})
}

func CreateProduct(c *fiber.Ctx) error {
	var newProduct models.Product
	if err := c.BodyParser(&newProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	store.Products = append(store.Products, newProduct)
	return c.Status(fiber.StatusCreated).JSON(newProduct)
}
