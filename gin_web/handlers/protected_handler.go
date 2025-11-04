package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProtectedData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is protected data"})
}

func ProtectedWithUser(c *gin.Context) {
	userID, _ := c.Get("userID")
	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{
		"message":  "Protected endpoint with user context",
		"userID":   userID,
		"username": username,
	})
}
