package main

import (
	"fmt"
	"os"

	handler "github.com/AquoDev/simple-imageboard-golang/server/handler"
	_ "github.com/joho/godotenv/autoload"
	iris "github.com/kataras/iris"
)

func main() {
	// Make empty Iris instance
	app := iris.New()

	// Register the template directory and engine
	template := iris.HTML("./views", ".html")
	app.RegisterView(template)

	// Set index handler
	app.Get("/", handler.GetIndex)

	// Start server
	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	app.Run(iris.Addr(addr))
}
