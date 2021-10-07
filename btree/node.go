package btree

type node struct {
	keys     []int
	children []*node
	leaf     bool
	n        int
}

func newNode(t int) *node {
	return &node{
		keys:     make([]int, 2*t-1),
		children: make([]*node, 2*t),
		leaf:     false,
		n:        0,
	}
}
