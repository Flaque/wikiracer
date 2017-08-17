package integration

import (
	"testing"

	"github.com/Flaque/wikiracer/search"
)

func TestSearch(t *testing.T) {
	search.Search("cats", "dogs")
}
