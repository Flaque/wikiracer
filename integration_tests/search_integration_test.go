package integration

import (
	"testing"

	"github.com/Flaque/wikiracer/search"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	ok, err := search.SearchConcurrently("Octal", "Number")
	assert.Nil(t, err)
	assert.True(t, ok)
}
