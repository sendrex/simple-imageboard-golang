package redis

import (
	"fmt"
	"time"
)

// makeKey returns a string built with both arguments.
func makeKey(prefix string, number uint64) string {
	return fmt.Sprintf("%s:%d", prefix, number)
}

// getCachedModel returns a generic cached model.
func getCachedModel(key string) (*map[string]interface{}, error) {
	// Get cached result
	result, err := client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	// Parse and return cached result
	return UnmarshalModel(result)
}

// setCachedModel caches any generic struct or interface as a JSON string.
func setCachedModel(key string, data interface{}, duration time.Duration) error {
	// Marshal data into a JSON string
	cachedData, err := MarshalModel(data)
	if err != nil {
		return err
	}

	return client.Set(key, string(cachedData), duration).Err()
}
