package redis

import "fmt"

// makeKey returns a string built with both arguments.
func makeKey(prefix string, number uint64) string {
	return fmt.Sprintf("%s:%d", prefix, number)
}

// GetPageKey returns a string that corresponds to that key in Redis.
func GetPageKey(id uint64) string {
	return makeKey("page", id)
}
