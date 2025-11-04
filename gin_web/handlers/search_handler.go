package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	query := c.Query("q")
	limit := c.DefaultQuery("limit", "10")
	category := c.Query("category")

	c.JSON(http.StatusOK, gin.H{
		"query":    query,
		"limit":    limit,
		"category": category,
	})
}
