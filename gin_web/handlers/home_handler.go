package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Gin Demo",
		"version": "1.0",
		"endpoints": []string{
			"GET /users",
			"POST /users",
			"GET /api/products",
			"GET /search?q=query",
		},
	})
}

func NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
}
