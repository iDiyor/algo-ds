package sorting

// Best - O(n*log n)
// Worst - O(n2)
// Average - O(n*log n)
// Space Complexity - O(log n)
// Stability - No

func QuickSort(array []int, low int, high int) {
	if low > high {
		return
	}

	pivot := partition(array, low, high)
	QuickSort(array, low, pivot-1)
	QuickSort(array, pivot+1, high)
}

func partition(array []int, low int, high int) int {
	// choose the rightmost element as pivot
	pivot := array[high]
	// pointer for greater element
	i := low - 1

	// traverse through all elements
	// compare each element with pivot
	for j := low; j < high; j++ {
		if array[j] < pivot {
			// if element smaller than pivot is found
			// swap it with the greater element pointed by i
			i++

			// swapping element at i with element at j
			array[i], array[j] = array[j], array[i]
		}
	}

	// swap the pivot element with the greater element specified by i
	array[i+1], array[high] = array[high], array[i+1]

	// return the position from where partition is done
	return i + 1
}
