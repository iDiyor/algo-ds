package sorting

func heapify(array []int, n int, i int) {
	// Find largest among root, left child and right child
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && array[l] > array[largest] {
		largest = l
	}

	if r < n && array[r] > array[largest] {
		largest = r
	}

	// Swap and continue heapifying if root is not largest
	if largest != i {
		array[i], array[largest] = array[largest], array[i]
		heapify(array, n, largest)
	}
}

func HeapSort(array []int) {
	n := len(array)

	// build a max heap
	for i := (n / 2) - 1; i >= 0; i-- {
		heapify(array, n, i)
	}

	// Heap sort; replace last with root and heapify; reduce the heap size (value of i passed to heapify)
	for i := n - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		heapify(array, i, 0)
	}
}
