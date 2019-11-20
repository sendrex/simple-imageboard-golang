package redis

import (
	"time"

	"github.com/AquoDev/simple-imageboard-golang/model"
)

// GetCachedPage returns a cached page.
func GetCachedPage(id uint64) ([]model.Post, error) {
	// Get key for this page
	key := GetPageKey(id)

	// Get cached page
	return getCachedPostSlice(key)
}

// SetCachedPage sets a page in cache.
func SetCachedPage(id uint64, page []model.Post) error {
	// Marshal page into JSON
	cachedPage, err := MarshalModel(page)
	if err != nil {
		return err
	}

	key := GetPageKey(id)
	return client.Set(key, string(cachedPage), 15*time.Second).Err()
}

// SetErrorCachedPage sets an error with a page ID in cache.
func SetErrorCachedPage(id uint64, data map[string]interface{}) error {
	// Marshal error into JSON
	cachedError, err := MarshalModel(data)
	if err != nil {
		return err
	}

	key := GetPageKey(id)
	return client.Set(key, string(cachedError), 15*time.Second).Err()
}
