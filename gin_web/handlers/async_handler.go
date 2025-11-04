package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AsyncHandler(c *gin.Context) {
	copyContext := c.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		log.Printf("Async task completed for path: %s", copyContext.Request.URL.Path)
	}()
	c.JSON(http.StatusOK, gin.H{"message": "Async task started"})
}

func DelayedResponse(c *gin.Context) {
	copyContext := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		copyContext.JSON(http.StatusOK, gin.H{
			"message": "Delayed response after 2 seconds",
		})
	}()
}
