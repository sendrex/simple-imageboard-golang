package handler

import "github.com/kataras/iris"

// GetRobotsTxt renders robots.txt file.
func GetRobotsTxt(ctx iris.Context) {
	ctx.ServeFile("./static/robots.txt", false)
}
