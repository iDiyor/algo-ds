package queue

import "fmt"

type QueueArray struct {
	Items []int
	front int
	rear  int
	size  int
}

func NewQueueArray(capacity int) *QueueArray {
	return &QueueArray{
		Items: make([]int, capacity),
		front: -1,
		rear:  -1,
		size:  0,
	}
}

func (q *QueueArray) Enqueue(data int) bool {
	if q.isFull() {
		return false
	}

	if q.front == -1 {
		q.front = 0
	}

	q.rear = (q.rear + 1) % cap(q.Items)
	q.Items[q.rear] = data
	q.size++

	return true
}

func (q *QueueArray) Dequeue() int {
	if q.isEmpty() {
		return -1
	}

	data := q.Items[q.front]
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else {
		q.front = (q.front + 1) % cap(q.Items)
	}

	q.size--
	return data
}

func (q *QueueArray) isEmpty() bool {
	return q.size == 0
}

func (q *QueueArray) isFull() bool {
	return q.size == cap(q.Items)
}

func (q *QueueArray) Size() int {
	return q.size
}

func (q *QueueArray) String() string {
	var result string

	front := q.front

	result += "\n########### QUEUE ###########\n"

	for front <= q.rear {
		result += fmt.Sprintf("[%d]<-", q.Items[front])
		front++
	}

	result += "\n########### END ###########"

	return result
}
