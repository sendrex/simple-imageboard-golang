package methods

import (
	"encoding/json"

	"github.com/AquoDev/simple-imageboard-golang/database"
)

// Thread has a 'posts' slice to be filled when it's requested.
type Thread struct {
	Posts []database.Post
}

// ThreadsToDelete has a 'posts' slice where we'll save old threads to delete.
type ThreadsToDelete struct {
	Collection []database.Post
}

// GetThread returns a JSON with a thread (original post + on_thread == original post ID).
func GetThread(id uint64) (string, error) {
	// Make empty thread
	thread := new(Thread)

	// Query posts that belong to a thread
	err := db.Select("id, content, pic, reply_to, created_at").Where("id = ?", id).Or("on_thread = ?", id).Order("id asc").Find(&thread.Posts).Error
	if err != nil {
		return "", err
	}

	// Convert result into JSON
	result, err := json.Marshal(thread.Posts)

	return string(result), err
}

// DeleteOldThreads deletes any thread older than the last bump date in page 9.
func DeleteOldThreads() (err error) {
	// Make empty collection
	threads := new(ThreadsToDelete)

	// Query old threads and save them in the collection
	result := db.Where("on_thread IS NULL").Order("updated_at desc").Offset(300).Find(&threads.Collection)

	// Check if any threads have been found
	if result.RecordNotFound() {
		// If any record hasn't been found, return the error
		err = result.Error
	} else {
		// Delete old threads inside the collection
		err = db.Delete(&threads.Collection).Error
	}

	return
}
