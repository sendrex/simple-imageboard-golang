package redis

import (
	"time"

	"github.com/AquoDev/simple-imageboard-golang/model"
)

// GetCachedThread returns a cached thread.
func GetCachedThread(id uint64) ([]model.Post, error) {
	// Get key for this thread
	key := GetThreadKey(id)

	// Get cached thread
	return getCachedPostSlice(key)
}

// SetCachedThread sets a thread in cache.
func SetCachedThread(id uint64, thread []model.Post) error {
	// Marshal thread into JSON
	cachedThread, err := MarshalModel(thread)
	if err != nil {
		return err
	}

	key := GetThreadKey(id)
	return client.Set(key, string(cachedThread), 15*time.Second).Err()
}

// SetErrorCachedThread sets an error with a thread ID in cache.
func SetErrorCachedThread(id uint64, data map[string]interface{}) error {
	// Marshal error into JSON
	cachedError, err := MarshalModel(data)
	if err != nil {
		return err
	}

	key := GetThreadKey(id)
	return client.Set(key, string(cachedError), 15*time.Second).Err()
}
