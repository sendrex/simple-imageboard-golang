package redis

import (
	"github.com/AquoDev/simple-imageboard-golang/model"
)

// getCachedModel returns a generic cached model.
func getCachedModel(key string) (*model.Cache, error) {
	// Get cached result
	result, err := client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	// Parse and return cached result
	return unmarshalModel(result)
}

// GetCachedPage returns a cached page or error.
func GetCachedPage(id uint64) (*model.Cache, error) {
	key := getPageKey(id)
	return getCachedModel(key)
}

// GetCachedThread returns a cached thread or error.
func GetCachedThread(id uint64) (*model.Cache, error) {
	key := getThreadKey(id)
	return getCachedModel(key)
}

// GetCachedPost returns a cached post or error.
func GetCachedPost(id uint64) (*model.Cache, error) {
	key := getPostKey(id)
	return getCachedModel(key)
}
