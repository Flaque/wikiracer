package wikimedia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLinksRoute(t *testing.T) {
	req, err := GetLinksRoute([]string{"Cat", "Dog", "Sea Lion"}, "")
	assert.Nil(t, err)
	assert.Equal(t, req.URL.String(), "https://en.wikipedia.org/w/api.php?action=query&format=json&prop=links&pllimit=max&titles=Cat|Dog|Sea%20Lion")
	assert.Equal(t, req.Method, "GET")
}
