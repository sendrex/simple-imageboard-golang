package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

var boardSettings map[string]map[string]uint64

// GetHealthcheck handles a JSON response with a 200 OK if everything's alright.
func GetHealthcheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": http.StatusText(http.StatusOK),
	})
}

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
	if parseInt, err := strconv.ParseUint(os.Getenv("THREADS_PER_PAGE"), 10, 0); err != nil {
		panic(err)
	} else {
		threadsPerPage = parseInt
	}

	if parseInt, err := strconv.ParseUint(os.Getenv("POSTS_PER_THREAD"), 10, 0); err != nil {
		panic(err)
	} else {
		postsPerThread = parseInt
	}

	// Parse Redis cache times
	if parseInt, err := strconv.ParseUint(os.Getenv("REDIS_EXPIRE_TIME_PAGE"), 10, 0); err != nil {
		panic(err)
	} else {
		cacheTimePage = parseInt
	}

	if parseInt, err := strconv.ParseUint(os.Getenv("REDIS_EXPIRE_TIME_THREAD"), 10, 0); err != nil {
		panic(err)
	} else {
		cacheTimeThread = parseInt
	}

	if parseInt, err := strconv.ParseUint(os.Getenv("REDIS_EXPIRE_TIME_POST"), 10, 0); err != nil {
		panic(err)
	} else {
		cacheTimePost = parseInt
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
