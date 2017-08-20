package wikimedia

import (
	"github.com/buger/jsonparser"
)

// Grabs the "plcontinue" from "continue"
func getPlcontinueFromJSONBytes(json []byte) string {
	cont, err := jsonparser.GetString(json, "continue", "plcontinue")
	if err != nil {
		return "" // Don't continue
	}
	return cont
}

func getLinksFromJSONBytes(json []byte) (map[string][]string, error) {

	linksByPage := make(map[string][]string)
	var funcErr error

	// Grab from query -> pages
	jsonparser.ObjectEach(json, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		title, err := jsonparser.GetString(value, "title")

		if err != nil {
			funcErr = err // Trap the error so we can report it
			return err
		}

		// Grab from dynamically page -> links
		links, _, _, err := jsonparser.Get(value, "links")
		if err != nil {
			// This condition is more or less expected, since not all JSON items have links
			return nil // So let's just continue on
		}

		// grab the "title"'s from "links"
		pageTitles, err := getTitlesFromPageJSON(links)
		if err != nil {
			funcErr = err
			return err
		}

		linksByPage[title] = pageTitles
		return nil
	}, "query", "pages")

	return linksByPage, funcErr
}

func getTitlesFromPageJSON(data []byte) ([]string, error) {

	links := []string{}
	var funcErr error

	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		link, err := jsonparser.GetString(value, "title")
		if err != nil {
			funcErr = err
		}

		links = append(links, link)
	})

	return links, funcErr
}
