package framework

import (
	"github.com/labstack/echo/v4"
)

// NewApp returns an empty App where routes will be set up.
func NewApp() *App {
	app := echo.New()
	app.HideBanner = true

	return app
}
