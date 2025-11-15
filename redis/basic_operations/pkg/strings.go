package redisops

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func DemonstrateStrings(ctx context.Context, rdb *redis.Client) {
	fmt.Println("\n--- STRING Operations ---")

	// SET - Set a key-value pair
	err := rdb.Set(ctx, "user:name", "John Doe", 0).Err()
	if err != nil {
		log.Printf("Error setting string: %v", err)
		return
	}
	fmt.Println("✓ SET user:name 'John Doe'")

	// GET - Get value by key
	val, err := rdb.Get(ctx, "user:name").Result()
	if err != nil {
		log.Printf("Error getting string: %v", err)
	} else {
		fmt.Printf("✓ GET user:name = %s\n", val)
	}

	// SET with expiration
	err = rdb.Set(ctx, "session:token", "abc123xyz", 10*time.Second).Err()
	if err != nil {
		log.Printf("Error setting string with expiration: %v", err)
	} else {
		fmt.Println("✓ SET session:token 'abc123xyz' EX 10")
	}

	// INCR - Increment a number
	count, err := rdb.Incr(ctx, "counter").Result()
	if err != nil {
		log.Printf("Error incrementing: %v", err)
	} else {
		fmt.Printf("✓ INCR counter = %d\n", count)
	}

	// MSET - Set multiple keys at once
	err = rdb.MSet(ctx, "key1", "value1", "key2", "value2", "key3", "value3").Err()
	if err != nil {
		log.Printf("Error setting multiple keys: %v", err)
	} else {
		fmt.Println("✓ MSET key1 value1 key2 value2 key3 value3")
	}

	// MGET - Get multiple values
	vals, err := rdb.MGet(ctx, "key1", "key2", "key3").Result()
	if err != nil {
		log.Printf("Error getting multiple keys: %v", err)
	} else {
		fmt.Printf("✓ MGET key1 key2 key3 = %v\n", vals)
	}
}
