package database

import (
	"github.com/AquoDev/simple-imageboard-golang/model"
)

// GetPost returns a post.
func GetPost(id uint64) (*model.Post, error) {
	// Make empty post
	post := new(model.Post)

	// Query post
	err := db.Select("id, content, pic, on_thread, created_at, updated_at").Where("id = ?", id).First(&post).Error

	return post, err
}

// SavePost returns a struct with the ID and delete code of the inserted post.
func SavePost(post *model.Post) (*model.DeleteData, error) {
	// Try to insert the post
	if err := db.Create(&post).Error; err != nil {
		return nil, err
	}

	// If it's inserted, parse data and return it
	return &model.DeleteData{
		ID:         post.ID,
		DeleteCode: post.DeleteCode,
	}, nil
}

// DeletePost returns an error that should be checked in the handler.
// Warning: if the post started a thread (on_thread == null), it will delete
// every post in the thread (on_thread == id).
func DeletePost(id uint64, code string) error {
	// Make empty post
	post := new(model.Post)

	// Query post and check if the post hasn't been found
	if result := db.Where("id = ? AND delete_code = ?", id, code).First(&post); result.RecordNotFound() {
		// If it hasn't, return the error
		return result.Error
	}

	// Delete the post if the delete code is correct
	return db.Delete(&post).Error
}
