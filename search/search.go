package search

import (
	"fmt"
	"time"

	"github.com/Flaque/wikiracer/wikimedia"
	"github.com/korovkin/limiter"

	"go.uber.org/zap"
)

const MaxDepth = 1

var logger, _ = zap.NewProduction()

func Search(start string, goal string) {

	fmt.Println("Starting Search...")

	root := queryLink(start, goal, MaxDepth)

	nodes := make(chan Node, len(root.children))

	startTime := time.Now()
	limit := limiter.NewConcurrencyLimiter(40)
	for _, child := range root.children {
		child := child // Capture child value
		limit.Execute(func() {
			nodes <- queryNode(child)
		})
	}
	limit.Wait()
	fmt.Printf("\nCompleted in %s \n", time.Since(startTime))

	fmt.Println("Finished some nodes...")
}

func queryNode(node Node) Node {
	return queryLink(node.current, node.goal, node.depth)
}

func queryLink(link string, goal string, depth int) Node {
	links, err := wikimedia.GetPageLinks(link)
	return NewNodeWithChildren(link, goal, depth, err, links)
}
