package framework

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ulule/limiter/v3"
)

// CheckHeader validates the request header.
func CheckHeader(expectedContentType string) MiddlewareFunc {
	return func(next HandlerFunc) HandlerFunc {
		return func(ctx Context) error {
			// Check "Content-Type"
			if ctx.Request().Header.Get("Content-Type") != expectedContentType {
				return SendError(http.StatusBadRequest)
			}
			return next(ctx)
		}
	}
}

// RateLimiter sets and configures a rate limiter.
func RateLimiter(ipRateLimiter *limiter.Limiter) MiddlewareFunc {
	return func(next HandlerFunc) HandlerFunc {
		return func(ctx Context) error {
			ip := ctx.RealIP()
			limiterCtx, err := ipRateLimiter.Get(ctx.Request().Context(), ip)
			if err != nil {
				return SendError(http.StatusInternalServerError)
			}

			header := ctx.Response().Header()
			header.Set("X-RateLimit-Limit", strconv.FormatInt(limiterCtx.Limit, 10))
			header.Set("X-RateLimit-Remaining", strconv.FormatInt(limiterCtx.Remaining, 10))
			header.Set("X-RateLimit-Reset", strconv.FormatInt(limiterCtx.Reset, 10))

			if limiterCtx.Reached {
				return SendError(http.StatusTooManyRequests)
			}

			return next(ctx)
		}
	}
}

// CORS returns a CORS middleware given the allowed methods.
func CORS(methods ...string) MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: methods,
	})
}

// Secure sets some headers to avoid some attacks.
func Secure() MiddlewareFunc {
	return middleware.Secure()
}

// RemoveTrailingSlash removes the trailing slash from URLs.
func RemoveTrailingSlash() MiddlewareFunc {
	return middleware.RemoveTrailingSlash()
}
