package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

var boardSettings map[string]map[string]uint64

// GetBoardSettings handles a JSON response with this board settings.
func GetBoardSettings(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, boardSettings)
}

func init() {
	var (
		threadsPerPage, postsPerThread                uint64
		cacheTimePage, cacheTimeThread, cacheTimePost uint64
	)

	// Parse board settings
	if parseUint, err := strconv.ParseUint(os.Getenv("THREADS_PER_PAGE"), 10, 0); err != nil {
		message := fmt.Errorf("[HANDLER] Couldn't parse THREADS_PER_PAGE @ %w", err)
		panic(message)
	} else {
		threadsPerPage = parseUint
	}

	if parseUint, err := strconv.ParseUint(os.Getenv("POSTS_PER_THREAD"), 10, 0); err != nil {
		message := fmt.Errorf("[HANDLER] Couldn't parse POSTS_PER_THREAD @ %w", err)
		panic(message)
	} else {
		postsPerThread = parseUint
	}

	// Parse Redis cache times
	if parseUint, err := strconv.ParseUint(os.Getenv("REDIS_EXPIRE_TIME_PAGE"), 10, 0); err != nil {
		message := fmt.Errorf("[HANDLER] Couldn't parse REDIS_EXPIRE_TIME_PAGE @ %w", err)
		panic(message)
	} else {
		cacheTimePage = parseUint
	}

	if parseUint, err := strconv.ParseUint(os.Getenv("REDIS_EXPIRE_TIME_THREAD"), 10, 0); err != nil {
		message := fmt.Errorf("[HANDLER] Couldn't parse REDIS_EXPIRE_TIME_THREAD @ %w", err)
		panic(message)
	} else {
		cacheTimeThread = parseUint
	}

	if parseUint, err := strconv.ParseUint(os.Getenv("REDIS_EXPIRE_TIME_POST"), 10, 0); err != nil {
		message := fmt.Errorf("[HANDLER] Couldn't parse REDIS_EXPIRE_TIME_POST @ %w", err)
		panic(message)
	} else {
		cacheTimePost = parseUint
	}

	// Fill board settings
	boardSettings = make(map[string]map[string]uint64)
	boardSettings["server"] = map[string]uint64{
		"THREADS_PER_PAGE": threadsPerPage,
		"POSTS_PER_THREAD": postsPerThread,
	}
	boardSettings["cache"] = map[string]uint64{
		"REDIS_EXPIRE_TIME_POST":   cacheTimePost,
		"REDIS_EXPIRE_TIME_THREAD": cacheTimeThread,
		"REDIS_EXPIRE_TIME_PAGE":   cacheTimePage,
	}
}
