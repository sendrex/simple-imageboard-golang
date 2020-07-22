package handler

import (
	"github.com/AquoDev/simple-imageboard-golang/env"
	"github.com/AquoDev/simple-imageboard-golang/framework"
)

var boardSettings = map[string]map[string]uint64{
	"limiter": {
		"LIMITER_TIME_WINDOW_REGULAR":         env.GetUint64("LIMITER_TIME_WINDOW_REGULAR"),
		"LIMITER_TIME_WINDOW_STRICT":          env.GetUint64("LIMITER_TIME_WINDOW_STRICT"),
		"LIMITER_REQUESTS_PER_WINDOW_REGULAR": env.GetUint64("LIMITER_REQUESTS_PER_WINDOW_REGULAR"),
		"LIMITER_REQUESTS_PER_WINDOW_STRICT":  env.GetUint64("LIMITER_REQUESTS_PER_WINDOW_STRICT"),
	},
	"server": {
		"POSTS_PER_THREAD":   env.GetUint64("POSTS_PER_THREAD"),
		"MAX_PARENT_THREADS": env.GetUint64("MAX_PARENT_THREADS"),
	},
	"cache": {
		"REDIS_EXPIRE_TIME_POST":   env.GetUint64("REDIS_EXPIRE_TIME_POST"),
		"REDIS_EXPIRE_TIME_THREAD": env.GetUint64("REDIS_EXPIRE_TIME_THREAD"),
		"REDIS_EXPIRE_TIME_INDEX":  env.GetUint64("REDIS_EXPIRE_TIME_INDEX"),
	},
}

// GetBoardSettings handles a JSON response with this board settings.
func GetBoardSettings(ctx framework.Context) error {
	return framework.SendOK(ctx, boardSettings)
}
