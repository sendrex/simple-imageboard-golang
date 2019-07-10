package handler

import "github.com/kataras/iris"

// GetIndex renders the index HTML.
func GetIndex(ctx iris.Context) {
	ctx.View("index.html")
}
