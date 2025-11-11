package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go-web-lab/fiber_gorm/database"
	"go-web-lab/fiber_gorm/models"
)

func GetTeachers(c *fiber.Ctx) error {
	var teachers []models.Teacher
	database.DB.Find(&teachers)
	return c.JSON(teachers)
}

func GetTeacher(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	var teacher models.Teacher

	if err := database.DB.Preload("Courses").First(&teacher, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Teacher not found"})
	}

	return c.JSON(teacher)
}

func CreateTeacher(c *fiber.Ctx) error {
	var teacher models.Teacher

	if err := c.BodyParser(&teacher); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	database.DB.Create(&teacher)
	return c.Status(fiber.StatusCreated).JSON(teacher)
}

func UpdateTeacher(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	var teacher models.Teacher

	if err := database.DB.First(&teacher, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Teacher not found"})
	}

	if err := c.BodyParser(&teacher); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	database.DB.Save(&teacher)
	return c.JSON(teacher)
}

func DeleteTeacher(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	var teacher models.Teacher

	if err := database.DB.First(&teacher, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Teacher not found"})
	}

	database.DB.Delete(&teacher)
	return c.JSON(fiber.Map{"message": "Teacher deleted successfully"})
}
