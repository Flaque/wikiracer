package wikimedia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombineMaps(t *testing.T) {
	main := make(map[string][]string)
	main["Dog"] = []string{"Fluffer", "Shmoofer", "Dippledorf"}
	main["Cat"] = []string{"Meowington"}

	extension := make(map[string][]string)
	extension["Dog"] = []string{"Spottyfloops"}

	pets := combineMaps(main, extension)

	assert.Equal(t, len(pets["Cat"]), 1)
	assert.Contains(t, pets["Cat"], "Meowington")

	assert.Equal(t, len(pets["Dog"]), 4)
	assert.Contains(t, pets["Dog"], "Spottyfloops")
	assert.Contains(t, pets["Dog"], "Fluffer")
}
