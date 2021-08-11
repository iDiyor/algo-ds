package singly_linked_list

type Node struct {
	Next *Node
	Data int
}

func NewNode(data int) *Node {
	return &Node{
		Next: nil,
		Data: data,
	}
}
