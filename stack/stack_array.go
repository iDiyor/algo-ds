package stack

import "fmt"

type Stack struct {
	Items []int
	top   int
	size  int
}

func NewStack(size int) *Stack {
	return &Stack{
		Items: make([]int, size),
		top:   -1,
		size:  0,
	}
}

func (s *Stack) Push(data int) bool {
	if s.IsFull() {
		return false
	}

	s.top++
	s.Items[s.top] = data
	s.size++
	return true
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	data := s.Items[s.top]
	s.Items[s.top] = 0
	s.top--
	s.size--
	return data
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) IsFull() bool {
	return s.top == cap(s.Items)-1
}

func (s *Stack) String() string {
	var result string

	result += fmt.Sprintln("\n########STACK########")

	for _, item := range s.Items {
		result += fmt.Sprintf("[%d]\n", item)
	}

	result += fmt.Sprintln("\n########STACK########")

	return result
}
