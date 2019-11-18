package main

import (
	"fmt"
	"os"

	"github.com/AquoDev/simple-imageboard-golang/server/handler"
	"github.com/AquoDev/simple-imageboard-golang/server/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kataras/iris"
)

func main() {
	// Make empty Iris instance
	app := iris.New()

	// Register the template directory and engine
	template := iris.HTML("./views", ".html")
	app.RegisterView(template)

	// Set index handler
	app.Get("/", handler.GetIndex)

	// Set "Content-Type" header for every response except for the index
	app.UseGlobal(middleware.UseContentTypeJSON)

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
