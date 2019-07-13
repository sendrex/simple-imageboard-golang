package handler

import (
	"time"

	"github.com/AquoDev/simple-imageboard-golang/database/methods"
	"github.com/AquoDev/simple-imageboard-golang/redis"
	"github.com/kataras/iris"
)

// GetPageExample handles a JSON response with a how-to example.
func GetPageExample(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"info":    "If you want to see any page from 0 to 9 (being 0 the first and 9 the last), see the example.",
		"example": "{url}/page/{number}",
	})
}

// GetPage handles a JSON response with a number of posts defined beforehand.
func GetPage(ctx iris.Context) {
	var response string
	var err error

	// Parse page ID
	id := ctx.Params().GetUint64Default("id", 0)
	redisKey := redis.GetPageKey(id)

	// Get page from cache
	response, err = redisClient.Get(redisKey).Result()
	// If it exists, return a response with it
	if err == nil {
		ctx.WriteString(response)
		return
	}

	// If it couldn't be found in cache, get it from the database
	response, err = methods.GetPage(id)
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
