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

	// TODO Request limiter middleware (example: https://github.com/ulule/limiter-examples/blob/master/echo/main.go)

	// Remove trailing slash from URLs
	app.Pre(middleware.RemoveTrailingSlash())

	// Register "secure" middleware
	app.Use(middleware.Secure())

	// Register default CORS on every route
	app.Use(middleware.GetCORSdefault())

	// Set all static content routing
	app.Static("/", "./static")

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
	posts.POST("", handler.SavePost, middleware.CheckHeaders)
	posts.DELETE("", handler.DeletePost, middleware.CheckHeaders)

	// Start server
	addr := fmt.Sprintf(":%d", port)
	app.Logger.Fatal(app.Start(addr))
}

func init() {
	if parsed, err := strconv.ParseUint(os.Getenv("PORT"), 10, 0); err != nil {
		panic(err)
	} else {
		port = parsed
	}

	fmt.Println("[SERVER]: Port OK")
}
