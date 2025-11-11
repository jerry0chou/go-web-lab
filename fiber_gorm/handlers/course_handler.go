package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go-web-lab/fiber_gorm/database"
	"go-web-lab/fiber_gorm/models"
)

func GetCourses(c *fiber.Ctx) error {
	var courses []models.Course
	database.DB.Preload("Teacher").Find(&courses)
	return c.JSON(courses)
}

func GetCourse(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	var course models.Course

	if err := database.DB.Preload("Teacher").Preload("Enrollments.Student").First(&course, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Course not found"})
	}

	return c.JSON(course)
}

func CreateCourse(c *fiber.Ctx) error {
	var course models.Course

	if err := c.BodyParser(&course); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	database.DB.Create(&course)
	database.DB.Preload("Teacher").First(&course, course.ID)
	return c.Status(fiber.StatusCreated).JSON(course)
}

func UpdateCourse(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	var course models.Course

	if err := database.DB.First(&course, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Course not found"})
	}

	if err := c.BodyParser(&course); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	database.DB.Save(&course)
	database.DB.Preload("Teacher").First(&course, course.ID)
	return c.JSON(course)
}

func DeleteCourse(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	var course models.Course

	if err := database.DB.First(&course, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Course not found"})
	}

	database.DB.Delete(&course)
	return c.JSON(fiber.Map{"message": "Course deleted successfully"})
}
