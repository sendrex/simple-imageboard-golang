package main

import (
	"fmt"
	"os"

	"github.com/AquoDev/simple-imageboard-golang/handler"
	"github.com/AquoDev/simple-imageboard-golang/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
	// Make empty Iris instance
	app := echo.New()

	// Remove trailing slash from URLs
	app.Pre(middleware.RemoveTrailingSlash())

	// Register "secure" middleware
	app.Use(middleware.Secure())

	// TODO Request limiter middleware (example: https://github.com/ulule/limiter-examples/blob/master/echo/main.go)

	// Set CORS middlewares
	corsPost := middleware.GetCORSpost()
	corsDefault := middleware.GetCORSdefault()

	// Set all static content
	app.Static("/", "./static")

	// Set page routing
	pages := app.Group("/page", corsDefault)
	pages.GET("", handler.GetPageExample)
	pages.GET("/:id", handler.GetPage)

	// Set thread routing
	threads := app.Group("/thread", corsDefault)
	threads.GET("", handler.GetThreadExample)
	threads.GET("/:id", handler.GetThread)

	// Set post routing
	posts := app.Group("/post", corsPost)
	posts.GET("", handler.GetPostExample)
	posts.GET("/:id", handler.GetPost)
	posts.POST("", handler.SavePost, middleware.CheckHeaders)
	posts.DELETE("", handler.DeletePost, middleware.CheckHeaders)

	// TODO Custom error handler

	// Start server
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	app.Logger.Fatal(app.Start(addr))
}
