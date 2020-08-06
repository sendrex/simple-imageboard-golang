package cache

import (
	"fmt"

	"github.com/AquoDev/simple-imageboard-golang/env"
	"github.com/go-redis/redis/v7"
)

// Redis client connection.
var client *redis.Client

func connect(opt *redis.Options) error {
	client = redis.NewClient(opt)
	return client.Ping().Err()
}

func init() {
	// Connect the client and panic if there are errors
	if err := connect(&redis.Options{
		Addr:       fmt.Sprintf("%s:%d", env.GetString("REDIS_HOST"), env.GetInt("REDIS_PORT")),
		PoolSize:   100,
		MaxRetries: 1,
		Password:   env.GetString("REDIS_PASSWORD"),
		DB:         env.GetInt("REDIS_DATABASE"),
	}); err != nil {
		message := fmt.Errorf("[CACHE] Redis connection failed @ %w", err)
		panic(message)
	}

	fmt.Println("[CACHE] Redis connection OK")
}
