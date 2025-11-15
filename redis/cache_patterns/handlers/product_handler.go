package handlers

import (
	"net/http"
	"strconv"

	"go-web-lab/redis/cache_patterns/cache"
	"go-web-lab/redis/cache_patterns/store"

	"github.com/gin-gonic/gin"
)

// ProductHandler handles product-related HTTP requests
type ProductHandler struct {
	cacheService *cache.CacheService
	db           *store.MockDB
}

// NewProductHandler creates a new product handler
func NewProductHandler(cacheService *cache.CacheService, db *store.MockDB) *ProductHandler {
	return &ProductHandler{
		cacheService: cacheService,
		db:           db,
	}
}

// GetProduct handles GET /products/:id
func (h *ProductHandler) GetProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	// Get product with cache statistics
	product, fromCache, err := h.cacheService.GetProductWithStats(
		c.Request.Context(),
		id,
		h.db.GetProduct,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":      err.Error(),
			"from_cache": fromCache,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product":    product,
		"from_cache": fromCache,
		"message":    getCacheMessage(fromCache),
	})
}

// GetAllProducts handles GET /products
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.db.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
		"count":    len(products),
	})
}

// InvalidateCache handles DELETE /products/:id/cache
func (h *ProductHandler) InvalidateCache(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	err = h.cacheService.InvalidateProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "cache invalidated",
		"product_id": id,
	})
}

// getCacheMessage returns a descriptive message based on cache hit/miss
func getCacheMessage(fromCache bool) string {
	if fromCache {
		return "✅ Data retrieved from Redis cache (fast!)"
	}
	return "⏳ Data retrieved from database and cached for next time"
}
