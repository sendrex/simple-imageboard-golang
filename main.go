package main

import (
	"fmt"
	"os"

	"github.com/AquoDev/simple-imageboard-golang/server/handler"
	"github.com/AquoDev/simple-imageboard-golang/server/middleware"
	"github.com/iris-contrib/middleware/secure"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kataras/iris"
)

func main() {
	// Make empty Iris instance
	app := iris.New()

	// Register "secure" middleware
	security := secure.New(secure.Options{
		STSSeconds:              315360000,
		STSIncludeSubdomains:    true,
		STSPreload:              true,
		ForceSTSHeader:          false,
		FrameDeny:               true,
		CustomFrameOptionsValue: "SAMEORIGIN",
		ContentTypeNosniff:      true,
		BrowserXSSFilter:        true,
	})
	app.UseGlobal(security.Serve)

	// Register the template directory and engine
	viewEngine := iris.HTML("./views", ".html")
	app.RegisterView(viewEngine)

	// Set robots.txt and all static content
	app.HandleDir("/", "./static")

	// Set index handler
	app.Get("/", handler.GetIndex)

	// TODO implement CORS middleware

	// Set page handler
	pages := app.Party("/page")
	{
		pages.Get("/", handler.GetPageExample)
		pages.Get("/{id:uint8 min(0) max(9)}", handler.GetPage)
	}

	// Set thread handler
	threads := app.Party("/thread")
	{
		threads.Get("/", handler.GetThreadExample)
		threads.Get("/{id:uint64 min(1)}", handler.GetThread)
	}

	// Set post handler
	posts := app.Party("/post")
	{
		posts.Get("/", handler.GetPostExample)
		posts.Get("/{id:uint64 min(1)}", handler.GetPost)
		posts.Post("/", middleware.CheckHeaders, handler.SavePost)
		posts.Delete("/", middleware.CheckHeaders, handler.DeletePost)
	}

	// Set 404 Not Found handler
	app.OnAnyErrorCode(handler.PathNotFound)

	// Start server
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	app.Run(iris.Addr(addr))
}
