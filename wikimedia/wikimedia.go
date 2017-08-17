package wikimedia

import (
	"net/http"
	"time"

	"github.com/go-resty/resty"
	cache "github.com/patrickmn/go-cache"
)

var linkCache = cache.New(5*time.Minute, 10*time.Minute)
var restyClient = getResty()

func GetPagesLinks(titles []string, cont string) (map[string][]string, error) {

	// Get our inital route that we'll use
	route, err := GetLinksRoute(titles, cont)
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
		newLinksPerPage, err := GetPagesLinks(titles, continueString)
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
		MaxIdleConns:        20,
		MaxIdleConnsPerHost: 20,
	}

	return resty.New().SetTransport(&transport).SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))
}

func combineMaps(main map[string][]string, extension map[string][]string) map[string][]string {
	for key, value := range extension {
		// If the main already has this key, let's combine them
		if _, ok := main[key]; ok {
			main[key] = append(main[key], value...)
		} else {
			main[key] = value
		}
	}
	return main
}
