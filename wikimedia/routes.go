package wikimedia

import (
	"net/http"
	"net/url"

	"github.com/dghubble/sling"
)

const baseURL string = "https://en.wikipedia.org/api/rest_v1/"

// GetPageHTMLRoute generates an http.Request that could be sent off to wikimedia
func GetPageHTMLRoute(title string) (*http.Request, error) {
	safeTitle := url.PathEscape(title)
	return sling.New().Get(baseURL).Path("page/").Path("html/").Path(safeTitle).Request()
}
