package database

import (
	"github.com/AquoDev/simple-imageboard-golang/model"
)

// GetIndex returns every post that started a thread.
func GetIndex() ([]model.Post, error) {
	// Query posts that started a thread (parent_thread IS NULL)
	index := make([]model.Post, 0)
	err := db.Select("id, content, pic, created_at, updated_at").Where("parent_thread IS NULL").Order("updated_at desc").Find(&index).Error

	// Return index and error
	return index, err
}
