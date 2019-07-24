package methods

import (
	"encoding/json"

	"github.com/AquoDev/simple-imageboard-golang/database"
)

// GetPost returns a JSON with a post.
func GetPost(id uint64) (string, error) {
	// Make empty post
	post := new(database.Post)

	// Query post
	err := db.Select("id, content, pic, on_thread, reply_to, created_at, updated_at").Where("id = ?", id).Find(&post).Error
	if err != nil {
		return "", err
	}

	// Convert result into JSON
	result, err := json.Marshal(post)

	return string(result), err
}

// TODO implement "INSERT INTO" data to the database
func SavePost() {

}

// DeletePost returns an error that should be checked in the handler.
// Warning: if the post started a thread (on_thread == null), it will delete
// every post in the thread (on_thread == id).
func DeletePost(id uint64, code string) (err error) {
	// Make empty post
	post := new(database.Post)

	// Select post to delete
	err = db.Where("id = ? AND delete_code = ?", id, code).Find(&post).Error
	if err != nil {
		// If it's not found, return the error
		return
	}

	// Delete the post if the delete code is correct
	err = db.Delete(&post).Error
	return
}
