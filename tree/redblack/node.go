package redblack

const (
	colorBlack = 0
	colorRed   = 1
)

type node struct {
	parent *node
	left   *node
	right  *node
	key    int
	color  int
}

func newNode(key int) *node {
	return &node{
		parent: nil,
		left:   nil,
		right:  nil,
		key:    key,
		color:  colorRed, // new nodes are always RED
	}
}

func (n node) isBlack() bool {
	return n.color == colorBlack
}

func (n node) isRed() bool {
	return n.color == colorRed
}
