package handler

import (
	"net/http"

	"github.com/AquoDev/simple-imageboard-golang/cache"
	"github.com/AquoDev/simple-imageboard-golang/database"
	"github.com/AquoDev/simple-imageboard-golang/framework"
)

var (
	examplePost = map[string]string{
		"info":    "If you want to see any post, see the example.",
		"example": "{url}/post/{id}",
	}

	okMessage = map[string]string{
		"message": http.StatusText(http.StatusOK),
	}
)

// GetPostExample handles a JSON response with a how-to example.
func GetPostExample(ctx framework.Context) error {
	return framework.SendOK(ctx, examplePost)
}

// GetPost handles a JSON response with a post.
func GetPost(ctx framework.Context) error {
	// Parse post ID. Is the parsing invalid?
	// ── Yes:	return a failed JSON response.
	// ── No:	continue. ID was parsed successfully.
	id, err := framework.GetID(ctx)
	if err != nil {
		return framework.SendError(http.StatusNotFound)
	}

	// Try to get the post from cache. Is it stored in cache?
	// ── Yes:	check if what's cached is an error or a post. Both options lead to a response.
	// ── No:	continue. It's not in cache, but it could be stored in the database.
	if response, err := cache.GetCachedPost(id); err == nil {
		// Is the post cached?
		// ── Yes:	return the cached post.
		// ── No:	it's an error, so return that error as a failed JSON response.
		if response.Status == http.StatusOK {
			return framework.SendOK(ctx, response.Data)
		}
		return framework.SendError(response.Status)
	}

	// Try to get the post from the database. Does it exist?
	// ── Yes:	return the post.
	// ── No:	continue. The post doesn't exist.
	if response, err := database.GetPost(id); err == nil {
		go cache.SetCachedPost(id, http.StatusOK, response)
		return framework.SendOK(ctx, response)
	}

	// Return a 404 NotFound JSON response after caching it.
	go cache.SetCachedPost(id, http.StatusNotFound, nil)
	return framework.SendError(http.StatusNotFound)
}

// SavePost handles a JSON response and saves the data as a post.
func SavePost(ctx framework.Context) error {
	// Read body (JSON) from the request. Is the body request bad formed?
	// ── Yes:	return a failed JSON response.
	// ── No:	continue. Body is well formed and could be read.
	post, err := framework.BindPost(ctx)
	if err != nil {
		return framework.SendError(http.StatusBadRequest)
	}

	// Does the post we receive reply to another post?
	// ── Yes:	check if the reply exists. We can't reply to a post that doesn't exists (foreign key constrain).
	// ── No:	continue. The post we receive starts a parent thread.
	if post.RepliesToAnotherPost() {
		// Does the reply exists and is allowed to be replied?
		// ── Yes:	make the relation between post and reply.
		// ── No:	return a failed JSON response. The reply doesn't exist or thread is closed.
		if reply, err := database.GetPost(*post.ReplyTo); err == nil && reply.AllowsReplies() {
			post.MakeRelationWith(reply)
		} else {
			return framework.SendError(http.StatusForbidden)
		}
	}

	// Try to save the post in the database. Does it fail to be saved?
	// ── Yes:	return a failed JSON response.
	// ── No:	continue.
	response, err := database.SavePost(post)
	if err != nil {
		return framework.SendError(http.StatusBadRequest)
	}

	// Does the post start a parent thread?
	// ── Yes: threads that are beyond the last page must be deleted.
	// ── No (post is a reply), but the post doesn't sage: bump that parent thread.
	// ── No (post is a reply), but the post sages: continue. Parent thread is not bumped and no threads should be deleted.
	go func() {
		// The following statements are meant to be ran in background, as they don't affect the response
		if post.IsAParentThread() {
			database.DeleteOldThreads()
		} else if !post.Sages() {
			database.BumpThread(post)
		}
	}()

	// Return a 201 Created JSON response with the post ID and its delete code.
	return framework.SendCreated(ctx, response)
}

// DeletePost handles a JSON response with a post.
func DeletePost(ctx framework.Context) error {
	// Read body (JSON) from the request. Is the body request bad formed?
	// ── Yes:	return a failed JSON response. Maybe some field's missing.
	// ── No:	continue. Body is well formed and could be read.
	data, err := framework.BindDeleteData(ctx)
	if err != nil {
		return framework.SendError(http.StatusBadRequest)
	}

	// Try to delete post (and thread if the post has "parent_thread == nil"). Couldn't it be deleted?
	// ── Yes:	return a failed JSON response. Incorrect delete code or post doesn't exist.
	// ── No:	continue. Correct delete code, the post was deleted.
	if err := database.DeletePost(data); err != nil {
		return framework.SendError(http.StatusUnauthorized)
	}

	// Return a 200 OK JSON response.
	return framework.SendOK(ctx, okMessage)
}

// UpdatePost handles a JSON response with a post.
func UpdatePost(ctx framework.Context) error {
	// Read body (JSON) from the request. Is the body request bad formed?
	// ── Yes:	return a failed JSON response. Maybe some field's missing.
	// ── No:	continue. Body is well formed and could be read.
	data, err := framework.BindUpdateData(ctx)
	if err != nil {
		return framework.SendError(http.StatusBadRequest)
	}

	// Try to update post (it must be a parent thread). Couldn't it be updated?
	// ── Yes:	return a failed JSON response. Incorrect password, post doesn't exist or post isn't a parent thread.
	// ── No:	continue. Correct password, the post was updated.
	if err := database.UpdatePost(data); err != nil {
		return framework.SendError(http.StatusUnauthorized)
	}

	// Return a 200 OK JSON response.
	return framework.SendOK(ctx, okMessage)
}
