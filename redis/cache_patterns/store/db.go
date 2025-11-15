package store

import (
	"fmt"
	"go-web-lab/redis/cache_patterns/models"
	"sync"
	"time"
)

// MockDB simulates a database
type MockDB struct {
	products map[int]*models.Product
	mu       sync.RWMutex
}

// NewMockDB creates a new mock database
func NewMockDB() *MockDB {
	db := &MockDB{
		products: make(map[int]*models.Product),
	}

	// Initialize with some sample products
	db.initData()
	return db
}

// initData initializes sample product data
func (db *MockDB) initData() {
	db.mu.Lock()
	defer db.mu.Unlock()

	products := []*models.Product{
		{ID: 1, Name: "Laptop", Price: 999.99, Stock: 10},
		{ID: 2, Name: "Mouse", Price: 29.99, Stock: 50},
		{ID: 3, Name: "Keyboard", Price: 79.99, Stock: 30},
		{ID: 4, Name: "Monitor", Price: 299.99, Stock: 15},
		{ID: 5, Name: "Headphones", Price: 149.99, Stock: 25},
	}

	for _, p := range products {
		db.products[p.ID] = p
	}
}

// GetProduct simulates a database query with latency
func (db *MockDB) GetProduct(id int) (*models.Product, error) {
	// Simulate database query latency (50-200ms)
	time.Sleep(time.Duration(50+id*10) * time.Millisecond)

	db.mu.RLock()
	defer db.mu.RUnlock()

	product, exists := db.products[id]
	if !exists {
		return nil, fmt.Errorf("product not found: %d", id)
	}

	return product, nil
}

// GetAllProducts returns all products
func (db *MockDB) GetAllProducts() ([]*models.Product, error) {
	// Simulate database query latency
	time.Sleep(100 * time.Millisecond)

	db.mu.RLock()
	defer db.mu.RUnlock()

	products := make([]*models.Product, 0, len(db.products))
	for _, p := range db.products {
		products = append(products, p)
	}

	return products, nil
}
