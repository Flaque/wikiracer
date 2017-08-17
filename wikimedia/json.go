package wikimedia

import (
	"github.com/buger/jsonparser"
)

// Grabs the "plcontinue "
func getPlcontinueFromJSONBytes(json []byte) string {
	cont, err := jsonparser.GetString(json, "continue", "plcontinue")
	if err != nil {
		return "" // Don't continue
	}
	return cont
}

func getLinksFromJSONBytes(json []byte) (map[string][]string, error) {

	linksByPage := make(map[string][]string)

	// Grab from query -> pages
	jsonparser.ObjectEach(json, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		title, err := jsonparser.GetString(value, "title")
		if err != nil {
			return err
		}

		// Grab from dynamically page -> links
		links, _, _, err := jsonparser.Get(value, "links")
		if err != nil {
			return err
		}

		// grab the "title"'s from "links"
		pageTitles, err := getTitlesFromPageJSON(links)
		if err != nil {
			return err
		}

		linksByPage[title] = pageTitles
		return nil
	}, "query", "pages")

	return linksByPage, nil
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
	}, "person", "avatars")

	return links, funcErr
}
