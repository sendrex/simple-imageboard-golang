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
		panic(err)
	}
	db, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		panic(err)
	}

	// Parse host and port into accepted address
	addr := fmt.Sprintf("%s:%d", host, port)

	// Connect the client
	client = redis.NewClient(&redis.Options{
		Addr:       addr,
		PoolSize:   1000,
		MaxRetries: 2,
		Password:   password,
		DB:         db,
	})

	// Check if it's connected
	if err := client.Ping().Err(); err != nil {
		panic(err)
	}

	// Show the result
	fmt.Println("[REDIS]: Redis client OK")
}
