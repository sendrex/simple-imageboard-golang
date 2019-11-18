package methods

import (
	"encoding/json"

	"github.com/AquoDev/simple-imageboard-golang/database"
	"github.com/AquoDev/simple-imageboard-golang/server/utils"
)

// GetPost returns a JSON with a post.
func GetPost(id uint64) (string, error) {
	// Make empty post
	post := new(database.Post)

	// Query post
	err := db.Select("id, content, pic, on_thread, created_at, updated_at").Where("id = ?", id).First(&post).Error
	if err != nil {
		return "", err
	}

	// Convert result into JSON
	result, err := json.Marshal(post)

	return string(result), err
}

// SavePost returns a JSON with the ID and delete code of the inserted post.
func SavePost(post *database.Post) (result string, err error) {
	// Try to insert the post
	if err = db.Create(&post).Error; err != nil {
		// If it fails, the result will be empty
		result = ""
	} else {
		// If it's inserted, parse it into JSON
		jsonObject, _ := json.Marshal(&utils.DeleteData{
			ID:         post.ID,
			DeleteCode: post.DeleteCode,
		})
		result = string(jsonObject)
	}

	return
}

// DeletePost returns an error that should be checked in the handler.
// Warning: if the post started a thread (on_thread == null), it will delete
// every post in the thread (on_thread == id).
func DeletePost(id uint64, code string) (err error) {
	// Make empty post
	post := new(database.Post)

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
