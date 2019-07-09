package redis

import "github.com/go-redis/redis"

// Redis client connection.
var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

// GetRedisClient returns an already existing Redis client for caching.
func GetRedisClient() (*redis.Client) {
	return client
}
