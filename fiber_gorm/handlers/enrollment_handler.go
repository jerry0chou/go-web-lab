package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go-web-lab/fiber_gorm/database"
	"go-web-lab/fiber_gorm/models"
)

func GetEnrollments(c *fiber.Ctx) error {
	var enrollments []models.Enrollment
	database.DB.Preload("Student").Preload("Course").Find(&enrollments)
	return c.JSON(enrollments)
}

func GetEnrollment(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	var enrollment models.Enrollment

	if err := database.DB.Preload("Student").Preload("Course").First(&enrollment, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Enrollment not found"})
	}

	return c.JSON(enrollment)
}

func CreateEnrollment(c *fiber.Ctx) error {
	var enrollment models.Enrollment

	if err := c.BodyParser(&enrollment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	database.DB.Create(&enrollment)
	database.DB.Preload("Student").Preload("Course").First(&enrollment, enrollment.ID)
	return c.Status(fiber.StatusCreated).JSON(enrollment)
}

func UpdateEnrollment(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	var enrollment models.Enrollment

	if err := database.DB.First(&enrollment, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Enrollment not found"})
	}

	if err := c.BodyParser(&enrollment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	database.DB.Save(&enrollment)
	database.DB.Preload("Student").Preload("Course").First(&enrollment, enrollment.ID)
	return c.JSON(enrollment)
}

func DeleteEnrollment(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	var enrollment models.Enrollment

	if err := database.DB.First(&enrollment, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Enrollment not found"})
	}

	database.DB.Delete(&enrollment)
	return c.JSON(fiber.Map{"message": "Enrollment deleted successfully"})
}
