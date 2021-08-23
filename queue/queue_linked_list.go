package queue

type node struct {
	Data int
	Next *node
}

func newNode(data int) *node {
	return &node{
		Data: data,
		Next: nil,
	}
}

type QueueLinkedList struct {
	Head *node
	size int
}

func NewQueueLinkedList() *QueueLinkedList {
	return &QueueLinkedList{
		Head: nil,
		size: 0,
	}
}

func (q *QueueLinkedList) Enqueue(data int) {
	if q.Head == nil {
		q.Head = newNode(data)
		return
	}

	tail := q.Head

	for tail.Next != nil {
		tail = tail.Next
	}

	tail.Next = newNode(data)
	q.size++
}

func (q *QueueLinkedList) Dequeue() int {
	if q.Head == nil {
		return -1
	}

	data := q.Head.Data
	q.Head = q.Head.Next
	q.size--
	return data
}
