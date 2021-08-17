package sorting

func InsertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		key := array[i]

		// Move elements of arr[0..i-1], that are
		// greater than key, to one position ahead
		// of their current position
		j := i - 1

		for j >= 0 && key < array[j] {
			array[j+1] = array[j]
			j--
		}

		array[j+1] = key
	}
}
