package main

import (
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
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx context.Context) {
		ctx.JSON(context.Map{"message": "Hello iris web framework."})
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"))
}
