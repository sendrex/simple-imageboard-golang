package database

import (
	"errors"
	"fmt"

	"github.com/AquoDev/simple-imageboard-golang/env"
	"github.com/AquoDev/simple-imageboard-golang/model"
)

var adminPassword = env.GetString("ADMIN_PASSWORD")

// GetPost returns a post given its ID.
func GetPost(id uint64) (*model.Post, error) {
	// Query post given its ID
	post := new(model.Post)
	err := db.Select("id, content, pic, parent_thread, reply_to, created_at, updated_at, sticky, closed").Where("id = ?", id).First(&post).Error

	// Return post and error
	return post, err
}

// SavePost returns a struct with the ID and delete code of the inserted post.
func SavePost(post *model.Post) (*model.DeleteData, error) {
	// Try to insert the post
	if err := db.Create(&post).Error; err != nil {
		return nil, err
	}

	// If it's been inserted, return delete data without error
	return &model.DeleteData{
		ID:       post.ID,
		Password: post.Password,
	}, nil
}

// DeletePost returns an error that should be checked in the handler.
// Warning: if the post started a thread (parent_thread == nil), it will delete
// every post in the thread (parent_thread == id).
func DeletePost(data *model.DeleteData) error {
	// Query post and check if the data is valid
	post := new(model.Post)
	query := "id = ?"
	args := []interface{}{data.ID}

	if data.Password != adminPassword {
		query = fmt.Sprintf("%s AND password = ?", query)
		args = append(args, data.Password)
	}

	err := db.Where(query, args...).First(&post).Error

	// If there's any error, return it
	if err != nil {
		return err
	}

	// Delete the post
	return db.Delete(&post).Error
}

// UpdatePost returns an error that should be checked in the handler.
// Warning: the post must start a thread (parent_thread == nil), or else this will fail.
func UpdatePost(data *model.UpdateData) error {
	// Check if password is the admin password
	if data.Password != adminPassword {
		return errors.New("password is not admin password")
	}

	// Query post from the database and check if it's a parent thread
	post := new(model.Post)
	err := db.Select("id, parent_thread").Where("id = ?", data.ID).First(&post).Error
	if err != nil {
		return err
	} else if !post.IsAParentThread() {
		return errors.New("post is not a parent thread")
	}

	// Update the post
	return db.Model(&post).Updates(map[string]interface{}{"sticky": data.Sticky, "closed": data.Closed}).Error
}
