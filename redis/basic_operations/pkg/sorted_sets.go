package redisops

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

// ============================================
// SORTED SET (ZSet) Operations
// ============================================
func DemonstrateSortedSets(ctx context.Context, rdb *redis.Client) {
	fmt.Println("\n--- SORTED SET (ZSet) Operations ---")

	// ZADD - Add members with scores
	err := rdb.ZAdd(ctx, "leaderboard", redis.Z{
		Score:  1000,
		Member: "player1",
	}, redis.Z{
		Score:  1500,
		Member: "player2",
	}, redis.Z{
		Score:  800,
		Member: "player3",
	}, redis.Z{
		Score:  2000,
		Member: "player4",
	}).Err()
	if err != nil {
		log.Printf("Error adding to sorted set: %v", err)
		return
	}
	fmt.Println("✓ ZADD leaderboard 1000 player1 1500 player2 800 player3 2000 player4")

	// ZRANGE - Get range by rank (ascending)
	players, err := rdb.ZRange(ctx, "leaderboard", 0, -1).Result()
	if err != nil {
		log.Printf("Error getting sorted set range: %v", err)
	} else {
		fmt.Printf("✓ ZRANGE leaderboard 0 -1 = %v\n", players)
	}

	// ZREVRANGE - Get range by rank (descending)
	topPlayers, err := rdb.ZRevRange(ctx, "leaderboard", 0, 2).Result()
	if err != nil {
		log.Printf("Error getting reverse range: %v", err)
	} else {
		fmt.Printf("✓ ZREVRANGE leaderboard 0 2 = %v\n", topPlayers)
	}

	// ZRANGEBYSCORE - Get range by score
	highScores, err := rdb.ZRangeByScore(ctx, "leaderboard", &redis.ZRangeBy{
		Min: "1000",
		Max: "2000",
	}).Result()
	if err != nil {
		log.Printf("Error getting range by score: %v", err)
	} else {
		fmt.Printf("✓ ZRANGEBYSCORE leaderboard 1000 2000 = %v\n", highScores)
	}

	// ZSCORE - Get score of member
	score, err := rdb.ZScore(ctx, "leaderboard", "player2").Result()
	if err != nil {
		log.Printf("Error getting score: %v", err)
	} else {
		fmt.Printf("✓ ZSCORE leaderboard player2 = %.0f\n", score)
	}

	// ZRANK - Get rank of member (0-based, ascending)
	rank, err := rdb.ZRank(ctx, "leaderboard", "player2").Result()
	if err != nil {
		log.Printf("Error getting rank: %v", err)
	} else {
		fmt.Printf("✓ ZRANK leaderboard player2 = %d\n", rank)
	}

	// ZINCRBY - Increment score
	newScore, err := rdb.ZIncrBy(ctx, "leaderboard", 100, "player1").Result()
	if err != nil {
		log.Printf("Error incrementing score: %v", err)
	} else {
		fmt.Printf("✓ ZINCRBY leaderboard 100 player1 = %.0f\n", newScore)
	}
}
