package redisops

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func DemonstrateExpiration(ctx context.Context, rdb *redis.Client) {
	fmt.Println("\n--- EXPIRATION Operations ---")

	// SET with expiration
	err := rdb.Set(ctx, "temp:key", "temporary value", 5*time.Second).Err()
	if err != nil {
		log.Printf("Error setting key with expiration: %v", err)
		return
	}
	fmt.Println("✓ SET temp:key 'temporary value' EX 5")

	// Check TTL (Time To Live)
	ttl, err := rdb.TTL(ctx, "temp:key").Result()
	if err != nil {
		log.Printf("Error getting TTL: %v", err)
	} else {
		fmt.Printf("✓ TTL temp:key = %v\n", ttl)
	}

	// EXPIRE - Set expiration on existing key
	err = rdb.Set(ctx, "session:123", "active", 0).Err()
	if err != nil {
		log.Printf("Error setting key: %v", err)
	} else {
		rdb.Expire(ctx, "session:123", 10*time.Second)
		fmt.Println("✓ EXPIRE session:123 10")
	}

	// PERSIST - Remove expiration
	err = rdb.Set(ctx, "persistent:key", "value", 30*time.Second).Err()
	if err != nil {
		log.Printf("Error setting key: %v", err)
	} else {
		rdb.Persist(ctx, "persistent:key")
		fmt.Println("✓ PERSIST persistent:key (removed expiration)")

		// Verify it's persistent
		ttl, _ := rdb.TTL(ctx, "persistent:key").Result()
		if ttl == -1 {
			fmt.Println("  → Key is now persistent (no expiration)")
		}
	}

	// EXPIREAT - Set expiration at specific timestamp
	futureTime := time.Now().Add(1 * time.Minute)
	err = rdb.Set(ctx, "scheduled:key", "value", 0).Err()
	if err != nil {
		log.Printf("Error setting key: %v", err)
	} else {
		rdb.ExpireAt(ctx, "scheduled:key", futureTime)
		fmt.Printf("✓ EXPIREAT scheduled:key %v\n", futureTime.Unix())
	}
}
