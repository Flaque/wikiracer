package integration

import (
	"sort"
	"testing"

	wikimedia "github.com/flaque/wikiracer/wikimedia"

	"github.com/stretchr/testify/assert"
)

func TestGetPageLinks(t *testing.T) {
	links, err := wikimedia.GetPagesLinks([]string{"Cat", "Dog"}, "")
	assert.Nil(t, err)

	// These tests technically have the posibility of failing.
	// Since it's possible for someone to edit the wikipedia page.
	assert.Contains(t, links["Cat"], "Animal")
	assert.Contains(t, links["Cat"], "Mammal")
	assert.Contains(t, links["Cat"], "Felis")

	assert.Contains(t, links["Dog"], "Animal")
	assert.Contains(t, links["Dog"], "Canis")
	assert.Contains(t, links["Dog"], "Mammal")

	// Test to make sure we didn't some how get the same links for Cats and Dogs
	dogs := links["Dog"][:]
	cats := links["Cat"][:]
	sort.Strings(dogs)
	sort.Strings(cats)
	assert.NotEqual(t, dogs, cats)
}
