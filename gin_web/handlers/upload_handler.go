package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadSingleFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"filename": file.Filename,
		"size":     file.Size,
		"header":   file.Header,
	})
}

func UploadMultipleFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	files := form.File["files"]

	var fileInfo []gin.H
	for _, file := range files {
		fileInfo = append(fileInfo, gin.H{
			"filename": file.Filename,
			"size":     file.Size,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"count": len(fileInfo),
		"files": fileInfo,
	})
}

func ProcessForm(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	age := c.DefaultPostForm("age", "0")

	c.JSON(http.StatusOK, gin.H{
		"name":  name,
		"email": email,
		"age":   age,
	})
}

func SaveUploadedFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filename := filepath.Base(file.Filename)
	dst := "./gin_web/uploads/" + filename

	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "File uploaded successfully",
		"filename": filename,
		"path":     dst,
	})
}
