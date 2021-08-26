package heap

import "testing"

func TestPriorityQueue(t *testing.T) {
	heap := NewPriorityQueue(10)

	testNumbers := []int{
		5, 7, 6, 2, 3, 4, 1,
	}

	testDequeue := []int{
		7, 6, 5, 4, 3, 2, 1,
	}

	for _, num := range testNumbers {
		heap.Enqueue(num)
	}

	// Dequeue
	for _, extract := range testDequeue {
		max := heap.Dequeue()
		if max != extract {
			t.Errorf("expected %d; got %d", extract, max)
		}
	}
}
