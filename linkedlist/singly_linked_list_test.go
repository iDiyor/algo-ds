package linkedlist

import "testing"

func TestSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	list.InsertAtHead(1)
	list.InsertAtTail(3)
	list.InsertAtTail(8)
	list.InsertAtTail(6)
	list.InsertAtTail(7)
	list.InsertAtTail(4)

	list.InsertAtIndex(1, 5)

	t.Log(list)

	valAtFirstIdx := list.Get(1)
	if valAtFirstIdx != 5 {
		t.Errorf("expected %d at %d, got %d", 2, 1, valAtFirstIdx)
	}

	list.Sort()
	t.Log(list)

}
