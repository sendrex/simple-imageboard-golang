package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"gopkg.in/guregu/null.v3"
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
	Replies    uint64       `json:"replies,omitempty"` // omitempty: "replies" is NOT shown if == 0
}

// Check returns an error if this post isn't valid, and nil otherwise.
func (post *Post) Check() error {
	// Check if content has at least one character
	if len(post.Content) == 0 {
		return errors.New("content must have at least 1 character")
	}

	// Check if content has surpassed 1000 characters
	if len(post.Content) > 1000 {
		return errors.New("content must not have more than 1000 characters")
	}

	// Check if delete_code has surpassed 128 characters
	if len(post.DeleteCode) > 128 {
		return errors.New("delete_code must not have more than 128 characters")
	}

	// Check if it has already an ID (it shouldn't have)
	if post.ID != 0 {
		return errors.New("id shouldn't be set")
	}

	// Check if pic exists and it's an invalid URL (it should be a valid URL)
	if post.Pic != nil && len(post.Pic.String) > 512 {
		return errors.New("pic must not have more than 512 characters")
	}

	// Check if pic exists and it's an invalid URL (it should be a valid URL)
	if post.Pic != nil && !govalidator.IsURL(post.Pic.String) {
		return errors.New("pic is a invalid url")
	}

	return nil
}
