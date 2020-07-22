package redis

import (
	"encoding/json"

	"github.com/AquoDev/simple-imageboard-golang/model"
)

// getCachedModel returns a generic cached model.
func getCachedModel(key string) (*model.Cache, error) {
	// Get cached JSON
	result, err := client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	// Parse JSON into a struct
	cache := new(model.Cache)
	err = json.Unmarshal([]byte(result), &cache)

	// Return cached data and error
	return cache, err
}

// GetCachedIndex returns a cached index or error.
func GetCachedIndex() (*model.Cache, error) {
	key := getIndexKey()
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
