package methods

import (
	"encoding/json"
	"time"

	"github.com/AquoDev/simple-imageboard-golang/database"
)

// Post is the same struct, but with less fields to show on page.
type Post struct {
	ID        uint64    `json:"id"`
	Content   string    `json:"content"`
	Pic       string    `json:"pic"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var db = database.GetDatabaseClient()

// GetPage returns a JSON with a thread list.
func GetPage(id uint64) (string, error) {
	// Make empty struct of posts slice
	var page struct {
		Posts []Post
	}

	// Query posts that started a thread (on_thread == null)
	err := db.Select("id, content, pic, created_at, updated_at").Offset(10 * id).Limit(10).Where("on_thread IS NULL").Order("updated_at desc").Find(&page.Posts).Error
	if err != nil {
		return "", err
	}

	// Convert result into JSON
	result, err := json.Marshal(page.Posts)

	return string(result), err
}
