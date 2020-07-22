package handler

import (
	"net/http"

	"github.com/AquoDev/simple-imageboard-golang/database"
	"github.com/AquoDev/simple-imageboard-golang/framework"
	"github.com/AquoDev/simple-imageboard-golang/redis"
)

var exampleThread = map[string]string{
	"info":    "If you want to see any thread, see the example.",
	"example": "{url}/thread/{id}",
}

// GetThreadExample handles a JSON response with a how-to example.
func GetThreadExample(ctx framework.Context) error {
	return framework.SendOK(ctx, exampleThread)
}

// GetThread handles a JSON response with a number of posts defined beforehand.
func GetThread(ctx framework.Context) error {
	// Parse thread ID. Is the parsing invalid?
	// ── Yes:	return a failed JSON response.
	// ── No:	continue. ID was parsed successfully.
	id, err := framework.GetID(ctx)
	if err != nil {
		return framework.SendError(http.StatusNotFound)
	}

	// Try to get the thread from cache. Is it stored in cache?
	// ── Yes:	check if what's cached is an error or a thread. Both options lead to a response.
	// ── No:	continue. It's not in cache, but posts could be stored in the database.
	if response, err := redis.GetCachedThread(id); err == nil {
		// Is the thread cached?
		// ── Yes:	return the cached thread.
		// ── No:	it's an error, so return that error as a failed JSON response.
		if response.Status == http.StatusOK {
			return framework.SendOK(ctx, response.Data)
		}
		return framework.SendError(response.Status)
	}

	// Try to get the posts from the database. Isn't the thread empty?
	// ── Yes:	cache the thread and return it.
	// ── No:	continue. The thread doesn't exist (if it's empty, it doesn't have the OP).
	if response, err := database.GetThread(id); err == nil {
		go redis.SetCachedThread(id, http.StatusOK, response)
		return framework.SendOK(ctx, response)
	}

	// Return a 404 NotFound JSON response after caching it.
	go redis.SetCachedThread(id, http.StatusNotFound, nil)
	return framework.SendError(http.StatusNotFound)
}
