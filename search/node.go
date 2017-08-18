package search

type Node struct {
	current  string
	goal     string
	children []Node
	path     []Node
	err      error
}

func (worker *Node) TunnyReady() bool {
	return true
}

// This is where the work actually happens
func (worker *Node) TunnyJob(data interface{}) interface{} {
	if outputStr, ok := data.(string); ok {
		return ("custom job done: " + outputStr)
	}
	return nil
}

func (n Node) AddChildren(links []string) Node {
	for _, link := range links {
		n.children = append(n.children, Node{
			current:  link,
			children: []Node{},
			path:     append(n.path, n),
			err:      nil,
		})
	}

	return n
}

func NewNode(link string, goal string, err error) Node {
	return Node{
		current:  link,
		goal:     goal,
		children: []Node{},
		path:     []Node{},
		err:      err,
	}
}

func NewNodeWithChildren(link string, goal string, err error, links []string) Node {
	node := NewNode(link, goal, err)
	return node.AddChildren(links)
}
