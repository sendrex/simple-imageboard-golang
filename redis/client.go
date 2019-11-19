package redis

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

// Redis client connection.
var client *redis.Client

// Client returns an already connected Redis client for caching.
func Client() *redis.Client {
	return client
}

func init() {
	// Read values from .env file
	godotenv.Load()

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
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Check if it's connected
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	// Show the result
	fmt.Printf("Redis: %s received\n", pong)
}
