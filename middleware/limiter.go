package middleware

import (
	"net/http"
	"strconv"

	"github.com/AquoDev/simple-imageboard-golang/env"
	"github.com/labstack/echo/v4"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

var (
	ipRateLimiterRegular, ipRateLimiterStrict *limiter.Limiter
	storeRegular, storeStrict                 limiter.Store
	limiterTimeWindowRegular                  = env.GetTime("LIMITER_TIME_WINDOW_REGULAR", "s")
	limiterTimeWindowStrict                   = env.GetTime("LIMITER_TIME_WINDOW_STRICT", "s")
	limiterRequestsPerWindowRegular           = env.GetInt64("LIMITER_REQUESTS_PER_WINDOW_REGULAR")
	limiterRequestsPerWindowStrict            = env.GetInt64("LIMITER_REQUESTS_PER_WINDOW_STRICT")
)

// IPRateLimitRegular returns a regular limiter middleware.
func IPRateLimitRegular() echo.MiddlewareFunc {
	rate := limiter.Rate{
		Period: limiterTimeWindowRegular,
		Limit:  limiterRequestsPerWindowRegular,
	}
	storeRegular = memory.NewStore()
	ipRateLimiterRegular = limiter.New(storeRegular, rate)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ip := ctx.RealIP()
			limiterCtx, err := ipRateLimiterRegular.Get(ctx.Request().Context(), ip)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError)
			}

			header := ctx.Response().Header()
			header.Set("X-RateLimit-Limit", strconv.FormatInt(limiterCtx.Limit, 10))
			header.Set("X-RateLimit-Remaining", strconv.FormatInt(limiterCtx.Remaining, 10))
			header.Set("X-RateLimit-Reset", strconv.FormatInt(limiterCtx.Reset, 10))

			if limiterCtx.Reached {
				return echo.NewHTTPError(http.StatusTooManyRequests)
			}

			return next(ctx)
		}
	}
}

// IPRateLimitStrict returns a strict limiter middleware.
func IPRateLimitStrict() echo.MiddlewareFunc {
	rate := limiter.Rate{
		Period: limiterTimeWindowStrict,
		Limit:  limiterRequestsPerWindowStrict,
	}
	storeStrict = memory.NewStore()
	ipRateLimiterStrict = limiter.New(storeStrict, rate)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ip := ctx.RealIP()
			limiterCtx, err := ipRateLimiterStrict.Get(ctx.Request().Context(), ip)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError)
			}

			header := ctx.Response().Header()
			header.Set("X-RateLimit-Limit", strconv.FormatInt(limiterCtx.Limit, 10))
			header.Set("X-RateLimit-Remaining", strconv.FormatInt(limiterCtx.Remaining, 10))
			header.Set("X-RateLimit-Reset", strconv.FormatInt(limiterCtx.Reset, 10))

			if limiterCtx.Reached {
				return echo.NewHTTPError(http.StatusTooManyRequests)
			}

			return next(ctx)
		}
	}
}
