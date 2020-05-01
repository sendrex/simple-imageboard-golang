package handler

import (
	"net/http"

	"github.com/AquoDev/simple-imageboard-golang/env"
	"github.com/labstack/echo/v4"
)

var boardSettings = map[string]map[string]uint64{
	"server": {
		"THREADS_PER_PAGE": env.GetUint64("THREADS_PER_PAGE"),
		"POSTS_PER_THREAD": env.GetUint64("POSTS_PER_THREAD"),
	},
	"cache": {
		"REDIS_EXPIRE_TIME_POST":   env.GetUint64("REDIS_EXPIRE_TIME_POST"),
		"REDIS_EXPIRE_TIME_THREAD": env.GetUint64("REDIS_EXPIRE_TIME_THREAD"),
		"REDIS_EXPIRE_TIME_PAGE":   env.GetUint64("REDIS_EXPIRE_TIME_PAGE"),
	},
}

// GetBoardSettings handles a JSON response with this board settings.
func GetBoardSettings(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, boardSettings)
}
