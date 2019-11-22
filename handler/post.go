package handler

import (
	"net/http"
	"strconv"

	"github.com/AquoDev/simple-imageboard-golang/database"
	"github.com/AquoDev/simple-imageboard-golang/model"
	"github.com/AquoDev/simple-imageboard-golang/redis"
	"github.com/AquoDev/simple-imageboard-golang/util"
	"github.com/labstack/echo/v4"
)

// GetPostExample handles a JSON response with a how-to example.
func GetPostExample(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{
		"info":    "If you want to see any post, see the example.",
		"example": "{url}/post/{id}",
	})
}

// GetPost handles a JSON response with a post.
func GetPost(ctx echo.Context) error {
	// Parse post ID
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil || id == 0 {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	// Get post from cache
	if response, err := redis.GetCachedPost(id); err == nil {
		// If it exists, return a response with it
		if response.Status == http.StatusOK {
			return ctx.JSON(response.Status, response.Data)
		}
		// If it's an error, return an error response
		return echo.NewHTTPError(response.Status)
	}

	// If it couldn't be found in cache, get it from the database
	if response, err := database.GetPost(id); err == nil {
		// If it exists, set the cache and send the post
		go redis.SetCachedPost(id, http.StatusOK, response)
		return ctx.JSON(http.StatusOK, response)
	}

	// At last, send 404 Not Found if the post doesn't exist
	go redis.SetCachedPost(id, http.StatusNotFound, nil)
	return echo.NewHTTPError(http.StatusNotFound)
}

// SavePost handles a JSON response and saves the data as a post.
func SavePost(ctx echo.Context) error {
	post := new(model.Post)

	// Read JSON from body
	if err := ctx.Bind(&post); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Check if post is not valid
	if err := post.Check(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Make delete code if it hasn't one
	if post.DeleteCode == "" {
		post.DeleteCode = util.RandomString(32)
	}

	// Try to save the post (or thread) and check if it has been saved
	response, err := database.SavePost(post)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if post.OnThread == nil {
		// Delete old threads when the post starts a new thread
		go database.DeleteOldThreads()
	} else {
		// Bump thread if it hasn't reached bump limit
		go database.BumpThread(uint64(post.OnThread.ValueOrZero()), post.CreatedAt)
	}

	return ctx.JSON(http.StatusCreated, response)
}

// DeletePost handles a JSON response with a post.
func DeletePost(ctx echo.Context) error {
	data := new(model.DeleteData)

	// Read JSON from body
	if err := ctx.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Try to get post from database to check if it exists
	if _, err := database.GetPost(data.ID); err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	// Try to delete post (and thread if the post has "on_thread == null")
	if err := database.DeletePost(data.ID, data.DeleteCode); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": http.StatusText(http.StatusOK),
	})
}
