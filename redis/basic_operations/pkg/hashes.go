package redisops

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func DemonstrateHashes(ctx context.Context, rdb *redis.Client) {
	fmt.Println("\n--- HASH Operations ---")

	// HSET - Set field in hash
	err := rdb.HSet(ctx, "user:1000", map[string]interface{}{
		"name":  "Alice",
		"email": "alice@example.com",
		"age":   "30",
	}).Err()
	if err != nil {
		log.Printf("Error setting hash: %v", err)
		return
	}
	fmt.Println("✓ HSET user:1000 name Alice email alice@example.com age 30")

	// HGET - Get field from hash
	name, err := rdb.HGet(ctx, "user:1000", "name").Result()
	if err != nil {
		log.Printf("Error getting hash field: %v", err)
	} else {
		fmt.Printf("✓ HGET user:1000 name = %s\n", name)
	}

	// HGETALL - Get all fields from hash
	user, err := rdb.HGetAll(ctx, "user:1000").Result()
	if err != nil {
		log.Printf("Error getting all hash fields: %v", err)
	} else {
		fmt.Printf("✓ HGETALL user:1000 = %v\n", user)
	}

	// HINCRBY - Increment numeric field
	age, err := rdb.HIncrBy(ctx, "user:1000", "age", 1).Result()
	if err != nil {
		log.Printf("Error incrementing hash field: %v", err)
	} else {
		fmt.Printf("✓ HINCRBY user:1000 age 1 = %d\n", age)
	}

	// HDEL - Delete field from hash
	err = rdb.HDel(ctx, "user:1000", "age").Err()
	if err != nil {
		log.Printf("Error deleting hash field: %v", err)
	} else {
		fmt.Println("✓ HDEL user:1000 age")
	}
}
