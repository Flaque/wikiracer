package integration

import (
	"testing"

	"github.com/Flaque/wikiracer/search"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	node, err := search.Search("Dog", "Airplane")
	assert.Nil(t, err)
	assert.NotEmpty(t, node.Path)
}
