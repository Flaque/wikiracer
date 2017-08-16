package wikimedia

import (
	"io/ioutil"
	"net/http"
)

// BodyString is a helper that returns the body string from a response.
func BodyString(resp *http.Response) (string, error) {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	return bodyString, err
}
