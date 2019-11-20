package redis

import (
	"github.com/AquoDev/simple-imageboard-golang/model"
)

// getCachedPage returns a cached post slice (used for pages and threads).
func getCachedPostSlice(key string) ([]model.Post, error) {
	// Get cached result for this key
	result, err := client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	// Unmarshal result to a post slice
	postSlice, err := UnmarshalPostSlice(result)
	if err != nil {
		return nil, err
	}

	return postSlice, nil
}
