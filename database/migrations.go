package database

import (
	"fmt"

	"github.com/AquoDev/simple-imageboard-golang/model"
)

func init() {
	// Make migration
	if err := db.AutoMigrate(&model.Post{}).Error; err != nil {
		panic(err)
	}

	// "on_thread" should be foreign key
	if err := db.Model(&model.Post{}).AddForeignKey("on_thread", "posts(id)", "CASCADE", "RESTRICT").Error; err != nil {
		panic(err)
	}

	fmt.Println("[DATABASE]: Migrations OK")
}
