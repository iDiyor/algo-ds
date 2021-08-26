package heap

type PriorityQueue struct {
	items    []int
	capacity int
}

func NewPriorityQueue(cap int) *PriorityQueue {
	return &PriorityQueue{
		items:    make([]int, 0),
		capacity: cap,
	}
}

func (q *PriorityQueue) heapify(i int) {
	n := len(q.items)
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && q.items[left] > q.items[largest] {
		largest = left
	}

	if right < n && q.items[right] > q.items[largest] {
		largest = right
	}

	if largest != i {
		q.items[i], q.items[largest] = q.items[largest], q.items[i]
		q.heapify(largest)
	}
}

func (q *PriorityQueue) bubbleUp(index int) {
	parent := (index - 1) / 2
	for q.items[index] > q.items[parent] {
		q.items[index], q.items[parent] = q.items[parent], q.items[index]
		index = (index - 1) / 2
	}
}

func (q *PriorityQueue) Enqueue(data int) bool {
	if len(q.items) >= q.capacity {
		return false
	}

	q.items = append(q.items, data)
	q.bubbleUp(len(q.items) - 1)

	return true
}

func (q *PriorityQueue) Peek() int {
	if len(q.items) == 0 {
		return -1
	}

	return q.items[0]
}

func (q *PriorityQueue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}

	maxValue := q.items[0]

	n := len(q.items) - 1

	// swap root with the last node
	q.items[0], q.items[n] = q.items[n], q.items[0]
	// remove the last node
	q.items = q.items[:n]

	q.heapify(0)

	return maxValue
}
