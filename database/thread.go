package database

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/AquoDev/simple-imageboard-golang/model"
)

var postsPerThread uint64

// GetThread returns a slice of posts.
func GetThread(id uint64) ([]model.Post, error) {
	// Make empty thread
	thread := make([]model.Post, 0)

	// Query posts that belong to a thread
	if err := db.Select("id, content, pic, reply_to, created_at, updated_at").Where("id = ?", id).Or("parent_thread = ?", id).Or("reply_to = ?", id).Order("id asc").Find(&thread).Error; err != nil {
		return nil, err
	}

	return thread, nil
}

// DeleteOldThreads deletes any thread older than the last bump date in page 9.
func DeleteOldThreads() error {
	// Make empty slice of IDs
	threadIDs := make([]uint64, 0)

	// Query IDs from old threads and save them in the slice and check if there aren't IDs found
	if result := db.Model(&model.Post{}).Offset(threadsPerPage*10).Where("parent_thread IS NULL").Order("updated_at desc").Pluck("id", &threadIDs); len(threadIDs) == 0 {
		// If there's none, return the error
		return result.Error
	}

	// Delete old threads as defined in the slice
	return db.Where("id IN (?)", threadIDs).Delete(&model.Post{}).Error
}

// BumpThread updates the post's "updated_at" field.
func BumpThread(id uint64, updatedAt *time.Time) error {
	var threadLength uint64

	// Query posts that belong to a thread
	if err := db.Model(&model.Post{}).Where("id = ?", id).Or("parent_thread = ?", id).Order("id asc").Count(&threadLength).Error; err != nil {
		return err
	}

	// If there are less than "max" posts in the thread, update it
	if threadLength < postsPerThread {
		return db.Model(&model.Post{}).Where("id = ?", id).Update("updated_at", updatedAt).Error
	}

	return nil
}

func init() {
	if parseUint, err := strconv.ParseUint(os.Getenv("POSTS_PER_THREAD"), 10, 0); err != nil {
		message := fmt.Errorf("[DATABASE] Couldn't parse POSTS_PER_THREAD @ %w", err)
		panic(message)
	} else {
		postsPerThread = parseUint
	}
}
