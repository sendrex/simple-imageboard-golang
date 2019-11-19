package database

import (
	"gopkg.in/guregu/null.v3"
	"time"
)

// Post defines the table in which the posts will be saved and how it's represented in JSON.
type Post struct {
	ID         uint64       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Content    string       `json:"content" gorm:"not null;size:1000"`
	OnThread   *null.Int    `json:"on_thread,omitempty"`
	Pic        *null.String `json:"pic,omitempty" gorm:"size:512"`
	DeleteCode string       `json:"delete_code,omitempty" gorm:"not null;size:128"`
	CreatedAt  *time.Time   `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt  *time.Time   `json:"updated_at,omitempty" gorm:"not null;default:CURRENT_TIMESTAMP"`
}

func init() {
	// Make migration
	connection.AutoMigrate(&Post{})

	// "on_thread" should be foreign key
	connection.Model(&Post{}).AddForeignKey("on_thread", "posts(id)", "CASCADE", "RESTRICT")
}
