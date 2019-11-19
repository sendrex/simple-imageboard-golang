package methods

import (
	db "github.com/AquoDev/simple-imageboard-golang/database"
	"github.com/AquoDev/simple-imageboard-golang/server/utils"
)

// GetPost returns a post.
func GetPost(id uint64) (post db.Post, err error) {
	// Query post
	err = db.Client().Select("id, content, pic, on_thread, created_at, updated_at").Where("id = ?", id).First(&post).Error
	return
}

// SavePost returns a struct with the ID and delete code of the inserted post.
func SavePost(post *db.Post) (result *utils.DeleteData, err error) {
	// Try to insert the post
	if err = db.Client().Create(&post).Error; err == nil {
		// If it's inserted, parse data
		result = &utils.DeleteData{
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
	post := new(db.Post)

	// Get post from the search
	result := db.Client().Where("id = ? AND delete_code = ?", id, code).First(&post)

	// Check if the post hasn't been found
	if result.RecordNotFound() {
		// If it hasn't, return the error
		err = result.Error
	} else {
		// Delete the post if the delete code is correct
		err = db.Client().Delete(&post).Error
	}

	return
}
