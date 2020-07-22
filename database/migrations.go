package database

import (
	"fmt"

	"github.com/AquoDev/simple-imageboard-golang/model"
)

func init() {
	tableName := (&model.Post{}).TableName()

	// Create table and update it with new fields if it's needed
	if err := db.AutoMigrate(&model.Post{}).Error; err != nil {
		message := fmt.Errorf("[DATABASE] Automigrate() failed @ %w", err)
		panic(message)
	}

	// "parent_thread" should be foreign key
	if err := db.Model(&model.Post{}).AddForeignKey("parent_thread", fmt.Sprintf("%s(id)", tableName), "CASCADE", "RESTRICT").Error; err != nil {
		message := fmt.Errorf("[DATABASE] AddForeignKey(\"parent_thread\") failed @ %w", err)
		panic(message)
	}

	// "reply_to" should be foreign key
	if err := db.Model(&model.Post{}).AddForeignKey("reply_to", fmt.Sprintf("%s(id)", tableName), "SET NULL", "RESTRICT").Error; err != nil {
		message := fmt.Errorf("[DATABASE] AddForeignKey(\"reply_to\") failed @ %w", err)
		panic(message)
	}
}
