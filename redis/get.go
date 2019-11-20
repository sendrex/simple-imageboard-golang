package redis

// GetCachedPage returns a cached page or error.
func GetCachedPage(id uint64) (*map[string]interface{}, error) {
	key := GetPageKey(id)
	return getCachedModel(key)
}

// GetCachedThread returns a cached thread or error.
func GetCachedThread(id uint64) (*map[string]interface{}, error) {
	key := GetThreadKey(id)
	return getCachedModel(key)
}

// GetCachedPost returns a cached post or error.
func GetCachedPost(id uint64) (*map[string]interface{}, error) {
	key := GetPostKey(id)
	return getCachedModel(key)
}
