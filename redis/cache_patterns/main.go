package main

import (
	"log"

	"go-web-lab/redis/cache_patterns/cache"
	"go-web-lab/redis/cache_patterns/database"
	"go-web-lab/redis/cache_patterns/handlers"
	"go-web-lab/redis/cache_patterns/store"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to Redis
	rdb, err := database.ConnectRedis()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("✅ Successfully connected to Redis!")

	// Initialize services
	db := store.NewMockDB()
	cacheService := cache.NewCacheService(rdb)
	productHandler := handlers.NewProductHandler(cacheService, db)

	// Setup router
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Product routes
	products := r.Group("/products")
	{
		products.GET("", productHandler.GetAllProducts)
		products.GET("/:id", productHandler.GetProduct)
		products.DELETE("/:id/cache", productHandler.InvalidateCache)
	}

	// Start server
	port := ":8080"
	log.Printf("🚀 Server starting on %s", port)
	log.Println("📖 API Endpoints:")
	log.Println("   GET    /products          - List all products")
	log.Println("   GET    /products/:id      - Get product by ID (with cache)")
	log.Println("   DELETE /products/:id/cache - Invalidate product cache")
	log.Println("")
	log.Println("💡 Try: curl http://localhost:8080/products/1")

	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
