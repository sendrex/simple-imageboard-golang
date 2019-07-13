package redis

import "fmt"

// makeKey returns a string built with both arguments.
func makeKey(prefix string, number uint64) string {
	return fmt.Sprintf("%s:%d", prefix, number)
}

// GetPageKey returns a string that corresponds to a 'page' key in Redis.
func GetPageKey(id uint64) string {
	return makeKey("page", id)
}

// GetThreadKey returns a string that corresponds to a 'thread' key in Redis.
func GetThreadKey(id uint64) string {
	return makeKey("thread", id)
}
