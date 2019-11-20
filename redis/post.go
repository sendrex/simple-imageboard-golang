package redis

import (
	"time"

	"github.com/AquoDev/simple-imageboard-golang/model"
)

// GetCachedPost returns a cached post.
func GetCachedPost(id uint64) (*model.Post, error) {
	// Get key from post ID
	key := GetPostKey(id)

	// Get cached post
	result, err := client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	// Unmarshal result to a post
	return UnmarshalPost(result)
}

// SetCachedPost sets a post in cache.
func SetCachedPost(id uint64, post *model.Post) error {
	// Marshal post into JSON
	cachedPost, err := MarshalModel(post)
	if err != nil {
		return err
	}

	key := GetPostKey(id)
	return client.Set(key, string(cachedPost), 30*time.Second).Err()
}

// SetErrorCachedPost sets an error with a post ID in cache.
func SetErrorCachedPost(id uint64, data map[string]interface{}) error {
	// Marshal error into JSON
	cachedError, err := MarshalModel(data)
	if err != nil {
		return err
	}

	key := GetPostKey(id)
	return client.Set(key, string(cachedError), 30*time.Second).Err()
}
