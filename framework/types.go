package framework

import (
	"github.com/labstack/echo/v4"
)

type (
	// App is a type alias for "echo.Echo".
	App = echo.Echo

	// Context is a type alias for "echo.Context"
	Context = echo.Context

	// HTTPError is a type alias for "echo.HTTPError".
	HTTPError = echo.HTTPError

	// MiddlewareFunc is a type alias for "echo.MiddlewareFunc".
	MiddlewareFunc = echo.MiddlewareFunc

	// HandlerFunc is a type alias for "echo.HandlerFunc".
	HandlerFunc = echo.HandlerFunc
)
