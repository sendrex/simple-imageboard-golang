package handler

import (
	"time"

	"github.com/AquoDev/simple-imageboard-golang/database"
	"github.com/AquoDev/simple-imageboard-golang/database/methods"
	"github.com/AquoDev/simple-imageboard-golang/redis"
	"github.com/AquoDev/simple-imageboard-golang/server/utils"
	"github.com/kataras/iris"
)

// DeleteData is a struct in which data to delete a post is parsed from JSON.
type DeleteData struct {
	DeleteCode string `json:"delete_code"`
}

// GetPostExample handles a JSON response with a how-to example.
func GetPostExample(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"info":    "If you want to see any post, see the example.",
		"example": "{url}/post/{id}",
	})
}

// GetPost handles a JSON response with a post.
func GetPost(ctx iris.Context) {
	var response string
	var err error

	// Parse post ID
	id := ctx.Params().GetUint64Default("id", 0)
	redisKey := redis.GetPostKey(id)

	// Get post from cache
	response, err = redisClient.Get(redisKey).Result()
	// If it exists, return a response with it
	if err == nil {
		ctx.WriteString(response)
		return
	}

	// If it couldn't be found in cache, get it from the database
	response, err = methods.GetPost(id)
	if err != nil {
		// If the post isn't found, the response will be 404 Not Found
		response = GetError(404)
	}

	// Set the cache and send the response (be it a correct or failed one)
	redisClient.Set(redisKey, response, 30*time.Second)
	ctx.WriteString(response)
}

// TODO test implementation for POST method, it's not finished
func SavePost(ctx iris.Context) {
	post := new(database.Post)

	// Read JSON from body
	if err := ctx.ReadJSON(&post); err != nil {
		invalidData := GetError(400)
		ctx.WriteString(invalidData)
		return
	}

	// Check if it has already an ID (it shouldn't have)
	if post.ID != 0 {
		invalidData := GetError(400)
		ctx.WriteString(invalidData)
		return
	}

	// Make delete code if it hasn't one
	if post.DeleteCode == "" {
		post.DeleteCode = utils.RandomString(32)
	}

	if response, err := methods.SavePost(post); err != nil {
		invalidData := GetError(400)
		ctx.WriteString(invalidData)
	} else {
		// TODO delete old threads when the post starts a new thread
		ctx.WriteString(response)
	}
}

// DeletePost handles a JSON response with a post.
func DeletePost(ctx iris.Context) {
	data := new(DeleteData)

	// Parse post ID
	id := ctx.Params().GetUint64Default("id", 0)

	// Try to get post from database to check if it exists
	if _, err := methods.GetPost(id); err != nil {
		postDoesntExist := GetError(404)
		ctx.WriteString(postDoesntExist)
		return
	}

	// Read JSON from body
	if err := ctx.ReadJSON(&data); err != nil {
		invalidData := GetError(400)
		ctx.WriteString(invalidData)
		return
	}

	// Try to delete post (and thread if the post has "on_thread == null")
	if err := methods.DeletePost(id, data.DeleteCode); err != nil {
		incorrectCode := GetError(400)
		ctx.WriteString(incorrectCode)
		return
	}

	ctx.JSON(iris.Map{
		"status":  200,
		"message": "Success",
	})
}
