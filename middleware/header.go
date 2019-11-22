package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var expectedContentType string

// CheckHeaders handles a response with an error if any check fails.
func CheckHeaders() echo.MiddlewareFunc {
	expectedContentType = "application/json; charset=UTF-8"

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// Check "Content-Type"
			if ctx.Request().Header.Get("Content-Type") != expectedContentType {
				return echo.NewHTTPError(http.StatusBadRequest)
			}
			return next(ctx)
		}
	}
}
