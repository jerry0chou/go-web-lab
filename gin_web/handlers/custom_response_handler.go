package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func XMLResponse(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{
		"message": "Hello XML",
		"status":  "success",
	})
}

func YAMLResponse(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{
		"message": "Hello YAML",
		"status":  "success",
	})
}

func ProtoBufResponse(c *gin.Context) {
	c.ProtoBuf(http.StatusOK, gin.H{
		"message": "Hello Protobuf",
		"status":  "success",
	})
}

func StringResponse(c *gin.Context) {
	c.String(http.StatusOK, "Hello, this is a plain text response")
}

func DataResponse(c *gin.Context) {
	data := []byte("This is raw binary data")
	c.Data(http.StatusOK, "application/octet-stream", data)
}

func HeaderOnlyResponse(c *gin.Context) {
	c.Header("Custom-Header", "CustomValue")
	c.Status(http.StatusNoContent)
}
