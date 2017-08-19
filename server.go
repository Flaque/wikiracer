package main

import (
	"strings"

	"github.com/Flaque/wikiracer/search"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.Default()

	// Sends a "Zoom" gif to imply that our API is super fast.
	// Method:   GET
	// Resource: http://localhost:8080/
	app.Handle("GET", "/", func(ctx context.Context) {
		ctx.HTML(ZoomGif)
	})

	// Method:   GET
	// Resource: http://localhost:8080/search/
	app.Get("/search/{start:string}/{goal:string}", func(ctx context.Context) {
		start := ctx.Params().Get("start")
		goal := ctx.Params().Get("goal")

		node, err := search.SearchConcurrently(start, goal)

		if err != nil {
			ctx.JSON(context.Map{"message": err.Error()}) // TODO: Probably not a good plan in a real prod service
		} else {
			ctx.JSON(context.Map{"path": strings.Join(node.Path, ", ")})
		}
	})

	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}
