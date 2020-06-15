package database

import (
	"github.com/AquoDev/simple-imageboard-golang/model"
)

// GetIndex returns every post that started a thread (post slice).
func GetIndex() ([]model.Post, error) {
	// Make empty list
	page := make([]model.Post, 0)

	// Query posts that started a thread (parent_thread IS NULL)
	if err := db.Select("id, content, pic, created_at, updated_at").Where("parent_thread IS NULL").Order("updated_at desc").Find(&page).Error; err != nil {
		return nil, err
	}

	return page, nil
}
