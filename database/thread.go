package database

import (
	"time"

	"github.com/AquoDev/simple-imageboard-golang/model"
)

// GetThread returns a slice of posts (original post + on_thread == original post ID).
func GetThread(id uint64) ([]model.Post, error) {
	// Make empty thread
	thread := make([]model.Post, 0)

	// Query posts that belong to a thread
	err := db.Select("id, content, pic, created_at").Where("id = ?", id).Or("on_thread = ?", id).Order("id asc").Find(&thread).Error

	return thread, err
}

// DeleteOldThreads deletes any thread older than the last bump date in page 9.
func DeleteOldThreads() error {
	// Make empty slice of IDs
	threadIDs := make([]uint64, 0)

	// Query ID from old threads and save them in the slice
	result := db.Model(&model.Post{}).Offset(100).Where("on_thread IS NULL").Order("updated_at desc").Pluck("id", &threadIDs)

	// Check if there aren't IDs found
	if len(threadIDs) == 0 {
		// If there's none, return the error
		return result.Error
	}

	// Delete old threads as defined in the slice
	return db.Where("id IN (?)", threadIDs).Delete(&model.Post{}).Error
}

// BumpThread updates the post's "updated_at" field.
func BumpThread(id uint64, updatedAt *time.Time) error {
	// Make empty thread
	thread := make([]model.Post, 0)

	// Query posts that belong to a thread
	if err := db.Where("id = ?", id).Or("on_thread = ?", id).Order("id asc").Find(&thread).Error; err != nil {
		return err
	}

	// If there are less than 300 posts in the thread, update it
	if len(thread) < 300 {
		return db.Model(&model.Post{}).Where("id = ?", id).Update("updated_at", updatedAt).Error
	}

	return nil
}
