package singly_linked_list

import "fmt"

type List struct {
	head   *Node
	tail   *Node
	length int
}

func NewList() *List {
	return &List{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// Append()
// Prepend()
// GetAtIndex()
// RemoveAtIndex()
// RemoveHead()
// RemoveTail()
// Size()
// Reverse()
func (l *List) Append(data int) {
	if l == nil {
		return
	}

	if l.head == nil {
		n := NewNode(data)
		l.head = n
		l.length++
		return
	}

	current := l.head
	for current.Next != nil {
		current = current.Next
	}

	newTail := NewNode(data)
	current.Next = newTail

	l.tail = newTail
	l.length++
}

func (l *List) Prepend(data int) {
	if l == nil {
		return
	}

	if l.head == nil {
		n := NewNode(data)
		l.head = n
		l.length++
		return
	}

	newHead := NewNode(data)
	current := l.head

	current.Next = newHead
	l.head = newHead
	l.length++
}

func (l *List) RemoveHead() {
	if l == nil {
		return
	}

	if l.head == nil {
		return
	}

	current := l.head
	l.head = current.Next
	l.length--
}

func (l *List) RemoveTail() {
	if l == nil {
		return
	}

	if l.tail == nil {
		return
	}

	if l.head.Next == nil {
		l.RemoveHead()
		return
	}

	current := l.head

	for current.Next.Next != nil {
		current = current.Next
	}

	l.tail = current
	current.Next = nil
	l.length--
}

func (l *List) Size() int {
	return l.length
}

func (l *List) Reverse() {

}

func (l *List) String() string {
	if l == nil {
		return ""
	}

	if l.head == nil {
		return ""
	}

	current := l.head
	result := ""

	for current.Next != nil {
		result += fmt.Sprintf("[%d]->", current.Data)
		current = current.Next
	}

	return result
}
