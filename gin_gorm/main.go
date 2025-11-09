package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go-web-lab/gin_gorm/database"
	"go-web-lab/gin_gorm/handlers"
)

func main() {
	// Connect to database
	database.Connect()

	// Run migrations
	database.Migrate()

	// Setup router
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Student routes
	students := r.Group("/students")
	{
		students.GET("", handlers.GetStudents)
		students.GET("/:id", handlers.GetStudent)
		students.POST("", handlers.CreateStudent)
		students.PUT("/:id", handlers.UpdateStudent)
		students.DELETE("/:id", handlers.DeleteStudent)
	}

	// Teacher routes
	teachers := r.Group("/teachers")
	{
		teachers.GET("", handlers.GetTeachers)
		teachers.GET("/:id", handlers.GetTeacher)
		teachers.POST("", handlers.CreateTeacher)
		teachers.PUT("/:id", handlers.UpdateTeacher)
		teachers.DELETE("/:id", handlers.DeleteTeacher)
	}

	// Course routes
	courses := r.Group("/courses")
	{
		courses.GET("", handlers.GetCourses)
		courses.GET("/:id", handlers.GetCourse)
		courses.POST("", handlers.CreateCourse)
		courses.PUT("/:id", handlers.UpdateCourse)
		courses.DELETE("/:id", handlers.DeleteCourse)
	}

	// Enrollment routes
	enrollments := r.Group("/enrollments")
	{
		enrollments.GET("", handlers.GetEnrollments)
		enrollments.GET("/:id", handlers.GetEnrollment)
		enrollments.POST("", handlers.CreateEnrollment)
		enrollments.PUT("/:id", handlers.UpdateEnrollment)
		enrollments.DELETE("/:id", handlers.DeleteEnrollment)
	}

	// Start server
	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
