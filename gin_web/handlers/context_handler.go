package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetContext(c *gin.Context) {
	c.Set("userID", 123)
	c.Set("username", "alice")
	c.JSON(http.StatusOK, gin.H{"message": "Context values set"})
}

func GetContext(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusOK, gin.H{"userID": "not set"})
		return
	}
	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{
		"userID":   userID,
		"username": username,
	})
}

func ContextChain(c *gin.Context) {
	c.Set("step1", "completed")
	c.Set("step2", "completed")
	c.JSON(http.StatusOK, gin.H{
		"step1": c.MustGet("step1"),
		"step2": c.MustGet("step2"),
	})
}
