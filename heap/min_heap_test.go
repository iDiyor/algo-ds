package heap

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap(10)

	testNumbers := []int{
		7, 6, 5, 4, 3, 2, 1,
	}

	testExtracts := []int{
		1, 2, 3, 4, 5, 6, 7,
	}

	for _, num := range testNumbers {
		heap.Insert(num)
	}

	if heap.Size() != len(testNumbers) {
		t.Errorf("expected size %d; got %d", len(testNumbers), heap.Size())
	}

	// Extract Min
	for _, extract := range testExtracts {
		min := heap.ExtractMin()
		if min != extract {
			t.Errorf("expected %d; got %d", extract, min)
		}
	}

	for _, num := range testNumbers {
		heap.Insert(num)
	}

	heap.Delete(1)

	testExtracts = []int{
		2, 3, 4, 5, 6, 7,
	}

	// Extract Min
	for _, extract := range testExtracts {
		min := heap.ExtractMin()
		if min != extract {
			t.Errorf("expected %d; got %d", extract, min)
		}
	}
}
