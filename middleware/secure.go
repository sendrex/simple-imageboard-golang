package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Secure adds to the response some headers to increase protection.
func Secure() echo.MiddlewareFunc {
	return middleware.Secure()
}
