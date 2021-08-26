package heap

import "fmt"

type MinHeap interface {
	Insert(data int)
	Delete(data int)
	ExtractMin() int
	GetMin() int
	Size() int
}

type minHeap struct {
	items    []int
	capacity int
}

func NewMinHeap(cap int) MinHeap {
	return &minHeap{
		items:    make([]int, 0),
		capacity: cap,
	}
}

func NewMinHeapWithItems(items []int) MinHeap {
	h := minHeap{
		items:    items,
		capacity: len(items),
	}

	for i := len(items)/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}

	return &h
}

func (h *minHeap) heapify(i int) {
	n := len(h.items)

	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && h.items[left] < h.items[smallest] {
		smallest = left
	}

	if right < n && h.items[right] < h.items[smallest] {
		smallest = right
	}

	if smallest != i {
		h.items[i], h.items[smallest] = h.items[smallest], h.items[i]
		h.heapify(smallest)
	}
}

func (h *minHeap) bubbleUp(index int) {
	for h.items[index] < h.items[parent(index)] {
		h.items[index], h.items[parent(index)] = h.items[parent(index)], h.items[index]
		index = parent(index)
	}
}

func (h *minHeap) Insert(data int) {
	if len(h.items) >= h.capacity {
		return
	}

	h.items = append(h.items, data)

	n := len(h.items) - 1

	h.bubbleUp(n)
}

func (h *minHeap) Delete(data int) {
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
	h.items = h.items[:n]

	h.heapify(idx)
}

func (h *minHeap) ExtractMin() int {
	if len(h.items) == 0 {
		return -1
	}

	minValue := h.items[0]

	n := len(h.items) - 1
	// Swap ith with the last element
	h.items[0], h.items[n] = h.items[n], h.items[0]
	// Remove the last element
	h.items = h.items[:n]

	h.heapify(0)

	return minValue
}

func (h *minHeap) GetMin() int {
	return h.items[0]
}

func (h *minHeap) Size() int {
	return len(h.items)
}

func (h *minHeap) String() string {
	var result string

	result += "\n########### MIN HEAP ###########\n"

	for _, item := range h.items {
		result += fmt.Sprintf("[%d]-", item)
	}

	result += "\n########### ######## ###########"

	return result
}
