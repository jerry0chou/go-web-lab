package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"go-web-lab/redis/cache_patterns/models"

	"github.com/redis/go-redis/v9"
)

// CacheService handles caching operations with various protection mechanisms
type CacheService struct {
	rdb     *redis.Client
	mu      sync.Mutex          // For cache breakdown protection
	locks   map[int]*sync.Mutex // Per-key locks
	muLocks sync.Mutex          // Protects the locks map
}

// NewCacheService creates a new cache service
func NewCacheService(rdb *redis.Client) *CacheService {
	return &CacheService{
		rdb:   rdb,
		locks: make(map[int]*sync.Mutex),
	}
}

// GetLock gets or creates a lock for a specific key
func (c *CacheService) GetLock(key int) *sync.Mutex {
	c.muLocks.Lock()
	defer c.muLocks.Unlock()

	if lock, exists := c.locks[key]; exists {
		return lock
	}

	lock := &sync.Mutex{}
	c.locks[key] = lock
	return lock
}

// GetProduct implements Cache-Aside pattern with protection mechanisms
func (c *CacheService) GetProduct(ctx context.Context, id int, dbGetter func(int) (*models.Product, error)) (*models.Product, error) {
	cacheKey := fmt.Sprintf("product:%d", id)

	// 1. Try to get from cache first
	val, err := c.rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		// Cache hit
		var product models.Product
		if err := json.Unmarshal([]byte(val), &product); err == nil {
			return &product, nil
		}
	}

	// Cache miss - check if it's a null cache (防止缓存穿透)
	if err == redis.Nil {
		nullKey := fmt.Sprintf("product:null:%d", id)
		exists, _ := c.rdb.Exists(ctx, nullKey).Result()
		if exists > 0 {
			// This is a known non-existent product, return nil
			return nil, fmt.Errorf("product not found: %d", id)
		}
	}

	// 2. Cache miss - get lock to prevent cache breakdown (防止缓存击穿)
	lock := c.GetLock(id)
	lock.Lock()
	defer lock.Unlock()

	// Double-check: another goroutine might have populated the cache
	val, err = c.rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		var product models.Product
		if err := json.Unmarshal([]byte(val), &product); err == nil {
			return &product, nil
		}
	}

	// 3. Get from database
	product, err := dbGetter(id)
	if err != nil {
		// Product not found - cache null value to prevent cache penetration (防止缓存穿透)
		nullKey := fmt.Sprintf("product:null:%d", id)
		c.rdb.Set(ctx, nullKey, "1", 5*time.Minute) // Short TTL for null cache
		return nil, err
	}

	// 4. Write to cache with random TTL to prevent cache avalanche (防止缓存雪崩)
	productJSON, _ := json.Marshal(product)

	// Random TTL: base 30 minutes + random 0-10 minutes
	baseTTL := 30 * time.Minute
	randomTTL := time.Duration(rand.Intn(600)) * time.Second // 0-10 minutes
	ttl := baseTTL + randomTTL

	c.rdb.Set(ctx, cacheKey, productJSON, ttl)

	return product, nil
}

// InvalidateProduct removes a product from cache
func (c *CacheService) InvalidateProduct(ctx context.Context, id int) error {
	cacheKey := fmt.Sprintf("product:%d", id)
	nullKey := fmt.Sprintf("product:null:%d", id)

	pipe := c.rdb.Pipeline()
	pipe.Del(ctx, cacheKey)
	pipe.Del(ctx, nullKey)
	_, err := pipe.Exec(ctx)

	return err
}

// GetProductWithStats returns product and cache statistics
func (c *CacheService) GetProductWithStats(ctx context.Context, id int, dbGetter func(int) (*models.Product, error)) (*models.Product, bool, error) {
	cacheKey := fmt.Sprintf("product:%d", id)

	// Check cache first
	val, err := c.rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		var product models.Product
		if err := json.Unmarshal([]byte(val), &product); err == nil {
			return &product, true, nil // true = cache hit
		}
	}

	// Cache miss - proceed with normal flow
	product, err := c.GetProduct(ctx, id, dbGetter)
	return product, false, err // false = cache miss
}
