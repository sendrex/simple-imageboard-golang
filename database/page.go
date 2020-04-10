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

	// Query posts that started a thread (parent_thread IS NULL)
	if err := db.Select("id, content, pic, created_at, updated_at").Offset(threadsPerPage * id).Limit(threadsPerPage).Where("parent_thread IS NULL").Order("updated_at desc").Find(&page).Error; err != nil {
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
