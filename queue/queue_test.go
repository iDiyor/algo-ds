package queue

import (
	"testing"
)

func TestQueueArray(t *testing.T) {
	tests := [][]int{
		{1, 2, 3, 4, 5, 6},
		{4, 5, 2, 1, 3, 0},
		{100, 245, 378, 4223, 1124, 5462},
	}

	queue := NewQueueArray(10)
	for _, test := range tests {
		for item := range test {
			queue.Enqueue(item)
		}

		for item := range test {
			data := queue.Dequeue()
			if item != data {
				t.Errorf("expected %d; got %d", item, data)
			}
		}

		emptyQueueData := queue.Dequeue()
		if emptyQueueData != -1 {
			t.Errorf("expected %d; got %d", -1, emptyQueueData)
		}
	}
}

func TestQueueLinkedList(t *testing.T) {
	tests := [][]int{
		{1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 15, 16, 122, 523},
		{4, 5, 2, 1, 3, 0},
		{100, 245, 378, 4223, 1124, 5462},
	}

	queue := NewQueueLinkedList()
	for _, test := range tests {
		for item := range test {
			queue.Enqueue(item)
		}

		for item := range test {
			data := queue.Dequeue()
			if item != data {
				t.Errorf("expected %d; got %d", item, data)
			}
		}

		emptyQueueData := queue.Dequeue()
		if emptyQueueData != -1 {
			t.Errorf("expected %d; got %d", -1, emptyQueueData)
		}
	}
}
