package main

import (
	"fmt"
	"os"

	"github.com/AquoDev/simple-imageboard-golang/server/handler"
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
	app.UseGlobal(func(ctx iris.Context) {
		ctx.Header("Content-Type", "application/json; charset=UTF-8")
		ctx.Next()
	})

	// Set page handler
	pages := app.Party("/page")
	{
		pages.Get("/", handler.GetPageExample)
		pages.Get("/{id:uint8 min(0) max(9)}", handler.GetPage)
	}

	// Set 404 Not Found handler
	app.OnAnyErrorCode(handler.PathNotFound)

	// Start server
	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	app.Run(
		iris.Addr(addr),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
