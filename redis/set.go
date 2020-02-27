package redis

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/AquoDev/simple-imageboard-golang/model"
)

var (
	maxTimePage, maxTimeThread, maxTimePost time.Duration
)

// setCachedModel caches any generic struct or interface as a JSON string.
func setCachedModel(key string, status int, data interface{}, duration time.Duration) error {
	// Marshal status and data into JSON
	cache, err := json.Marshal(&model.Cache{
		Status: status,
		Data:   data,
	})
	if err != nil {
		return err
	}
	// Set JSON in cache
	return client.Set(key, string(cache), duration).Err()
}

// SetCachedPage sets a page or error in cache.
func SetCachedPage(id uint64, status int, data interface{}) error {
	key := getPageKey(id)
	return setCachedModel(key, status, data, maxTimePage)
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

func init() {
	stringTimePage := fmt.Sprintf("%ss", os.Getenv("REDIS_EXPIRE_TIME_PAGE"))
	stringTimeThread := fmt.Sprintf("%ss", os.Getenv("REDIS_EXPIRE_TIME_THREAD"))
	stringTimePost := fmt.Sprintf("%ss", os.Getenv("REDIS_EXPIRE_TIME_POST"))

	// Parse every expiry time
	if parsedTime, err := time.ParseDuration(stringTimePage); err != nil {
		message := fmt.Errorf("[REDIS] Couldn't parse REDIS_EXPIRE_TIME_PAGE: %w", err)
		panic(message)
	} else {
		maxTimePage = parsedTime
	}

	if parsedTime, err := time.ParseDuration(stringTimeThread); err != nil {
		message := fmt.Errorf("[REDIS] Couldn't parse REDIS_EXPIRE_TIME_THREAD: %w", err)
		panic(message)
	} else {
		maxTimeThread = parsedTime
	}

	if parsedTime, err := time.ParseDuration(stringTimePost); err != nil {
		message := fmt.Errorf("[REDIS] Couldn't parse REDIS_EXPIRE_TIME_POST: %w", err)
		panic(message)
	} else {
		maxTimePost = parsedTime
	}
}
