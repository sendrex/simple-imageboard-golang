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
	if err := db.Select("posts.id, posts.content, posts.pic, posts.created_at, posts.updated_at, count(b.on_thread) as replies").Joins("LEFT JOIN posts b ON b.on_thread = posts.id").Where("posts.id = ?", id).Or("posts.on_thread = ?", id).Group("posts.id").Order("posts.id asc").Find(&thread).Error; err != nil {
		return nil, err
	}

	return thread, nil
}

// DeleteOldThreads deletes any thread older than the last bump date in page 9.
func DeleteOldThreads() error {
	// Make empty slice of IDs
	threadIDs := make([]uint64, 0)

	// Query ID from old threads and save them in the slice and check if there aren't IDs found
	if result := db.Model(&model.Post{}).Offset(100).Where("on_thread IS NULL").Order("updated_at desc").Pluck("id", &threadIDs); len(threadIDs) == 0 {
		// If there's none, return the error
		return result.Error
	}

	// Delete old threads as defined in the slice
	return db.Where("id IN (?)", threadIDs).Delete(&model.Post{}).Error
}

// BumpThread updates the post's "updated_at" field.
func BumpThread(id uint64, updatedAt *time.Time) error {
	// Make empty thread
	var threadLength uint64

	// Query posts that belong to a thread
	if err := db.Model(&model.Post{}).Where("id = ?", id).Or("on_thread = ?", id).Order("id asc").Count(&threadLength).Error; err != nil {
		return err
	}

	// If there are less than 300 posts in the thread, update it
	if threadLength < 300 {
		return db.Model(&model.Post{}).Where("id = ?", id).Update("updated_at", updatedAt).Error
	}

	return nil
}
