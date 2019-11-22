package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// GetCORSpost returns CORS for any post request.
func GetCORSpost() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodOptions, http.MethodGet, http.MethodHead, http.MethodPost, http.MethodDelete},
	})
}

// GetCORSdefault returns CORS for any thread or page request.
func GetCORSdefault() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodOptions, http.MethodGet, http.MethodHead},
	})
}
