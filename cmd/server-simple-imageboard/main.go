package main

import (
	"fmt"

	"github.com/AquoDev/simple-imageboard-golang/env"
	"github.com/AquoDev/simple-imageboard-golang/framework"
	"github.com/AquoDev/simple-imageboard-golang/handler"
	"github.com/AquoDev/simple-imageboard-golang/middleware"
)

func main() {
	// Make empty instance
	app := framework.NewApp()

	// Remove trailing slash from URLs
	app.Pre(middleware.RemoveTrailingSlash())

	// Register "secure" middleware
	app.Use(middleware.Secure())

	// Register default CORS on every route
	app.Use(middleware.DefaultCORS())

	// Set regular limiter for every route
	app.Use(middleware.IPRateLimitRegular())

	// Set all static content routing
	app.Static("/", "./static")

	// Set healthcheck and server settings routing
	app.GET("/health", handler.GetHealthcheck)
	app.GET("/settings", handler.GetBoardSettings)

	// Set index routing
	app.GET("/index", handler.GetIndex)

	// Set thread routing
	threads := app.Group("/thread")
	threads.GET("", handler.GetThreadExample)
	threads.GET("/:id", handler.GetThread)

	// Set post routing
	posts := app.Group("/post", middleware.ExtendedCORS())
	posts.GET("", handler.GetPostExample)
	posts.GET("/:id", handler.GetPost)
	posts.POST("", handler.SavePost, middleware.IPRateLimitStrict(), middleware.CheckHeader())
	posts.DELETE("", handler.DeletePost, middleware.IPRateLimitStrict(), middleware.CheckHeader())
	posts.PUT("", handler.UpdatePost, middleware.IPRateLimitStrict(), middleware.CheckHeader())

	// Start server
	addr := fmt.Sprintf(":%d", env.GetInt("PORT"))
	app.Logger.Fatal(app.Start(addr))
}
