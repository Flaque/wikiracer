package wikimedia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIgnoreHashes(t *testing.T) {
	assert.Equal(t, "Cat", IgnoreHashes("Cat#cite_note-NC012913-167"))
	assert.Equal(t, "Cat", IgnoreHashes("Cat"))
	assert.Equal(t, "", IgnoreHashes(""))
	assert.Equal(t, "Cat", IgnoreHashes("Cat##"))
	assert.Equal(t, "Cat", IgnoreHashes("Cat#something#anotherthing"))
	assert.Equal(t, "", IgnoreHashes("#Cat#ThenSomethingElse"))
}

func TestTrimLinkPrefix(t *testing.T) {
	assert.Equal(t, "Cat", TrimLinkPrefix("./Cat"))
	assert.Equal(t, "", TrimLinkPrefix(""))
	assert.Equal(t, "Cat#hey", TrimLinkPrefix("./Cat#hey"))
}

func TestLinkToTitle(t *testing.T) {
	assert.Equal(t, "Cat", LinkToTitle("./Cat#cite_note-NC012913-167"))
	assert.Equal(t, "Cat", LinkToTitle("./Cat#cite_note-NC012913-167#Heythere"))
	assert.Equal(t, "", LinkToTitle(""))
}
