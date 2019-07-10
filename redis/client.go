package redis

import (
	"fmt"
	"os"
	"strconv"

	redis "github.com/go-redis/redis"
)

// Redis client connection.
var client *redis.Client

// GetRedisClient returns an already connected Redis client for caching.
func GetRedisClient() *redis.Client {
	return client
}

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
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}
