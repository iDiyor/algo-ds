package heap

import (
	"fmt"
)

type MaxHeap interface {
	Insert(data int)
	Delete(data int)
	ExtractMax() int
	GetMax() int
	Size() int
}

type maxHeap struct {
	items []int
}

func NewMaxHeap() MaxHeap {
	return &maxHeap{
		items: make([]int, 0),
	}
}

func (h *maxHeap) heapify(i int) {
	n := len(h.items)

	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && h.items[left] > h.items[largest] {
		largest = left
	}

	if right < n && h.items[right] > h.items[largest] {
		largest = right
	}

	if largest != i {
		h.items[i], h.items[largest] = h.items[largest], h.items[i]
		h.heapify(largest)
	}
}

func (h *maxHeap) Insert(data int) {
	h.items = append(h.items, data)

	for i := len(h.items)/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *maxHeap) Delete(data int) {
	var idx int

	for i := 0; i < len(h.items); i++ {
		if h.items[i] == data {
			idx = i
			break
		}
	}

	n := len(h.items) - 1

	// Swap ith with the last element
	h.items[idx], h.items[n] = h.items[n], h.items[idx]
	// Remove the last element
	h.items = append(h.items[:n], h.items[n+1:]...)

	for i := len(h.items)/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *maxHeap) ExtractMax() int {
	if len(h.items) == 0 {
		return -1
	}

	maxValue := h.items[0]

	n := len(h.items) - 1

	// Swap ith with the last element
	h.items[0], h.items[n] = h.items[n], h.items[0]
	// Remove the last element
	h.items = h.items[:n]

	h.heapify(0)

	return maxValue
}

func (h *maxHeap) GetMax() int {
	return h.items[0]
}

func (h *maxHeap) String() string {
	return fmt.Sprintf("\nitems - %v", h.items)
}

func (h *maxHeap) Size() int {
	return len(h.items)
}
