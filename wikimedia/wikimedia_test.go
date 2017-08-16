package wikimedia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPageLinks(t *testing.T) {
	links, err := GetPageLinks("cats")
	assert.Nil(t, err)

	// These tests technically have the posibility of failing.
	// Since it's possible for someone to edit the wikipedia page.
	assert.Contains(t, links, "Animal")
	assert.Contains(t, links, "Felis")
	assert.Contains(t, links, "Mammal")
}
