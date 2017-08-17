package wikimedia

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/dghubble/sling"
)

const baseURL string = "https://en.wikipedia.org/w/api.php?action=query?prop=links?pllimit=max"

type Params struct {
	titles string `url:"titles,omitempty"`
	cont   string `url:"plcontinue,omitempty"`
}

// GetLinksRoute generates an http.Request to get a bunch of links for a title from Wikimedia
func GetLinksRoute(titles []string, cont string) (*http.Request, error) {
	safeTitles := url.PathEscape(strings.Join(titles, "|"))
	params := &Params{titles: safeTitles, cont: cont}
	return sling.New().Get(baseURL).QueryStruct(params).Request()
}
