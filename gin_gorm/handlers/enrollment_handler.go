package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-web-lab/gin_gorm/database"
	"go-web-lab/gin_gorm/models"
)

func GetEnrollments(c *gin.Context) {
	var enrollments []models.Enrollment
	database.DB.Preload("Student").Preload("Course").Find(&enrollments)
	c.JSON(http.StatusOK, enrollments)
}

func GetEnrollment(c *gin.Context) {
	id := c.Param("id")
	var enrollment models.Enrollment

	if err := database.DB.Preload("Student").Preload("Course").First(&enrollment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Enrollment not found"})
		return
	}

	c.JSON(http.StatusOK, enrollment)
}

func CreateEnrollment(c *gin.Context) {
	var enrollment models.Enrollment

	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&enrollment)
	database.DB.Preload("Student").Preload("Course").First(&enrollment, enrollment.ID)
	c.JSON(http.StatusCreated, enrollment)
}

func UpdateEnrollment(c *gin.Context) {
	id := c.Param("id")
	var enrollment models.Enrollment

	if err := database.DB.First(&enrollment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Enrollment not found"})
		return
	}

	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&enrollment)
	database.DB.Preload("Student").Preload("Course").First(&enrollment, enrollment.ID)
	c.JSON(http.StatusOK, enrollment)
}

func DeleteEnrollment(c *gin.Context) {
	id := c.Param("id")
	var enrollment models.Enrollment

	if err := database.DB.First(&enrollment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Enrollment not found"})
		return
	}

	database.DB.Delete(&enrollment)
	c.JSON(http.StatusOK, gin.H{"message": "Enrollment deleted successfully"})
}
