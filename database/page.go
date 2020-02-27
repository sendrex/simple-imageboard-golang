package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AquoDev/simple-imageboard-golang/model"
)

var threadsPerPage uint64

// GetPage returns a thread list (post slice).
func GetPage(id uint64) ([]model.Post, error) {
	// Make empty page
	page := make([]model.Post, 0)

	// Query posts that started a thread (on_thread == null)
	if err := db.Select("posts.id, posts.content, posts.pic, posts.created_at, posts.updated_at, count(replies.on_thread) as replies").Joins("LEFT JOIN posts replies ON replies.on_thread = posts.id").Offset(threadsPerPage * id).Limit(threadsPerPage).Where("posts.on_thread IS NULL").Group("posts.id").Order("posts.updated_at desc").Find(&page).Error; err != nil {
		return nil, err
	}

	return page, nil
}

func init() {
	if parseUint, err := strconv.ParseUint(os.Getenv("THREADS_PER_PAGE"), 10, 0); err != nil {
		message := fmt.Errorf("[DATABASE] Couldn't parse THREADS_PER_PAGE @ %w", err)
		panic(message)
	} else {
		threadsPerPage = parseUint
	}
}
