package methods

import (
	"encoding/json"
	"time"
)

// PostThread is the same struct, but with less fields to show on thread.
type PostThread struct {
	ID        uint64    `json:"id"`
	Content   string    `json:"content"`
	Pic       string    `json:"pic"`
	ReplyTo   uint64    `json:"reply_to"`
	CreatedAt time.Time `json:"created_at"`
}

// GetThread returns a JSON with a thread (original post + on_thread == OP.id).
func GetThread(id uint64) (string, error) {
	// Make empty struct of posts slice
	var thread struct {
		Posts []PostThread
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
