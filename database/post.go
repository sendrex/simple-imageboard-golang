package database

import "time"

// Post defines the table in which the posts will be saved.
type Post struct {
	ID         uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	Content    string `gorm:"not null;size:1000"`
	OnThread   uint64
	ReplyTo    uint64
	Pic        string    `gorm:"not null;size:512"`
	DeleteCode string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}

func init() {
	// Make migration
	db.AutoMigrate(&Post{})
}
