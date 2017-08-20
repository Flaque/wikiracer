package main

import (
	"strings"
	"time"

	"github.com/Flaque/wikiracer/search"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	cache "github.com/patrickmn/go-cache"
)

var requestCache = cache.New(5*time.Minute, 10*time.Minute)

func cacheString(start string, goal string) string {
	return start + "," + goal
}

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

		// Check that we haven't recently had this request
		item, ok := requestCache.Get(cacheString(start, goal))
		if ok {
			node := item.(search.Node)
			ctx.JSON(context.Map{"path": strings.Join(node.Path, ", ")})
			return
		}

		node, err := search.SearchConcurrently(start, goal)
		requestCache.Set(cacheString(start, goal), node, cache.DefaultExpiration) // Add to our cache

		if err != nil {
			ctx.JSON(context.Map{"message": err.Error()}) // TODO: Probably not a good plan in a real prod service
		} else {
			ctx.JSON(context.Map{"path": strings.Join(node.Path, ", ")})
		}
	})

	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}
