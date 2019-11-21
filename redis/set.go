package redis

import (
	"time"
)

// setCachedModel caches any generic struct or interface as a JSON string.
func setCachedModel(key string, data interface{}, duration time.Duration) error {
	// Marshal data into a JSON string
	cachedData, err := marshalModel(data)
	if err != nil {
		return err
	}
	// Set JSON string in cache
	return client.Set(key, string(cachedData), duration).Err()
}

// SetCachedPage sets a page or error in cache.
func SetCachedPage(id uint64, data interface{}) error {
	key := getPageKey(id)
	return setCachedModel(key, data, maxTimePage)
}

// SetCachedThread sets a thread or error in cache.
func SetCachedThread(id uint64, data interface{}) error {
	key := getThreadKey(id)
	return setCachedModel(key, data, maxTimeThread)
}

// SetCachedPost sets a post or error in cache.
func SetCachedPost(id uint64, data interface{}) error {
	key := getPostKey(id)
	return setCachedModel(key, data, maxTimePost)
}
