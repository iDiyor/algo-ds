package stack

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func NewNode(data int) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}

type StackLL struct {
	Head *Node
	Size int
}

func NewStackLL() *StackLL {
	return &StackLL{
		Head: nil,
		Size: 0,
	}
}

func (s *StackLL) Push(data int) {
	if s.Head == nil {
		s.Head = NewNode(data)
		s.Size++
		return
	}

	node := NewNode(data)
	node.Next = s.Head
	s.Head = node
	s.Size++
}

func (s *StackLL) Pop() int {
	if s.Head == nil {
		return -1
	}

	head := s.Head
	s.Head = head.Next
	s.Size--
	return head.Data
}

func (s *StackLL) String() string {
	var result string

	current := s.Head

	result += fmt.Sprintln("\n########STACK########")

	for current != nil {
		result += fmt.Sprintf("[%d]\n", current.Data)
		current = current.Next
	}

	result += fmt.Sprintln("########################")

	return result
}
