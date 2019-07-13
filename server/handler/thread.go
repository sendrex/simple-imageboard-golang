package handler

import (
	"time"

	"github.com/AquoDev/simple-imageboard-golang/database/methods"
	"github.com/AquoDev/simple-imageboard-golang/redis"
	"github.com/kataras/iris"
)

// GetThreadExample handles a JSON response with a how-to example.
func GetThreadExample(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"info":    "If you want to see any thread, see the example.",
		"example": "{url}/thread/{id}",
	})
}

// GetThread handles a JSON response with a number of posts defined beforehand.
func GetThread(ctx iris.Context) {
	var response string
	var err error

	// Parse page ID
	id := ctx.Params().GetUint64Default("id", 0)
	redisKey := redis.GetThreadKey(id)

	// Get page from cache
	response, err = redisClient.Get(redisKey).Result()
	// If it exists, return a response with it
	if err == nil {
		ctx.WriteString(response)
		return
	}

	// If it couldn't be found in cache, get it from the database
	response, err = methods.GetThread(id)
	if err != nil {
		// If there are errors, the response will be 400 Bad Request
		response = GetError(400)
	} else if response == "[]" {
		// If the list is empty, the response will be 404 Not Found
		response = GetError(404)
	}

	// Set the cache and send the response (be it a correct or failed one)
	redisClient.Set(redisKey, response, 15*time.Second)
	ctx.WriteString(response)
}
