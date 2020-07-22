package database

import (
	"github.com/AquoDev/simple-imageboard-golang/model"
)

// GetPost returns a post given its ID.
func GetPost(id uint64) (*model.Post, error) {
	// Query post given its ID
	post := new(model.Post)
	err := db.Select("id, content, pic, parent_thread, reply_to, created_at, updated_at").Where("id = ?", id).First(&post).Error

	// Return post and error
	return post, err
}

// SavePost returns a struct with the ID and delete code of the inserted post.
func SavePost(post *model.Post) (*model.DeleteData, error) {
	// Try to insert the post
	if err := db.Create(&post).Error; err != nil {
		return nil, err
	}

	// If it's been inserted, return delete data without error
	return &model.DeleteData{
		ID:         post.ID,
		DeleteCode: post.DeleteCode,
	}, nil
}

// DeletePost returns an error that should be checked in the handler.
// Warning: if the post started a thread (parent_thread == nil), it will delete
// every post in the thread (parent_thread == id).
func DeletePost(data *model.DeleteData) error {
	// Query post and check if the data is valid
	post := new(model.Post)
	err := db.Where("id = ? AND delete_code = ?", data.ID, data.DeleteCode).First(&post).Error

	// If there's any error, return it
	if err != nil {
		return err
	}

	// Delete the post
	return db.Delete(&post).Error
}
