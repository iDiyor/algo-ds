package avl

type node struct {
	parent        *node
	left          *node
	right         *node
	key           int
	height        int
	balanceFactor int
}

func newNode(key int) *node {
	return &node{
		parent:        nil,
		left:          nil,
		right:         nil,
		key:           key,
		height:        0,
		balanceFactor: 0,
	}
}
