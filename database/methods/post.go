package methods

// TODO implement POST/DELETE methods

import (
	"encoding/json"

	"github.com/AquoDev/simple-imageboard-golang/database"
)

// GetPost returns a JSON with a post.
func GetPost(id uint64) (string, error) {
	// Make empty post
	post := new(database.Post)

	// Query post
	err := db.Select("id, content, pic, on_thread, reply_to, created_at, updated_at").Where("id = ?", id).Find(&post).Error
	if err != nil {
		return "", err
	}

	// Convert result into JSON
	result, err := json.Marshal(post)

	return string(result), err
}
