package database

import "time"

// Post defines the table in which the posts will be saved.
type Post struct {
	ID         uint64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Content    string     `json:"content" gorm:"not null;size:1000"`
	OnThread   uint64     `json:"on_thread,omitempty"`
	ReplyTo    uint64     `json:"reply_to,omitempty"`
	Pic        string     `json:"pic" gorm:"size:512"`
	DeleteCode string     `json:"delete_code,omitempty" gorm:"not null"`
	CreatedAt  *time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" gorm:"not null;default:CURRENT_TIMESTAMP"`
}

func init() {
	// Make migration
	db.AutoMigrate(&Post{})
}
