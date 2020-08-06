package model

import (
	"errors"
	"math/rand"
	"regexp"
	"time"

	"github.com/AquoDev/simple-imageboard-golang/env"
	"github.com/jinzhu/gorm"
)

var (
	tableName = env.GetString("DB_TABLE_NAME")
	regex     = regexp.MustCompile(`^(http(s?):)([/|.|\w|\s|-])*\.(?:jp(e?)g|gif|png|webm|webp)$`)
)

// Post is the main struct. It defines how posts will be stored and represented.
type Post struct {
	ID           uint64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Content      string     `json:"content" gorm:"not null;size:1000"`
	Password     string     `json:"password,omitempty" gorm:"not null;size:128"`
	Pic          *string    `json:"pic,omitempty" gorm:"size:512"`
	ParentThread *uint64    `json:"parent_thread,omitempty"`
	ReplyTo      *uint64    `json:"reply_to,omitempty"`
	CreatedAt    *time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty" gorm:"not null;default:CURRENT_TIMESTAMP"`
	Sticky       bool       `json:"sticky,omitempty"`
	Closed       bool       `json:"closed,omitempty"`
	Sage         bool       `json:"sage,omitempty" gorm:"-"` // This field is not stored
}

// TableName sets the table name for the model Post.
func (post *Post) TableName() string {
	return tableName
}

// BeforeCreate makes the delete code if it's empty.
func (post *Post) BeforeCreate(scope *gorm.Scope) error {
	// Make delete code if it hasn't one
	if post.Password == "" {
		post.Password = randomString(32)
	}

	return nil
}

// Validate checks if all fields are valid.
func (post *Post) Validate() error {
	if post.ID != 0 {
		// Check if post hasn't got an ID
		return errors.New("id must not be set")
	} else if len(post.Content) == 0 {
		// Check if content has at least one character
		return errors.New("content must not be empty")
	} else if len(post.Content) > 1000 {
		// Check if content hasn't more than 1000 characters
		return errors.New("content must not have more than 1000 characters")
	} else if len(post.Password) > 128 {
		// Check if password hasn't more than 128 characters
		return errors.New("password must not have more than 128 characters")
	} else if post.Pic != nil {
		// Check if pic is not empty and...
		if len(*post.Pic) > 512 {
			// ... it hasn't more than 512 characters
			return errors.New("pic must not have more than 512 characters")
		} else if !regex.MatchString(*post.Pic) {
			// ... it's a valid URL
			return errors.New("pic must be a valid URL")
		}
	} else if post.Sticky {
		// Check if sticky is set to false
		return errors.New("sticky must not be set to true")
	} else if post.Closed {
		// Check if closed is set to false
		return errors.New("closed must not be set to true")
	}

	return nil
}

// RepliesToAnotherPost checks if this post replies to another post.
func (post *Post) RepliesToAnotherPost() bool {
	return post.ReplyTo != nil
}

// IsAParentThread checks if this post starts a parent thread.
func (post *Post) IsAParentThread() bool {
	return post.ParentThread == nil
}

// AllowsReplies checks if this post allows being replied.
func (post *Post) AllowsReplies() bool {
	return !post.Closed
}

// Sages checks if parent thread must not be bumped.
func (post *Post) Sages() bool {
	return post.Sage
}

// MakeRelationWith adds the parent thread ID to the post based on the reply.
func (post *Post) MakeRelationWith(reply *Post) {
	if reply.IsAParentThread() {
		post.ParentThread = &reply.ID
	} else {
		post.ParentThread = reply.ParentThread
	}
}

// randomString returns a random string from predefined characters.
func randomString(length int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, length)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}
