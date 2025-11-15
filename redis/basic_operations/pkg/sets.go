package redisops

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

// ============================================
// SET Operations
// ============================================
func DemonstrateSets(ctx context.Context, rdb *redis.Client) {
	fmt.Println("\n--- SET Operations ---")

	// SADD - Add members to set
	err := rdb.SAdd(ctx, "tags", "golang", "redis", "docker", "backend").Err()
	if err != nil {
		log.Printf("Error adding to set: %v", err)
		return
	}
	fmt.Println("✓ SADD tags golang redis docker backend")

	// SMEMBERS - Get all members
	members, err := rdb.SMembers(ctx, "tags").Result()
	if err != nil {
		log.Printf("Error getting set members: %v", err)
	} else {
		fmt.Printf("✓ SMEMBERS tags = %v\n", members)
	}

	// SISMEMBER - Check if member exists
	exists, err := rdb.SIsMember(ctx, "tags", "golang").Result()
	if err != nil {
		log.Printf("Error checking set membership: %v", err)
	} else {
		fmt.Printf("✓ SISMEMBER tags golang = %v\n", exists)
	}

	// SCARD - Get set cardinality (size)
	size, err := rdb.SCard(ctx, "tags").Result()
	if err != nil {
		log.Printf("Error getting set size: %v", err)
	} else {
		fmt.Printf("✓ SCARD tags = %d\n", size)
	}

	// SREM - Remove member from set
	err = rdb.SRem(ctx, "tags", "docker").Err()
	if err != nil {
		log.Printf("Error removing from set: %v", err)
	} else {
		fmt.Println("✓ SREM tags docker")
	}

	// Set intersection example
	rdb.SAdd(ctx, "set1", "a", "b", "c")
	rdb.SAdd(ctx, "set2", "b", "c", "d")
	intersection, _ := rdb.SInter(ctx, "set1", "set2").Result()
	fmt.Printf("✓ SINTER set1 set2 = %v\n", intersection)
}
