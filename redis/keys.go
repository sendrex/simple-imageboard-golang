package redis

// GetPageKey returns a string that corresponds to a 'page' key in Redis.
func GetPageKey(id uint64) string {
	return makeKey("page", id)
}

// GetThreadKey returns a string that corresponds to a 'thread' key in Redis.
func GetThreadKey(id uint64) string {
	return makeKey("thread", id)
}

// GetPostKey returns a string that corresponds to a 'post' key in Redis.
func GetPostKey(id uint64) string {
	return makeKey("post", id)
}
