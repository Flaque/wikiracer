package wikimedia

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const baseURL string = "https://en.wikipedia.org/w/api.php?action=query&format=json&prop=links&pllimit=max"

type Params struct {
	Titles string `url:"titles,omitempty"`
	Cont   string `url:"plcontinue,omitempty"`
}

// GetLinksRoute generates an http.Request to get a bunch of links for a title from Wikimedia
func GetLinksRoute(titles []string, cont string) (*http.Request, error) {
	safeTitles := strings.Join(escapeStrings(titles), "|")
	routeUrl := fmt.Sprintf("%s&titles=%s", baseURL, safeTitles)
	if cont != "" {
		routeUrl += fmt.Sprintf("&plcontinue=%s", url.PathEscape(cont))
	}

	return http.NewRequest("GET", routeUrl, nil)
}

// escapes a bunch of strings so we can safely use them in a URL
func escapeStrings(strs []string) []string {
	safes := []string{}
	for _, s := range strs {
		safes = append(safes, url.PathEscape(s))
	}
	return safes
}
