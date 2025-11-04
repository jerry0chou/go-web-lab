package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Gin Demo",
		"name":  "World",
	})
}

func RenderUserTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "user.html", gin.H{
		"userID": c.Param("id"),
		"name":   c.Query("name"),
	})
}

func RenderWithLayout(c *gin.Context) {
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"content": "This is dynamic content",
		"items":   []string{"Item 1", "Item 2", "Item 3"},
	})
}
