package wikimedia

import (
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func GetPageLinks(title string) ([]string, error) {
	resp, err := getPageHTML(title)
	if err != nil {
		return []string{}, err
	}

	links, err := getPageLinksFromHTMLResponse(resp)
	defer resp.Body.Close() // No resource leaks please

	return links, nil
}

// getPageHTML queries the wikimedia api for a title and gets the HTML back
func getPageHTML(title string) (*http.Response, error) {
	req, err := GetPageHTMLRoute(title)

	// Check that we didn't have an error getting the page
	if err != nil {
		return nil, err
	}

	// Create our client and send our request
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)

	// Check we didn't have an error when we sent our request
	if err != nil {
		return nil, err
	}

	// Check we got a correct response
	if resp.StatusCode == 200 { // OK
		return resp, nil
	}

	return nil, nil //TODO Do something different here.
}

// getPageLinksFromHTMLResponse parses the page links from a response
func getPageLinksFromHTMLResponse(resp *http.Response) ([]string, error) {

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return []string{}, err
	}

	links := parseLinks(doc)

	return links, nil
}

// inSet is a helper that returns true if the item is in our set
func inSet(set map[string]bool, item string) bool {
	_, ok := set[item]
	return ok
}

// parseLinks gives the links from a goquery document (the HTML)
func parseLinks(doc *goquery.Document) []string {

	previouslySeenLinks := map[string]bool{}
	links := []string{}

	// Find all the links (or "a" tags)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {

		href, exists := s.Attr("href")
		title := LinkToTitle(href)

		weShouldAdd := exists && IsValidLink(href) && !inSet(previouslySeenLinks, title)
		if weShouldAdd {
			links = append(links, title)
			previouslySeenLinks[title] = true
		}
	})

	return links
}
