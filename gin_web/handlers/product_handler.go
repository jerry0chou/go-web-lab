package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-web-lab/gin_web/models"
	"go-web-lab/gin_web/store"
)

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"products": store.Products})
}

func CreateProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newProduct.ID = len(store.Products) + 1
	store.Products = append(store.Products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}
