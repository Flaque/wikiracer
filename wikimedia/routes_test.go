package wikimedia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPageHTMLRouteStandardCase(t *testing.T) {
	req, err := GetPageHTMLRoute("Cats")
	assert.Nil(t, err)
	assert.Equal(t, req.URL.String(),
		"https://en.wikipedia.org/api/rest_v1/page/html/Cats",
		"Wikimedia GetPageHTMLRoute should produce a valid URL.")
}

func TestGetPageHTMLRouteBadCharacters(t *testing.T) {
	req, err := GetPageHTMLRoute(";/?@=😈")
	assert.Nil(t, err)
	assert.Equal(t, req.URL.String(),
		"https://en.wikipedia.org/api/rest_v1/page/html/%3B%2F%3F@=%F0%9F%98%88",
		"Wikimedia GetPageHTMLRoute should handle bad URL characters")
}
