package redisops

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func DemonstrateLists(ctx context.Context, rdb *redis.Client) {
	fmt.Println("\n--- LIST Operations ---")

	// LPUSH - Push to left (head)
	err := rdb.LPush(ctx, "tasks", "task1", "task2", "task3").Err()
	if err != nil {
		log.Printf("Error pushing to list: %v", err)
		return
	}
	fmt.Println("✓ LPUSH tasks task1 task2 task3")

	// RPUSH - Push to right (tail)
	err = rdb.RPush(ctx, "tasks", "task4", "task5").Err()
	if err != nil {
		log.Printf("Error pushing to list: %v", err)
	} else {
		fmt.Println("✓ RPUSH tasks task4 task5")
	}

	// LRANGE - Get range of elements
	tasks, err := rdb.LRange(ctx, "tasks", 0, -1).Result()
	if err != nil {
		log.Printf("Error getting list range: %v", err)
	} else {
		fmt.Printf("✓ LRANGE tasks 0 -1 = %v\n", tasks)
	}

	// LLEN - Get list length
	length, err := rdb.LLen(ctx, "tasks").Result()
	if err != nil {
		log.Printf("Error getting list length: %v", err)
	} else {
		fmt.Printf("✓ LLEN tasks = %d\n", length)
	}

	// LPOP - Pop from left
	task, err := rdb.LPop(ctx, "tasks").Result()
	if err != nil {
		log.Printf("Error popping from list: %v", err)
	} else {
		fmt.Printf("✓ LPOP tasks = %s\n", task)
	}

	// RPOP - Pop from right
	task, err = rdb.RPop(ctx, "tasks").Result()
	if err != nil {
		log.Printf("Error popping from list: %v", err)
	} else {
		fmt.Printf("✓ RPOP tasks = %s\n", task)
	}
}
