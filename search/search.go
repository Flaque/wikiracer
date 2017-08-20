package search

import (
	"errors"
	"strings"
	"time"

	"gopkg.in/oleiade/lane.v1"

	t "github.com/Flaque/wikiracer/tracer"
	"github.com/Flaque/wikiracer/wikimedia"
	cache "github.com/patrickmn/go-cache"

	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

func SearchConcurrently(start string, goal string) (Node, error) {
	logStart(start, goal)

	// Make sure we're not just doing a big loop
	if start == goal {
		return Node{Path: []string{start}}, nil
	}

	visitedCache := cache.New(5*time.Minute, 10*time.Minute)
	pqueue := lane.NewPQueue(lane.MINPQ)
	nodes := make(chan Node, 200)
	done := make(chan Node)

	// Start an item
	go searchAtWorker(NewNode(start, Node{goal: goal, depth: 0}, nil), pqueue, nodes, done, visitedCache)

	for {
		select {
		case <-nodes:
			node, _ := pqueue.Pop()
			go searchAtWorker(node.(Node), pqueue, nodes, done, visitedCache)

		case finalNode := <-done:
			logEnd(finalNode)
			return finalNode, nil

		case <-time.After(time.Second * 30):
			logTimeout(start, goal, 30)
			return Node{}, errors.New("timed out")
		}
	}
}

func searchAtWorker(node Node, pqueue *lane.PQueue, nodes chan<- Node, done chan<- Node, visitedCache *cache.Cache) {
	defer t.Un(t.Trace("Search: " + node.current))
	_, seenBefore := visitedCache.Get(node.current)

	if seenBefore {
		return // Skip this node since we've already seen it
	}

	logNode(node)

	pages, err := wikimedia.GetPagesLinks(node.current, "")
	if err != nil {
		logError(err)
		return // Don't search down paths with errors
	}

	for _, link := range pages[node.current] {
		if link == node.goal {
			done <- NewNode(link, node, nil) // We've found what we want!
			return
		}

		// Add a new node, tell the channel that we're going to be ready, and update our queue so we search in
		// a proper breadth first search order (and don't waste a ton of time going down some path we don't care about).
		newNode := NewNode(link, node, nil)
		pqueue.Push(newNode, newNode.depth)
		nodes <- newNode // Tell the for that we're ready to pull from the queue
		visitedCache.Set(link, true, cache.DefaultExpiration)
	}
}

func logNode(node Node) {
	logger.Info("Run node",
		zap.String("link", node.current),
		zap.Int("Depth", node.depth),
	)
}

func logStart(start string, goal string) {
	logger.Info("Starting search",
		zap.String("start", start),
		zap.String("goal", goal),
	)
}

func logEnd(node Node) {
	logger.Info("Found path",
		zap.String("path", strings.Join(node.Path, ", ")),
	)
}

func logTimeout(start string, goal string, time int) {
	logger.Warn("Timed out",
		zap.String("start", start),
		zap.String("goal", goal),
		zap.Int("seconds", time),
	)
}

func logError(err error) {
	logger.Error(err.Error(),
		zap.String("Message", err.Error()),
	)
}
