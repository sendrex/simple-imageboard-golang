package methods

import (
	"time"

	db "github.com/AquoDev/simple-imageboard-golang/database"
)

// GetThread returns a slice of posts (original post + on_thread == original post ID).
func GetThread(id uint64) (thread []db.Post, err error) {
	// Query posts that belong to a thread
	err = db.Client().Select("id, content, pic, created_at").Where("id = ?", id).Or("on_thread = ?", id).Order("id asc").Find(&thread).Error
	return
}

// DeleteOldThreads deletes any thread older than the last bump date in page 9.
func DeleteOldThreads() (err error) {
	// Make empty slice of IDs
	threadIDs := make([]uint64, 0)

	// Query ID from old threads and save them in the slice
	result := db.Client().Model(&db.Post{}).Offset(100).Where("on_thread IS NULL").Order("updated_at desc").Pluck("id", &threadIDs)

	// Check if there aren't IDs found
	if len(threadIDs) == 0 {
		// If there's none, return the error
		err = result.Error
	} else {
		// Delete old threads as defined in the slice
		err = db.Client().Where("id IN (?)", threadIDs).Delete(&db.Post{}).Error
	}

	return
}

// BumpThread updates the post's "updated_at" field.
func BumpThread(id uint64, updatedAt *time.Time) (err error) {
	// Make empty thread
	thread := make([]db.Post, 0)

	// Query posts that belong to a thread
	err = db.Client().Where("id = ?", id).Or("on_thread = ?", id).Order("id asc").Find(&thread).Error
	if err != nil {
		return
	}

	// If there are less than 300 posts in the thread, update it
	if len(thread) < 300 {
		err = db.Client().Model(&db.Post{}).Where("id = ?", id).Update("updated_at", updatedAt).Error
	}

	return
}
