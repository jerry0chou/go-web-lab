package main

import (
	"context"
	"fmt"
	"log"
	"os"

	redisops "go-web-lab/redis/basic_operations/pkg"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	// Get Redis address from environment variable, default to localhost for local development
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost"
	}
	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Test connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	fmt.Println("✅ Successfully connected to Redis!")

	// Run examples
	fmt.Println("\n" + redisops.Repeat("=", 60))
	fmt.Println("REDIS FUNDAMENTALS DEMONSTRATION")
	fmt.Println(redisops.Repeat("=", 60))

	redisops.DemonstrateStrings(ctx, rdb)
	redisops.DemonstrateHashes(ctx, rdb)
	redisops.DemonstrateLists(ctx, rdb)
	redisops.DemonstrateSets(ctx, rdb)
	redisops.DemonstrateSortedSets(ctx, rdb)
	redisops.DemonstrateExpiration(ctx, rdb)
	redisops.DemonstratePersistence(ctx, rdb)

	fmt.Println("\n✅ All Redis examples completed successfully!")
}
