package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go-web-lab/fiber_gorm/database"
	"go-web-lab/fiber_gorm/models"
)

func GetStudents(c *fiber.Ctx) error {
	var students []models.Student
	database.DB.Find(&students)
	return c.JSON(students)
}

func GetStudent(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	var student models.Student

	if err := database.DB.Preload("Enrollments.Course").First(&student, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Student not found"})
	}

	return c.JSON(student)
}

func CreateStudent(c *fiber.Ctx) error {
	var student models.Student

	if err := c.BodyParser(&student); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	database.DB.Create(&student)
	return c.Status(fiber.StatusCreated).JSON(student)
}

func UpdateStudent(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	var student models.Student

	if err := database.DB.First(&student, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Student not found"})
	}

	if err := c.BodyParser(&student); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	database.DB.Save(&student)
	return c.JSON(student)
}

func DeleteStudent(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	var student models.Student

	if err := database.DB.First(&student, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Student not found"})
	}

	database.DB.Delete(&student)
	return c.JSON(fiber.Map{"message": "Student deleted successfully"})
}
