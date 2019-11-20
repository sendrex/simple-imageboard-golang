package handler

import (
	"github.com/AquoDev/simple-imageboard-golang/database"
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
	// Parse thread ID
	id := ctx.Params().GetUint64Default("id", 0)

	// Get thread from cache
	if response, err := redis.GetCachedThread(id); err == nil {
		// If it exists, return a response with it
		ctx.JSON(response)
		return
	}

	// If it couldn't be found in cache, get it from the database
	if response, err := database.GetThread(id); err == nil && len(response) > 0 {
		// If the thread is not empty, set the cache and send the response
		redis.SetCachedThread(id, response)
		ctx.JSON(response)
		return
	}

	// At last, send 404 Not Found if the thread doesn't exist and set the cache
	response := GetError(404)
	redis.SetErrorCachedThread(id, response)
	ctx.JSON(response)
}
