package handler

import (
	"net/http"
	"strconv"

	"github.com/AquoDev/simple-imageboard-golang/database"
	"github.com/AquoDev/simple-imageboard-golang/redis"
	"github.com/labstack/echo/v4"
)

// GetThreadExample handles a JSON response with a how-to example.
func GetThreadExample(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{
		"info":    "If you want to see any thread, see the example.",
		"example": "{url}/thread/{id}",
	})
}

// GetThread handles a JSON response with a number of posts defined beforehand.
func GetThread(ctx echo.Context) error {
	// Parse thread ID
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil || id == 0 {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	// Get thread from cache
	if response, err := redis.GetCachedThread(id); err == nil {
		// If it exists, return a response with it
		return ctx.JSON(http.StatusOK, response)
	}

	// If it couldn't be found in cache, get it from the database
	if response, err := database.GetThread(id); err == nil && len(response) > 0 {
		// If the thread is not empty, set the cache and send the response
		redis.SetCachedThread(id, response)
		return ctx.JSON(http.StatusOK, response)
	}

	// At last, send 404 Not Found if the thread doesn't exist
	return echo.NewHTTPError(http.StatusNotFound)
}
