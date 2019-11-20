package database

import (
	"github.com/AquoDev/simple-imageboard-golang/model"
)

// GetPage returns a thread list (post slice).
func GetPage(id uint64) (page []model.Post, err error) {
	// Query posts that started a thread (on_thread == null)
	err = db.Select("id, content, pic, created_at, updated_at").Offset(10 * id).Limit(10).Where("on_thread IS NULL").Order("updated_at desc").Find(&page).Error
	return
}
