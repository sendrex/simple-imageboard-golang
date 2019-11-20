package database

import (
	"fmt"

	"github.com/AquoDev/simple-imageboard-golang/model"
)

func init() {
	// Make migration
	db.AutoMigrate(&model.Post{})

	// "on_thread" should be foreign key
	db.Model(&model.Post{}).AddForeignKey("on_thread", "posts(id)", "CASCADE", "RESTRICT")

	fmt.Println("Migrations are done")
}
