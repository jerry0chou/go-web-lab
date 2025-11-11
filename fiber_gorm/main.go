package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go-web-lab/fiber_gorm/database"
	"go-web-lab/fiber_gorm/handlers"
)

func main() {
	database.Connect()
	database.Migrate()

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	students := app.Group("/students")
	{
		students.Get("", handlers.GetStudents)
		students.Get("/:id", handlers.GetStudent)
		students.Post("", handlers.CreateStudent)
		students.Put("/:id", handlers.UpdateStudent)
		students.Delete("/:id", handlers.DeleteStudent)
	}

	teachers := app.Group("/teachers")
	{
		teachers.Get("", handlers.GetTeachers)
		teachers.Get("/:id", handlers.GetTeacher)
		teachers.Post("", handlers.CreateTeacher)
		teachers.Put("/:id", handlers.UpdateTeacher)
		teachers.Delete("/:id", handlers.DeleteTeacher)
	}

	courses := app.Group("/courses")
	{
		courses.Get("", handlers.GetCourses)
		courses.Get("/:id", handlers.GetCourse)
		courses.Post("", handlers.CreateCourse)
		courses.Put("/:id", handlers.UpdateCourse)
		courses.Delete("/:id", handlers.DeleteCourse)
	}

	enrollments := app.Group("/enrollments")
	{
		enrollments.Get("", handlers.GetEnrollments)
		enrollments.Get("/:id", handlers.GetEnrollment)
		enrollments.Post("", handlers.CreateEnrollment)
		enrollments.Put("/:id", handlers.UpdateEnrollment)
		enrollments.Delete("/:id", handlers.DeleteEnrollment)
	}

	log.Println("Server starting on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
