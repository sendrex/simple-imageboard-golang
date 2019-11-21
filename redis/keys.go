package redis

import (
	"fmt"
)

// makeKey returns a string built with both arguments.
func makeKey(prefix string, number uint64) string {
	return fmt.Sprintf("%s:%d", prefix, number)
}

// getPageKey returns a string that corresponds to a 'page' key in Redis.
func getPageKey(id uint64) string {
	return makeKey("page", id)
}

// getThreadKey returns a string that corresponds to a 'thread' key in Redis.
func getThreadKey(id uint64) string {
	return makeKey("thread", id)
}

// getPostKey returns a string that corresponds to a 'post' key in Redis.
func getPostKey(id uint64) string {
	return makeKey("post", id)
}
