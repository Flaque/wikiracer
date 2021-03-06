package wikimedia

import (
	"net/http"
	"time"

	t "github.com/Flaque/wikiracer/tracer"
	"github.com/go-resty/resty"
	cache "github.com/patrickmn/go-cache"
)

var linkCache = cache.New(5*time.Minute, 10*time.Minute)
var restyClient = getResty()

// GetPagesLinks returns a list of links from a wikipedia title
func GetPagesLinks(title string) ([]string, error) {
	defer t.Un(t.Trace(title))

	links, ok := linkCache.Get(title)
	if ok {
		return links.([]string), nil
	}

	pages, err := GetManyPagesLinks([]string{title}, "")
	linkCache.Set(title, pages[title], cache.DefaultExpiration)
	return pages[title], err
}

// GetManyPagesLinks returns a map of article title to a list of child links.
func GetManyPagesLinks(titles []string, cont string) (map[string][]string, error) {

	// Get our inital route that we'll use
	route, err := getLinksRoute(titles, cont)
	if err != nil {
		return nil, err
	}

	// Send off the request
	resp, err := restyClient.R().Get(route.URL.String())
	if err != nil {
		return nil, err
	}

	// Parse out our JSON
	linksPerPage, err := getLinksFromJSONBytes(resp.Body())
	if err != nil {
		return nil, err
	}

	// If we didn't get everything in the first request, we'll get a continue string
	// Which we can pass on to the next request and get ther rest of our data.
	continueString := getPlcontinueFromJSONBytes(resp.Body())
	if continueString != "" {
		newLinksPerPage, err := GetManyPagesLinks(titles, continueString)
		if err != nil {
			return linksPerPage, err // Don't mess with any error-y data
		}
		return combineMaps(linksPerPage, newLinksPerPage), nil
	}

	// If we don't need to continue, then we should just finish.
	return linksPerPage, nil
}

func getResty() *resty.Client {
	transport := http.Transport{
		MaxIdleConns:        30,
		MaxIdleConnsPerHost: 30,
	}

	return resty.New().SetTransport(&transport).SetRetryCount(3).SetTimeout(time.Duration(25 * time.Second)).SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))
}
