package linkedlist

type DLLNode struct {
	Data int
	Next *DLLNode
	Prev *DLLNode
}

func NewDLLNode(data int) *DLLNode {
	return &DLLNode{
		Data: data,
		Next: nil,
		Prev: nil,
	}
}

type DLLList struct {
	Head *DLLNode
	Size int
}

func NewDDLList() *DLLList {
	return &DLLList{
		Head: nil,
		Size: 0,
	}
}

// Insertion
// Insert at head
// Insert at the end
// Insert after given node
// Insert before given node

func (l *DLLList) InsertAtHead(data int) bool {
	if l.Head == nil {
		node := NewDLLNode(data)
		l.Head = node
		return true
	}

	node := NewDLLNode(data)
	node.Next = l.Head
	l.Head.Prev = node

	l.Head = node

	return true
}

func (l *DLLList) InsertAtTail(data int) bool {
	if l.Head == nil {
		node := NewDLLNode(data)
		l.Head = node
		return true
	}

	last := l.Head
	for last.Next != nil { // find the last node
		last = last.Next
	}

	node := NewDLLNode(data)
	// assign prev of node to last
	node.Prev = last
	// assign next of last to new node
	last.Next = node

	return true
}

func (l *DLLList) InsertAfter(prevNode *DLLNode, data int) bool {
	if l.Head == nil {
		node := NewDLLNode(data)
		l.Head = node
		return true
	}

	if prevNode == nil {
		return false
	}

	if prevNode.Next == nil {
		return l.InsertAtTail(data)
	}

	nextNode := NewDLLNode(data)

	nextNode.Next = prevNode.Next
	nextNode.Prev = prevNode

	prevNode.Next.Prev = nextNode
	prevNode.Next = nextNode

	return true
}

func (l *DLLList) InsertBefore(nextNode *DLLNode, data int) bool {
	if l.Head == nil {
		node := NewDLLNode(data)
		l.Head = node
		return true
	}

	if nextNode == nil {
		return false
	}

	if nextNode.Prev == nil {
		return l.InsertAtHead(data)
	}

	prevNode := NewDLLNode(data)
	prevNode.Next = nextNode
	prevNode.Prev = nextNode.Prev

	nextNode.Prev.Next = prevNode
	nextNode.Prev = prevNode

	return true
}

// Deletion
// Delete head
// Delete tail
// Delete given node

func (l *DLLList) DeleteAtHead() bool {
	if l.Head == nil {
		return false
	}

	if l.Head.Next != nil {
		l.Head.Next.Prev = nil
	}

	l.Head = l.Head.Next

	return true
}

func (l *DLLList) DeleteAtTail() bool {
	if l.Head == nil {
		return false
	}

	last := l.Head
	for last.Next != nil {
		last = last.Next
	}

	last.Prev.Next = nil
	last.Next = nil

	return true
}

func (l *DLLList) DeleteNode(node *DLLNode) bool {
	if l.Head == nil {
		return false
	}

	if node == nil {
		return false
	}

	if node.Prev == nil {
		return l.DeleteAtHead()
	}

	if node.Next == nil {
		return l.DeleteAtTail()
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	return true
}

func (l *DLLList) DeleteData(data int) bool {
	node := l.Head

	for node.Data != data {
		node = node.Next
	}

	if node == l.Head {
		return false
	}

	return l.DeleteNode(node)
}

func (l *DLLList) DeleteAfter(prevNode *DLLNode) bool {
	if l.Head == nil {
		return false
	}

	if prevNode.Next == nil {
		return false
	}

	delNode := prevNode.Next
	prevNode.Next = delNode.Next
	delNode.Prev = prevNode

	return true
}

func (l *DLLList) DeleteBefore(nextNode *DLLNode) bool {
	if l.Head == nil {
		return false
	}

	if nextNode.Prev == nil {
		return false
	}

	delNode := nextNode.Prev
	nextNode.Prev = delNode.Prev
	delNode.Prev.Next = delNode.Next

	return true
}
