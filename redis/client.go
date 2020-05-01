package redis

import (
	"fmt"

	"github.com/AquoDev/simple-imageboard-golang/env"
	"github.com/go-redis/redis/v7"
)

// Redis client connection.
var client *redis.Client

func init() {
	// Connect the client
	client = redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%s:%d", env.GetString("REDIS_HOST"), env.GetInt("REDIS_PORT")),
		PoolSize:   100,
		MaxRetries: 1,
		Password:   env.GetString("REDIS_PASSWORD"),
		DB:         env.GetInt("REDIS_DATABASE"),
	})

	// Check if it's connected
	if err := client.Ping().Err(); err != nil {
		message := fmt.Errorf("[REDIS] Client connection FAILED @ %w", err)
		panic(message)
	} else {
		fmt.Println("   [REDIS] Client connection OK")
	}
}
