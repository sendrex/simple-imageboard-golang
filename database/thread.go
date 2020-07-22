package database

import (
	"errors"

	"github.com/AquoDev/simple-imageboard-golang/env"
	"github.com/AquoDev/simple-imageboard-golang/model"
)

var (
	postsPerThread   = env.GetUint64("POSTS_PER_THREAD")
	maxParentThreads = env.GetUint64("MAX_PARENT_THREADS")
)

// GetThread returns every post that belong to a thread.
func GetThread(id uint64) ([]model.Post, error) {
	// Query every post that belong to a thread given its ID
	thread := make([]model.Post, 0)
	err := db.Select("id, content, pic, reply_to, created_at").Where("id = ?", id).Or("parent_thread = ?", id).Or("reply_to = ?", id).Order("id asc").Find(&thread).Error

	if err != nil {
		// If there's any error, return it
		return nil, err
	} else if len(thread) == 0 {
		// If it's empty, return a new error
		return nil, errors.New("thread doesn't exist")
	}

	// Return thread without error
	return thread, nil
}

// DeleteOldThreads deletes any thread that's fallen from the index.
func DeleteOldThreads() error {
	// Query IDs from old threads
	threadIDs := make([]uint64, 0)
	err := db.Model(&model.Post{}).Offset(maxParentThreads).Where("parent_thread IS NULL").Order("updated_at desc").Pluck("id", &threadIDs).Error

	if err != nil {
		// If there's any error, return it
		return err
	} else if len(threadIDs) > 0 {
		// If there are threads to delete, delete them
		return db.Where("id IN (?)", threadIDs).Delete(&model.Post{}).Error
	}

	// No threads to delete
	return nil
}

// BumpThread updates the parent thread from a given post.
func BumpThread(post *model.Post) error {
	// Query count of posts that belong to a thread given the parent post ID from a post
	var threadLength uint64
	err := db.Model(&model.Post{}).Where("id = ?", *post.ParentThread).Or("parent_thread = ?", *post.ParentThread).Order("id asc").Count(&threadLength).Error

	if err != nil {
		// If there's any error, return it
		return err
	} else if threadLength < postsPerThread {
		// If there are less than "max" posts in the thread, update it
		return db.Model(&model.Post{}).Where("id = ?", *post.ParentThread).Update("updated_at", post.UpdatedAt).Error
	}

	// Thread is not bumped because it reached bump limit
	return nil
}
