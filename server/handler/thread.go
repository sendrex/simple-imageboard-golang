package handler

import (
	"encoding/json"
	"time"

	"github.com/AquoDev/simple-imageboard-golang/database"
	"github.com/AquoDev/simple-imageboard-golang/database/methods"
	"github.com/AquoDev/simple-imageboard-golang/redis"
	"github.com/kataras/iris"
)

// GetThreadExample handles a JSON response with a how-to example.
func GetThreadExample(ctx iris.Context) {
	ctx.JSON(map[string]string{
		"info":    "If you want to see any thread, see the example.",
		"example": "{url}/thread/{id}",
	})
}

// GetThread handles a JSON response with a number of posts defined beforehand.
func GetThread(ctx iris.Context) {
	var response interface{}
	var err error

	// Parse thread ID
	id := ctx.Params().GetUint64Default("id", 0)
	redisKey := redis.GetThreadKey(id)

	// Get thread from cache
	response, err = redis.Client().Get(redisKey).Result()
	// If it exists, return a response with it
	if err == nil {
		thread := make([]database.Post, 0)
		cachedThread := []byte(response.(string))
		json.Unmarshal(cachedThread, &thread)
		ctx.JSON(response)
		return
	}

	// If it couldn't be found in cache, get it from the database
	response, err = methods.GetThread(id)
	if err != nil {
		// If there are errors, the response will be 400 Bad Request
		response = GetError(400)
	} else if len(response.([]database.Post)) == 0 {
		// If the list is empty, the response will be 404 Not Found
		response = GetError(404)
	}

	// Set the cache
	cachedThread, _ := json.Marshal(response)
	redis.Client().Set(redisKey, string(cachedThread), 15*time.Second)
	ctx.JSON(response)
}
