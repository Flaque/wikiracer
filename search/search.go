package search

import (
	"fmt"
	"time"

	"github.com/Flaque/wikiracer/wikimedia"
	cache "github.com/patrickmn/go-cache"
	lane "gopkg.in/oleiade/lane.v1"

	"go.uber.org/zap"
)

const MaxDepth = 5
const MaxBulkLinks = 15
const MaxWorkers = 10

var logger, _ = zap.NewProduction()

func SearchConcurrently(start string, goal string) (bool, error) {
	fmt.Println("Starting conncurent search")

	queue := lane.NewQueue()
	nodes := make(chan Node)
	done := make(chan bool)

	linkCache := cache.New(5*time.Minute, 10*time.Minute)

	// Start an item
	go searchAtWorker(NewNode(start, goal, nil), queue, nodes, done, linkCache)

	for {
		select {

		case <-nodes:
			node := queue.Dequeue()
			go searchAtWorker(node.(Node), queue, nodes, done, linkCache)

		case isDone := <-done:
			fmt.Println("Finished.")
			return isDone, nil
		}
	}
}

func searchAtWorker(node Node, queue *lane.Queue, nodes chan<- Node, done chan<- bool, linkCache *cache.Cache) {
	_, seenBefore := linkCache.Get(node.current)

	if seenBefore {
		return // Skip this node since we've already seen it
	}

	pages, err := wikimedia.GetPagesLinks([]string{node.current}, "")

	for _, link := range pages[node.current] {
		if link == node.goal {

			fmt.Println("Found goal: ", node.goal)
			done <- true // We've found what we want!
			return
		}

		newNode := NewNode(link, node.goal, err)
		queue.Enqueue(newNode)
		nodes <- newNode
		linkCache.Set(link, true, cache.DefaultExpiration)
	}
}

// Returns a slice of "chunks", where each chunk is a list of links to search.
// This is so we can batch up our links that we search in one request to avoid
// 100's of HTTP requests a second. (Faster and more polite!)
func getChunkOfLinks(nodes *lane.PQueue) []Node {
	chunk := []Node{}
	for i := 0; i < MaxBulkLinks && !nodes.Empty(); i++ {
		node, _ := nodes.Pop()
		chunk = append(chunk, node.(Node))
	}

	return chunk
}

func priority() int {
	return 1 // TODO find a better way to priority this
}

// Returns true if the item is in the array
func contains(arr []string, str string) bool {
	for _, current := range arr {
		if str == current {
			return true
		}
	}
	return false
}
