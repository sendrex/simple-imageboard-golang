package handler

import (
	"time"

	"github.com/AquoDev/simple-imageboard-golang/database/methods"
	"github.com/AquoDev/simple-imageboard-golang/redis"
	"github.com/kataras/iris"
)

// GetPageExample handles a JSON response with a how-to example.
func GetPageExample(ctx iris.Context) {
	ctx.JSON(map[string]string{
		"info":    "If you want to see any page from 0 to 9 (being 0 the first and 9 the last), see the example.",
		"example": "{url}/page/{number}",
	})
}

// GetPage handles a JSON response with a number of posts defined beforehand.
func GetPage(ctx iris.Context) {
	var response interface{}
	var err error

	// Parse page ID
	id := ctx.Params().GetUint64Default("id", 0)
	redisKey := redis.GetPageKey(id)

	// Get page from cache
	response, err = redis.Client().Get(redisKey).Result()
	// If it exists, return a response with it
	if err == nil {
		ctx.JSON(response)
		return
	}

	// If it couldn't be found in cache, get it from the database
	response, err = methods.GetPage(id)
	if err != nil {
		// If there are errors, the response will be 400 Bad Request
		response = GetError(400)
	}

	// Set the cache and send the response (it could be empty or not)
	redis.Client().Set(redisKey, response, 15*time.Second)
	ctx.JSON(response)
}
