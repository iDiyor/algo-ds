package stack

import (
	"log"
	"testing"
)

func TestStackLinkedListImpl(t *testing.T) {
	stack := NewStackLL()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	log.Println(stack)

	stack.Pop()
	stack.Pop()
	stack.Pop()

	log.Println(stack)

	if stack.Size != 2 {
		t.Errorf("expected size = %d; got %d", 2, stack.Size)
	}
}

func TestStackArrayImpl(t *testing.T) {
	stack := NewStack(10)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	log.Println(stack)

	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()

	log.Println(stack)

	if !stack.IsEmpty() {
		t.Errorf("expected stack to be empty")
	}
}
