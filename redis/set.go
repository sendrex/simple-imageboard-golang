package redis

// SetCachedPage sets a page or error in cache.
func SetCachedPage(id uint64, data interface{}) error {
	key := GetPageKey(id)
	return setCachedModel(key, data, maxTimePage)
}

// SetCachedThread sets a thread or error in cache.
func SetCachedThread(id uint64, data interface{}) error {
	key := GetThreadKey(id)
	return setCachedModel(key, data, maxTimeThread)
}

// SetCachedPost sets a post or error in cache.
func SetCachedPost(id uint64, data interface{}) error {
	key := GetPostKey(id)
	return setCachedModel(key, data, maxTimePost)
}
