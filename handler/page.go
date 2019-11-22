package handler

import (
	"net/http"
	"strconv"

	"github.com/AquoDev/simple-imageboard-golang/database"
	"github.com/AquoDev/simple-imageboard-golang/redis"
	"github.com/labstack/echo/v4"
)

// GetPageExample handles a JSON response with a how-to example.
func GetPageExample(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{
		"info":    "If you want to see any page from 0 to 9 (being 0 the first and 9 the last), see the example.",
		"example": "{url}/page/{number}",
	})
}

// GetPage handles a JSON response with a number of posts defined beforehand.
func GetPage(ctx echo.Context) error {
	// Parse page ID
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil || id > 9 {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	// Get page from cache
	if response, err := redis.GetCachedPage(id); err == nil {
		// If it exists, return a response with it
		if response.Status == http.StatusOK {
			return ctx.JSON(response.Status, response.Data)
		}
		// If it's an error, return an error response
		return echo.NewHTTPError(response.Status)
	}

	// If it couldn't be found in cache, get it from the database
	if response, err := database.GetPage(id); err == nil {
		// Even if the page is empty, set the cache and send the response
		go redis.SetCachedPage(id, http.StatusOK, response)
		return ctx.JSON(http.StatusOK, response)
	}

	// At last, send 500 Internal Server error
	go redis.SetCachedPage(id, http.StatusInternalServerError, nil)
	return echo.NewHTTPError(http.StatusInternalServerError)
}
