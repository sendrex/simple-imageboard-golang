package database

import (
	"github.com/AquoDev/simple-imageboard-golang/model"
)

// GetPage returns a thread list (post slice).
func GetPage(id uint64) ([]model.Post, error) {
	// Make empty page
	page := make([]model.Post, 0)

	// Query posts that started a thread (on_thread == null)
	if err := db.Select("posts.id, posts.content, posts.pic, posts.created_at, posts.updated_at, count(b.on_thread) as replies").Joins("LEFT JOIN posts b ON b.on_thread = posts.id").Offset(threadsPerPage * id).Limit(threadsPerPage).Where("posts.on_thread IS NULL").Group("posts.id").Order("posts.updated_at desc").Find(&page).Error; err != nil {
		return nil, err
	}

	return page, nil
}
