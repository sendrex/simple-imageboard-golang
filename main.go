package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AquoDev/simple-imageboard-golang/handler"
	"github.com/AquoDev/simple-imageboard-golang/middleware"
	"github.com/labstack/echo/v4"
)

var port uint64

func main() {
	// Make empty Echo instance
	app := echo.New()

	// Hide Echo banner
	app.HideBanner = true

	// Remove trailing slash from URLs
	app.Pre(middleware.RemoveTrailingSlash())

	// Register "secure" middleware
	app.Use(middleware.Secure())

	// Register default CORS on every route
	app.Use(middleware.GetCORSdefault())

	// Set regular limiter for every route
	app.Use(middleware.IPRateLimitRegular())

	// Set all static content routing
	app.Static("/", "./static")

	// Set healthcheck and server settings routing
	app.GET("/health", handler.GetHealthcheck)
	app.GET("/settings", handler.GetBoardSettings)

	// Set page routing
	pages := app.Group("/page")
	pages.GET("", handler.GetPageExample)
	pages.GET("/:id", handler.GetPage)

	// Set thread routing
	threads := app.Group("/thread")
	threads.GET("", handler.GetThreadExample)
	threads.GET("/:id", handler.GetThread)

	// Set post routing
	posts := app.Group("/post", middleware.GetCORSpost())
	posts.GET("", handler.GetPostExample)
	posts.GET("/:id", handler.GetPost)
	posts.POST("", handler.SavePost, middleware.IPRateLimitStrict(), middleware.CheckHeaders())
	posts.DELETE("", handler.DeletePost, middleware.IPRateLimitStrict(), middleware.CheckHeaders())

	// Start server
	addr := fmt.Sprintf(":%d", port)
	app.Logger.Fatal(app.Start(addr))
}

func init() {
	if parsed, err := strconv.ParseUint(os.Getenv("PORT"), 10, 0); err != nil {
		message := fmt.Errorf("[SERVER] Couldn't parse PORT @ %w", err)
		panic(message)
	} else {
		port = parsed
	}
}
