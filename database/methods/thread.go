package methods

import (
	"encoding/json"

	"github.com/AquoDev/simple-imageboard-golang/database"
)

// Thread has a 'posts' slice to be filled when it's requested.
type Thread struct {
	Posts []database.Post
}

// GetThread returns a JSON with a thread (original post + on_thread == original post ID).
func GetThread(id uint64) (string, error) {
	// Make empty thread
	thread := new(Thread)

	// Query posts that belong to a thread
	err := db.Select("id, content, pic, created_at").Where("id = ?", id).Or("on_thread = ?", id).Order("id asc").Find(&thread.Posts).Error
	if err != nil {
		return "", err
	}

	// Convert result into JSON
	result, err := json.Marshal(thread.Posts)

	return string(result), err
}

// DeleteOldThreads deletes any thread older than the last bump date in page 9.
func DeleteOldThreads() (err error) {
	// Make empty slice of IDs
	var threads []uint64

	// Query ID from old threads and save them in the slice
	result := db.Model(&database.Post{}).Offset(100).Where("on_thread IS NULL").Order("updated_at desc").Pluck("id", &threads)

	// Check if there aren't IDs found
	if len(threads) == 0 {
		// If there's none, return the error
		err = result.Error
	} else {
		// Delete old threads as defined in the slice
		err = db.Where("id IN (?)", threads).Delete(&database.Post{}).Error
	}

	return
}
