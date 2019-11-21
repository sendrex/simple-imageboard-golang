package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CheckHeaders handles a response with an error if any check fails.
func CheckHeaders(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Check "Content-Type"
		if ctx.Request().Header.Get("Content-Type") != "application/json; charset=UTF-8" {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return next(ctx)
	}
}
