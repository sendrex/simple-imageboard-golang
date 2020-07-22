package middleware

import (
	"github.com/AquoDev/simple-imageboard-golang/env"
	"github.com/AquoDev/simple-imageboard-golang/framework"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

var (
	ipRateLimiterRegular, ipRateLimiterStrict *limiter.Limiter
	storeRegular, storeStrict                 limiter.Store
)

// IPRateLimitRegular returns a regular limiter middleware.
func IPRateLimitRegular() framework.MiddlewareFunc {
	storeRegular = memory.NewStore()
	ipRateLimiterRegular = limiter.New(storeRegular, limiter.Rate{
		Period: env.GetTime("LIMITER_TIME_WINDOW_REGULAR", "s"),
		Limit:  env.GetInt64("LIMITER_REQUESTS_PER_WINDOW_REGULAR"),
	})

	return framework.RateLimiter(ipRateLimiterRegular)
}

// IPRateLimitStrict returns a strict limiter middleware.
func IPRateLimitStrict() framework.MiddlewareFunc {
	storeStrict = memory.NewStore()
	ipRateLimiterStrict = limiter.New(storeStrict, limiter.Rate{
		Period: env.GetTime("LIMITER_TIME_WINDOW_STRICT", "s"),
		Limit:  env.GetInt64("LIMITER_REQUESTS_PER_WINDOW_STRICT"),
	})

	return framework.RateLimiter(ipRateLimiterStrict)
}
