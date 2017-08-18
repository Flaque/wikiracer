package wikimedia

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