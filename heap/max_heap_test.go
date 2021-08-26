package heap

import (
	"testing"
)

func TestMaxHeapWithCap(t *testing.T) {
	heap := NewMaxHeapWithCap(10)

	testNumbers := []int{
		1, 5, 2, 4, 3, 7, 6,
	}

	testExtracts := []int{
		7, 6, 5, 4, 3, 2, 1,
	}

	for _, num := range testNumbers {
		heap.Insert(num)
	}

	if heap.Size() != len(testNumbers) {
		t.Errorf("expected size %d; got %d", len(testNumbers), heap.Size())
	}

	for _, extract := range testExtracts {
		max := heap.ExtractMax()
		if max != extract {
			t.Errorf("expected %d; got %d", extract, max)
		}
	}
}

func TestMaxHeapSlize(t *testing.T) {
	heap := NewMaxHeap()

	testNumbers := []int{
		1, 5, 2, 4, 3, 7, 6,
	}

	testExtracts := []int{
		7, 6, 5, 4, 3, 2, 1,
	}

	for _, num := range testNumbers {
		heap.Insert(num)
	}

	if heap.Size() != len(testNumbers) {
		t.Errorf("expected size %d; got %d", len(testNumbers), heap.Size())
	}

	for _, extract := range testExtracts {
		max := heap.ExtractMax()
		if max != extract {
			t.Errorf("expected %d; got %d", extract, max)
		}
	}
}
