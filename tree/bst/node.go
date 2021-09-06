package bst

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func NewNode(data int) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}
