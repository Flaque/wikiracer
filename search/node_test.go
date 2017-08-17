package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddChildren(t *testing.T) {
	node := Node{
		current:  "cats",
		children: []Node{},
		path:     []Node{},
	}

	node = node.AddChildren([]string{"dogs", "bears", "lions"})

	assert.Equal(t, len(node.children), 3)
}
