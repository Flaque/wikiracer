package search

import (
	"errors"
	"fmt"
	"time"

	"github.com/Flaque/wikiracer/wikimedia"
	cache "github.com/patrickmn/go-cache"
	lane "gopkg.in/oleiade/lane.v1"

	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()
var linkCache = cache.New(5*time.Minute, 10*time.Minute)

func SearchConcurrently(start string, goal string) (Node, error) {
	fmt.Println("Starting conncurent search")

	queue := lane.NewQueue()
	nodes := make(chan Node)
	done := make(chan Node)

	// Start an item
	go searchAtWorker(NewNode(start, Node{goal: goal}, nil), queue, nodes, done, linkCache)

	for {
		select {

		case <-nodes:
			node := queue.Dequeue()
			go searchAtWorker(node.(Node), queue, nodes, done, linkCache)

		case finalNode := <-done:
			fmt.Println("Finished.")
			return finalNode, nil

		case <-time.After(time.Second * 20):
			fmt.Printf("Search for %s to %s timed out.\n", start, goal)
			return Node{}, errors.New("timed out")
		}
	}
}

func searchAtWorker(node Node, queue *lane.Queue, nodes chan<- Node, done chan<- Node, linkCache *cache.Cache) {
	_, seenBefore := linkCache.Get(node.current)

	if seenBefore {
		return // Skip this node since we've already seen it
	}

	pages, err := wikimedia.GetPagesLinks([]string{node.current}, "")

	for _, link := range pages[node.current] {
		if link == node.goal {

			fmt.Println("Found goal: ", node.goal)
			done <- NewNode(link, node, err) // We've found what we want!
			return
		}

		// Add a new node, tell the channel that we're going to be ready, and update our queue so we search in
		// a proper breadth first search order (and don't waste a ton of time going down some path we don't care about).
		newNode := NewNode(link, node, err)
		queue.Enqueue(newNode)
		nodes <- newNode // Tell the for that we're ready to pull from the queue
		linkCache.Set(link, true, cache.DefaultExpiration)
	}
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
