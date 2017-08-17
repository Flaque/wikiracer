package search

type Node struct {
	current  string
	goal     string
	children []Node
	path     []Node
	err      error
	depth    int
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

func NewNode(link string, goal string, depth int, err error) Node {
	return Node{
		current:  link,
		goal:     goal,
		children: []Node{},
		path:     []Node{},
		depth:    depth,
		err:      err,
	}
}

func NewNodeWithChildren(link string, goal string, depth int, err error, links []string) Node {
	node := NewNode(link, goal, depth, err)
	return node.AddChildren(links)
}
