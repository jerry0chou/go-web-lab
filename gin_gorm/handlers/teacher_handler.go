package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-web-lab/gin_gorm/database"
	"go-web-lab/gin_gorm/models"
)

func GetTeachers(c *gin.Context) {
	var teachers []models.Teacher
	database.DB.Find(&teachers)
	c.JSON(http.StatusOK, teachers)
}

func GetTeacher(c *gin.Context) {
	id := c.Param("id")
	var teacher models.Teacher

	if err := database.DB.Preload("Courses").First(&teacher, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}

	c.JSON(http.StatusOK, teacher)
}

func CreateTeacher(c *gin.Context) {
	var teacher models.Teacher

	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&teacher)
	c.JSON(http.StatusCreated, teacher)
}

func UpdateTeacher(c *gin.Context) {
	id := c.Param("id")
	var teacher models.Teacher

	if err := database.DB.First(&teacher, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}

	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&teacher)
	c.JSON(http.StatusOK, teacher)
}

func DeleteTeacher(c *gin.Context) {
	id := c.Param("id")
	var teacher models.Teacher

	if err := database.DB.First(&teacher, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}

	database.DB.Delete(&teacher)
	c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted successfully"})
}
