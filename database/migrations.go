package database

import (
	"fmt"

	"github.com/AquoDev/simple-imageboard-golang/model"
)

func init() {
	// Make migration
	if err := db.AutoMigrate(&model.Post{}).Error; err != nil {
		message := fmt.Errorf("[DATABASE] Migrations FAILED @ %w", err)
		panic(message)
	}

	// "parent_thread" should be foreign key
	if err := db.Model(&model.Post{}).AddForeignKey("parent_thread", "posts(id)", "CASCADE", "RESTRICT").Error; err != nil {
		message := fmt.Errorf("[DATABASE] Migrations FAILED @ %w", err)
		panic(message)
	}

	// "reply_to" should be foreign key
	if err := db.Model(&model.Post{}).AddForeignKey("reply_to", "posts(id)", "SET NULL", "RESTRICT").Error; err != nil {
		message := fmt.Errorf("[DATABASE] Migrations FAILED @ %w", err)
		panic(message)
	}
}
