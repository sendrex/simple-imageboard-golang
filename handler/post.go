package handler

import (
	"github.com/AquoDev/simple-imageboard-golang/database"
	"github.com/AquoDev/simple-imageboard-golang/model"
	"github.com/AquoDev/simple-imageboard-golang/redis"
	"github.com/AquoDev/simple-imageboard-golang/util"
	"github.com/kataras/iris/v12"
)

// GetPostExample handles a JSON response with a how-to example.
func GetPostExample(ctx iris.Context) {
	ctx.JSON(map[string]string{
		"info":    "If you want to see any post, see the example.",
		"example": "{url}/post/{id}",
	})
}

// GetPost handles a JSON response with a post.
func GetPost(ctx iris.Context) {
	// Parse post ID
	id := ctx.Params().GetUint64Default("id", 0)

	// Get post from cache
	if response, err := redis.GetCachedPost(id); err == nil {
		// If it exists, return a response with it
		ctx.JSON(response)
		return
	}

	// If it couldn't be found in cache, get it from the database
	if response, err := database.GetPost(id); err == nil {
		// If it exists, set the cache and send the post
		redis.SetCachedPost(id, response)
		ctx.JSON(response)
		return
	}

	// At last, send 404 Not Found if the post doesn't exist and set the cache
	response := GetError(404)
	redis.SetCachedPost(id, response)
	ctx.JSON(response)
}

// SavePost handles a JSON response and saves the data as a post.
func SavePost(ctx iris.Context) {
	post := new(model.Post)

	// Read JSON from body
	if err := ctx.ReadJSON(&post); err != nil {
		invalidData := GetError(400)
		ctx.JSON(invalidData)
		return
	}

	// Check if post is not valid
	if err := post.Check(); err != nil {
		invalidData := GetError(400)
		ctx.JSON(invalidData)
		return
	}

	// Make delete code if it hasn't one
	if post.DeleteCode == "" {
		post.DeleteCode = util.RandomString(32)
	}

	// Try to save the post (or thread) and check if it has been saved
	if response, err := database.SavePost(post); err != nil {
		invalidData := GetError(400)
		ctx.JSON(invalidData)
	} else {
		if post.OnThread == nil {
			// Delete old threads when the post starts a new thread
			database.DeleteOldThreads()
		} else {
			// Bump thread if it hasn't reached bump limit
			database.BumpThread(uint64(post.OnThread.ValueOrZero()), post.CreatedAt)
		}
		ctx.JSON(response)
	}
}

// DeletePost handles a JSON response with a post.
func DeletePost(ctx iris.Context) {
	data := new(model.DeleteData)

	// Read JSON from body
	if err := ctx.ReadJSON(&data); err != nil {
		invalidData := GetError(400)
		ctx.JSON(invalidData)
		return
	}

	// Try to get post from database to check if it exists
	if _, err := database.GetPost(data.ID); err != nil {
		postDoesntExist := GetError(404)
		ctx.JSON(postDoesntExist)
		return
	}

	// Try to delete post (and thread if the post has "on_thread == null")
	if err := database.DeletePost(data.ID, data.DeleteCode); err != nil {
		incorrectCode := GetError(400)
		ctx.JSON(incorrectCode)
		return
	}

	ctx.JSON(map[string]interface{}{
		"message": "success",
		"status":  200,
	})
}
