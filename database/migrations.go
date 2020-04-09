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

	// "parent_post" should be foreign key
	if err := db.Model(&model.Post{}).AddForeignKey("parent_post", "posts(id)", "CASCADE", "RESTRICT").Error; err != nil {
		message := fmt.Errorf("[DATABASE] Migrations FAILED @ %w", err)
		panic(message)
	}

	// "reply_to" should be foreign key
	if err := db.Model(&model.Post{}).AddForeignKey("reply_to", "posts(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		message := fmt.Errorf("[DATABASE] Migrations FAILED @ %w", err)
		panic(message)
	}

	fmt.Println("[DATABASE] Migrations OK")
}
