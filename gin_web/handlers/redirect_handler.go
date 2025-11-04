package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedirectToURL(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://github.com/gin-gonic/gin")
}

func RedirectInternal(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}

func RedirectWithStatus(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/api/products")
}

func ConditionalRedirect(c *gin.Context) {
	redirect := c.Query("redirect")
	if redirect == "true" {
		c.Redirect(http.StatusFound, "/")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "No redirect requested"})
}
