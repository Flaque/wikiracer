package integration

import (
	"testing"

	wikimedia "github.com/flaque/wikiracer/wikimedia"

	"github.com/stretchr/testify/assert"
)

func TestGetPageLinks(t *testing.T) {
	links, err := wikimedia.GetPageLinks("cats")
	assert.Nil(t, err)

	// These tests technically have the posibility of failing.
	// Since it's possible for someone to edit the wikipedia page.
	assert.Contains(t, links, "Animal")
	assert.Contains(t, links, "Felis")
	assert.Contains(t, links, "Mammal")
}

func TestPingSeveral(t *testing.T) {
	links, err := wikimedia.GetPageLinks("Cats")
	assert.Nil(t, err)

	for i, link := range links {
		if i > 10 {
			break
		}

		links, err := wikimedia.GetPageLinks(link)
		assert.Nil(t, err)
		assert.NotEmpty(t, links)
	}
}
