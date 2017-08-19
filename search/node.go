package search

type Node struct {
	current string
	goal    string
	Path    []string
	Err     error
}

func NewNode(link string, parent Node, err error) Node {
	path := []string{link}
	if parent.Path != nil {
		path = append(parent.Path, link)
	}

	return Node{
		current: link,
		goal:    parent.goal,
		Path:    path,
		Err:     err,
	}
}
