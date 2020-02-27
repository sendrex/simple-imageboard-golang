package redis

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v7"
)

// Redis client connection.
var client *redis.Client

func init() {
	// Get Redis host, port, password and database
	host := os.Getenv("REDIS_HOST")
	password := os.Getenv("REDIS_PASSWORD")
	port, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		message := fmt.Errorf("[REDIS] Couldn't parse REDIS_PORT @ %w", err)
		panic(message)
	}
	db, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		message := fmt.Errorf("[REDIS] Couldn't parse REDIS_DATABASE @ %w", err)
		panic(message)
	}

	// Parse host and port into accepted address
	addr := fmt.Sprintf("%s:%d", host, port)

	// Connect the client
	client = redis.NewClient(&redis.Options{
		Addr:       addr,
		PoolSize:   100,
		MaxRetries: 1,
		Password:   password,
		DB:         db,
	})

	// Check if it's connected
	if err := client.Ping().Err(); err != nil {
		message := fmt.Errorf("[REDIS] Client connection FAILED @ %w", err)
		panic(message)
	}

	// Show the result
	fmt.Println("   [REDIS] Client connection OK")
}
