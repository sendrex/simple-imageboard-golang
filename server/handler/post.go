package handler

import (
	"time"

	"github.com/AquoDev/simple-imageboard-golang/database/methods"
	"github.com/AquoDev/simple-imageboard-golang/redis"
	"github.com/kataras/iris"
)

// TODO implement SavePost, DeletePost funcs

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
