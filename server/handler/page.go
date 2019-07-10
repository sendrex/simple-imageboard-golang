package handler

import (
	"time"

	"github.com/AquoDev/simple-imageboard-golang/redis"
	"github.com/kataras/iris"
)

// Redis client for getting from/setting to cache.
var client = redis.GetRedisClient()

// GetPageExample handles a JSON response with a how-to example.
func GetPageExample(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"info":    "If you want to see any page from 0 to 9 (being 0 the first and 9 the last), see the example.",
		"example": "{url}/page/0",
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
	response, err = client.Get(redisKey).Result()
	// If it exists, return a response with it
	if err == nil {
		ctx.WriteString(response)
		return
	}

	// If it couldn't be found in cache, get it from the database
	// TODO implement database logic
	// TODO look for an ORM (migrations, maybe seeding, etc)
	//response, err = database.GetPage(id)
	// If it doesn't exist, the response will be 404
	if err != nil {
		response = GetNotFoundResponse()
	}

	// Set the cache and send the response (be it a regular one or 404)
	client.Set(redisKey, response, 15*time.Second)
	ctx.WriteString(response)
}