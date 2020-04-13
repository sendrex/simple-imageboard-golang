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
	// Parse page number. Is the parsing invalid?
	// ── Yes:	return a failed JSON response.
	// ── No:	continue. Page number was parsed successfully.
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil || id > 9 {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	// Try to get the page from cache. Is it stored in cache?
	// ── Yes:	check if what's cached is an error or a page. Both options lead to a response.
	// ── No:	continue.
	if response, err := redis.GetCachedPage(id); err == nil {
		// Is the page cached?
		// ── Yes:	return the cached page.
		// ── No:	it's an error, so return that error as a failed JSON response.
		if response.Status == http.StatusOK {
			return ctx.JSON(response.Status, response.Data)
		}
		return echo.NewHTTPError(response.Status)
	}

	// Try to get the parent threads from the database, even if the slice we get is empty. Did it go correctly?
	// ── Yes:	cache the page and return it, even if it's empty.
	// ── No:	continue. There must be a server-side error. This means something has gone seriously wrong.
	if response, err := database.GetPage(id); err == nil {
		go redis.SetCachedPage(id, http.StatusOK, response)
		return ctx.JSON(http.StatusOK, response)
	}

	// Return a 500 InternalServerError JSON response after caching it.
	go redis.SetCachedPage(id, http.StatusInternalServerError, nil)
	return echo.NewHTTPError(http.StatusInternalServerError)
}
