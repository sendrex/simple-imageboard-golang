package handler

import (
	"github.com/AquoDev/simple-imageboard-golang/database"
	"github.com/AquoDev/simple-imageboard-golang/redis"
	"github.com/kataras/iris/v12"
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
	// Parse page ID
	id := ctx.Params().GetUint64Default("id", 0)

	// Get page from cache
	if response, err := redis.GetCachedPage(id); err == nil {
		// If it exists, return a response with it
		ctx.JSON(response)
		return
	}

	// If it couldn't be found in cache, get it from the database
	if response, err := database.GetPage(id); err == nil {
		// Even if the page is empty, set the cache and send the response
		redis.SetCachedPage(id, response)
		ctx.JSON(response)
		return
	}

	// At last, send 500 Internal Server error and set the cache
	response := GetError(500)
	redis.SetCachedPage(id, response)
	ctx.JSON(response)
}
