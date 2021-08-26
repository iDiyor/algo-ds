package heap

import (
	"fmt"
)

type maxHeapWithCap struct {
	items    []int
	size     int
	capacity int
}

func NewMaxHeapWithCap(cap int) MaxHeap {
	return &maxHeapWithCap{
		items:    make([]int, cap),
		size:     0,
		capacity: cap,
	}
}

func (h *maxHeapWithCap) heapify(i int) {
	n := h.size

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

func (h *maxHeapWithCap) Insert(data int) {
	if h.size == h.capacity {
		return
	}

	h.items[h.size] = data
	h.size++

	for i := h.size/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *maxHeapWithCap) Delete(data int) {
	var idx int

	for i := 0; i < h.size; i++ {
		if h.items[i] == data {
			idx = i
			break
		}
	}

	n := h.size - 1

	// Swap ith with the last element
	h.items[idx], h.items[n] = h.items[n], h.items[idx]
	// Remove the last element
	h.items = append(h.items[:n], h.items[n+1:]...)
	h.size--

	for i := h.size/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *maxHeapWithCap) ExtractMax() int {
	if h.size == 0 {
		return -1
	}

	maxValue := h.items[0]

	n := h.size - 1
	// Swap ith with the last element
	h.items[0], h.items[n] = h.items[n], h.items[0]
	// Remove the last element
	h.items = append(h.items[:n], h.items[n+1:]...)
	h.size--

	h.heapify(0)

	return maxValue
}

func (h *maxHeapWithCap) GetMax() int {
	return h.items[0]
}

func (h *maxHeapWithCap) String() string {
	return fmt.Sprintf("\nitems - %v\nsize - %d", h.items, h.size)
}

func (h *maxHeapWithCap) Size() int {
	return h.size
}
