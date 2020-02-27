package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetHealthcheck handles a JSON response with a 200 OK if everything's alright.
func GetHealthcheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": http.StatusText(http.StatusOK),
	})
}
