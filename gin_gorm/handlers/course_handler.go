package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-web-lab/gin_gorm/database"
	"go-web-lab/gin_gorm/models"
)

func GetCourses(c *gin.Context) {
	var courses []models.Course
	database.DB.Preload("Teacher").Find(&courses)
	c.JSON(http.StatusOK, courses)
}

func GetCourse(c *gin.Context) {
	id := c.Param("id")
	var course models.Course

	if err := database.DB.Preload("Teacher").Preload("Enrollments.Student").First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, course)
}

func CreateCourse(c *gin.Context) {
	var course models.Course

	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&course)
	database.DB.Preload("Teacher").First(&course, course.ID)
	c.JSON(http.StatusCreated, course)
}

func UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	var course models.Course

	if err := database.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&course)
	database.DB.Preload("Teacher").First(&course, course.ID)
	c.JSON(http.StatusOK, course)
}

func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	var course models.Course

	if err := database.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	database.DB.Delete(&course)
	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}
