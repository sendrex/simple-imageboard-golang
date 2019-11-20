package database

import (
	"github.com/AquoDev/simple-imageboard-golang/model"
)

// GetPost returns a post.
func GetPost(id uint64) (post model.Post, err error) {
	// Query post
	err = db.Select("id, content, pic, on_thread, created_at, updated_at").Where("id = ?", id).First(&post).Error
	return
}

// SavePost returns a struct with the ID and delete code of the inserted post.
func SavePost(post *model.Post) (result *model.DeleteData, err error) {
	// Try to insert the post
	if err = db.Create(&post).Error; err == nil {
		// If it's inserted, parse data
		result = &model.DeleteData{
			ID:         post.ID,
			DeleteCode: post.DeleteCode,
		}
	}

	return
}

// DeletePost returns an error that should be checked in the handler.
// Warning: if the post started a thread (on_thread == null), it will delete
// every post in the thread (on_thread == id).
func DeletePost(id uint64, code string) (err error) {
	// Make empty post
	post := new(model.Post)

	// Get post from the search
	result := db.Where("id = ? AND delete_code = ?", id, code).First(&post)

	// Check if the post hasn't been found
	if result.RecordNotFound() {
		// If it hasn't, return the error
		err = result.Error
	} else {
		// Delete the post if the delete code is correct
		err = db.Delete(&post).Error
	}

	return
}
