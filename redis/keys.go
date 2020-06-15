package redis

import (
	"fmt"

	"github.com/AquoDev/simple-imageboard-golang/model"
)

const separator = ":"

// prefix returns a string with the initial key given the resource name.
func prefix(resource string) string {
	return fmt.Sprintf("%s%s%s", (&model.Post{}).TableName(), separator, resource)
}

// suffix returns the end of the key given the resource ID.
func suffix(id uint64) string {
	return fmt.Sprintf("%s%d", separator, id)
}

// getIndexKey returns a string that corresponds to 'index' key in Redis.
// Key format: "<table name>:index"
func getIndexKey() string {
	return prefix("index")
}

// getPageKey returns a string that corresponds to a 'page' key in Redis.
// Key format: "<table name>:page:<id>"
func getPageKey(id uint64) string {
	return fmt.Sprintf("%s%s", prefix("page"), suffix(id))
}

// getThreadKey returns a string that corresponds to a 'thread' key in Redis.
// Key format: "<table name>:thread:<id>"
func getThreadKey(id uint64) string {
	return fmt.Sprintf("%s%s", prefix("thread"), suffix(id))
}

// getPostKey returns a string that corresponds to a 'post' key in Redis.
// Key format: "<table name>:post:<id>"
func getPostKey(id uint64) string {
	return fmt.Sprintf("%s%s", prefix("post"), suffix(id))
}
