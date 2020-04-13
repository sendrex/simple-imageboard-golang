package handler

import (
	"net/http"
	"strconv"

	"github.com/AquoDev/simple-imageboard-golang/database"
	"github.com/AquoDev/simple-imageboard-golang/model"
	"github.com/AquoDev/simple-imageboard-golang/redis"
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
	// Parse post ID. Is the parsing invalid?
	// ── Yes:	return a failed JSON response.
	// ── No:	continue. ID was parsed successfully.
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil || id == 0 {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	// Try to get the post from cache. Is it stored in cache?
	// ── Yes:	check if what's cached is an error or a post. Both options lead to a response.
	// ── No:	continue. It's not in cache, but it could be stored in the database.
	if response, err := redis.GetCachedPost(id); err == nil {
		// Is the post cached?
		// ── Yes:	return the cached post.
		// ── No:	it's an error, so return that error as a failed JSON response.
		if response.Status == http.StatusOK {
			return ctx.JSON(response.Status, response.Data)
		}
		return echo.NewHTTPError(response.Status)
	}

	// Try to get the post from the database. Does it exist?
	// ── Yes:	return the post.
	// ── No:	continue. The post doesn't exist.
	if response, err := database.GetPost(id); err == nil {
		go redis.SetCachedPost(id, http.StatusOK, response)
		return ctx.JSON(http.StatusOK, response)
	}

	// Return a 404 NotFound JSON response after caching it.
	go redis.SetCachedPost(id, http.StatusNotFound, nil)
	return echo.NewHTTPError(http.StatusNotFound)
}

// SavePost handles a JSON response and saves the data as a post.
func SavePost(ctx echo.Context) error {
	// Read body (JSON) from the request. Is the body request bad formed?
	// ── Yes:	return a failed JSON response.
	// ── No:	continue. Body is well formed and could be read.
	post := new(model.Post)
	if err := ctx.Bind(&post); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Does the post we receive reply to another post?
	// ── Yes:	check if the reply exists. We can't reply to a post that doesn't exists (foreign key constrain).
	// ── No:	continue. The post we receive starts a parent thread.
	if post.ReplyTo != nil {
		// Does the reply exists?
		// ── Yes:	check if the reply is a regular reply or a parent thread.
		// ── No:	return a failed JSON response.
		if reply, err := database.GetPost(*post.ReplyTo); err == nil {
			// Is the reply a regular post that doesn't start a thread?
			// ── Yes:	the post we receive will belong to the same parent thread as the reply.
			// ── No:	the reply started a parent thread, so we copy the reply ID to the parent thread ID
			//			as they're referencing the same post.
			if reply.ParentThread != nil {
				post.ParentThread = reply.ParentThread
			} else {
				post.ParentThread = post.ReplyTo
			}
		} else {
			return echo.NewHTTPError(http.StatusForbidden)
		}
	}

	// Try to save the post in the database. Does it fail to be saved?
	// ── Yes:	return a failed JSON response.
	// ── No:	continue. Body is well formed and could be read.
	response, err := database.SavePost(post)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Does the post belong to a parent thread?
	// ── Yes:	bump that parent thread.
	// ── No:	this post starts a parent thread, so threads that are beyond the last page must be deleted.
	if post.ParentThread != nil {
		go database.BumpThread(*post.ParentThread, post.CreatedAt)
	} else {
		go database.DeleteOldThreads()
	}

	// Return a 201 Created JSON response with the post ID and its delete code.
	return ctx.JSON(http.StatusCreated, response)
}

// DeletePost handles a JSON response with a post.
func DeletePost(ctx echo.Context) error {
	// Read body (JSON) from the request. Is the body request bad formed?
	// ── Yes:	return a failed JSON response. Maybe there's some field missing.
	// ── No:	continue. Body is well formed and could be read.
	data := new(model.DeleteData)
	if err := ctx.Bind(&data); err != nil || data.ID == 0 || data.DeleteCode == "" {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Try to get the post that's being deleted. Isn't it found?
	// ── Yes:	return a failed JSON response.
	// ── No:	continue. Post exists and can be deleted.
	if _, err := database.GetPost(data.ID); err != nil {
		return echo.NewHTTPError(http.StatusForbidden)
	}

	// Try to delete post (and thread if the post has "parent_thread == nil"). Couldn't it be deleted?
	// ── Yes:	return a failed JSON response. Incorrect delete code.
	// ── No:	continue. Correct delete code, the post was deleted.
	if err := database.DeletePost(data); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	// Return a 200 OK JSON response.
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": http.StatusText(http.StatusOK),
	})
}
