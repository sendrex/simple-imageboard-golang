package methods

import (
	"encoding/json"

	"github.com/AquoDev/simple-imageboard-golang/database"
)

// GetThread returns a JSON with a thread (original post + on_thread == original post ID).
func GetThread(id uint64) (string, error) {
	// Make empty struct of posts slice
	var thread struct {
		Posts []database.Post
	}

	// Query posts that belong to a thread
	err := db.Select("id, content, pic, reply_to, created_at").Where("id = ?", id).Or("on_thread = ?", id).Order("id asc").Find(&thread.Posts).Error
	if err != nil {
		return "", err
	}

	// Convert result into JSON
	result, err := json.Marshal(thread.Posts)

	return string(result), err
}
