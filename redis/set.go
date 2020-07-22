package redis

import (
	"encoding/json"
	"time"

	"github.com/AquoDev/simple-imageboard-golang/env"
	"github.com/AquoDev/simple-imageboard-golang/model"
)

var (
	maxTimeIndex  = env.GetTime("REDIS_EXPIRE_TIME_INDEX", "s")
	maxTimeThread = env.GetTime("REDIS_EXPIRE_TIME_THREAD", "s")
	maxTimePost   = env.GetTime("REDIS_EXPIRE_TIME_POST", "s")
)

// setCachedModel caches any struct as JSON.
func setCachedModel(key string, status int, data interface{}, duration time.Duration) error {
	// Parse struct into JSON
	cache, err := json.Marshal(&model.Cache{
		Status: status,
		Data:   data,
	})
	if err != nil {
		return err
	}

	// Cache JSON
	return client.Set(key, string(cache), duration).Err()
}

// SetCachedIndex sets an index or error in cache.
func SetCachedIndex(status int, data interface{}) error {
	key := getIndexKey()
	return setCachedModel(key, status, data, maxTimeIndex)
}

// SetCachedThread sets a thread or error in cache.
func SetCachedThread(id uint64, status int, data interface{}) error {
	key := getThreadKey(id)
	return setCachedModel(key, status, data, maxTimeThread)
}

// SetCachedPost sets a post or error in cache.
func SetCachedPost(id uint64, status int, data interface{}) error {
	key := getPostKey(id)
	return setCachedModel(key, status, data, maxTimePost)
}
