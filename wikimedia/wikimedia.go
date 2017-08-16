package wikimedia

import (
	"fmt"
	"net/http"
	"time"
)

func GetPageHTML(title string) (string, error) {
	req, err := GetPageHTMLRoute(title)

	// Check that we didn't have an error getting the page
	if err != nil {
		fmt.Print(err)
	}

	// Create our client and send our request
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)

	// Check we didn't have an error when we sent our request
	if err != nil {
		return "", err
	}

	// Close our request
	defer resp.Body.Close()

	// Check we got a correct response
	if resp.StatusCode == 200 { // OK
		return BodyString(resp)
	}

	return "", nil //TODO Do something different here.
}
