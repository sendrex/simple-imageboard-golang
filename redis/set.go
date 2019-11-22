package redis

import (
	"time"

	"github.com/AquoDev/simple-imageboard-golang/model"
)

// setCachedModel caches any generic struct or interface as a JSON string.
func setCachedModel(key string, cache *model.Cache, duration time.Duration) error {
	// Marshal data into a JSON string
	cachedData, err := marshalModel(cache)
	if err != nil {
		return err
	}
	// Set JSON string in cache
	return client.Set(key, string(cachedData), duration).Err()
}

// SetCachedPage sets a page or error in cache.
func SetCachedPage(id uint64, status int, data interface{}) error {
	key := getPageKey(id)
	cache := &model.Cache{
		Status: status,
		Data:   data,
	}
	return setCachedModel(key, cache, maxTimePage)
}

// SetCachedThread sets a thread or error in cache.
func SetCachedThread(id uint64, status int, data interface{}) error {
	key := getThreadKey(id)
	cache := &model.Cache{
		Status: status,
		Data:   data,
	}
	return setCachedModel(key, cache, maxTimeThread)
}

// SetCachedPost sets a post or error in cache.
func SetCachedPost(id uint64, status int, data interface{}) error {
	key := getPostKey(id)
	cache := &model.Cache{
		Status: status,
		Data:   data,
	}
	return setCachedModel(key, cache, maxTimePost)
}
