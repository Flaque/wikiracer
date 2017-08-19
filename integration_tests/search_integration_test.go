package integration

import (
	"strings"
	"testing"

	"github.com/Flaque/wikiracer/search"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	node, err := search.SearchConcurrently("Dog", "Airplane")
	assert.Nil(t, err)
	assert.NotNil(t, strings.Join(node.Path, ", "))
}
