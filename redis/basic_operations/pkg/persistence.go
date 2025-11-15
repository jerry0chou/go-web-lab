package redisops

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func DemonstratePersistence(ctx context.Context, rdb *redis.Client) {
	fmt.Println("\n--- PERSISTENCE Operations ---")

	// Save data
	err := rdb.Set(ctx, "persist:data1", "important data", 0).Err()
	if err != nil {
		log.Printf("Error setting key: %v", err)
		return
	}
	rdb.Set(ctx, "persist:data2", "more data", 0)
	rdb.Set(ctx, "persist:data3", "even more", 0)
	fmt.Println("✓ Set multiple keys for persistence demo")

	// SAVE - Force save (synchronous, blocks until complete)
	// Note: In production, use BGSAVE instead
	fmt.Println("✓ SAVE command would force synchronous save to disk")
	fmt.Println("  (Not executing in demo to avoid blocking)")

	// BGSAVE - Background save (asynchronous)
	// Note: This is typically configured in redis.conf
	fmt.Println("✓ BGSAVE command would trigger background save")
	fmt.Println("  (Redis auto-saves based on configuration)")

	// Check if data persists (simulated)
	val, err := rdb.Get(ctx, "persist:data1").Result()
	if err != nil {
		log.Printf("Error getting key: %v", err)
	} else {
		fmt.Printf("✓ Retrieved persisted data: persist:data1 = %s\n", val)
	}

	fmt.Println("\n  Note: Redis persistence is configured via:")
	fmt.Println("  - RDB (snapshot): save points in redis.conf")
	fmt.Println("  - AOF (Append Only File): logs every write operation")
	fmt.Println("  - Both can be enabled for maximum durability")
}
