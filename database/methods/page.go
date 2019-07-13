package methods

import (
	"encoding/json"

	"github.com/AquoDev/simple-imageboard-golang/database"
)

// GetPage returns a JSON with a thread list.
func GetPage(id uint64) (string, error) {
	// Make empty struct of posts slice
	var page struct {
		Posts []database.Post
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
