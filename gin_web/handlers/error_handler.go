package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TriggerError(c *gin.Context) {
	c.Error(errors.New("sample error"))
	c.Error(errors.New("another error"))

	err := c.Errors.Last()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "No errors"})
}

func AllErrors(c *gin.Context) {
	c.Error(errors.New("error 1"))
	c.Error(errors.New("error 2"))
	c.Error(errors.New("error 3"))

	var errorMessages []string
	for _, err := range c.Errors {
		errorMessages = append(errorMessages, err.Error())
	}
	c.JSON(http.StatusInternalServerError, gin.H{"errors": errorMessages})
}

func CustomError(c *gin.Context) {
	type CustomErr struct {
		Code    int
		Message string
	}

	err := CustomErr{
		Code:    404,
		Message: "Resource not found",
	}
	c.Error(errors.New(err.Message)).SetType(gin.ErrorTypePublic)
	c.JSON(http.StatusNotFound, gin.H{"error": err.Message})
}
