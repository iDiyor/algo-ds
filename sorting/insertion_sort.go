package sorting

func InsertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1

		for j >= 0 && key < array[j] {
			array[j+1] = array[j]
			j--
		}

		array[j+1] = key
	}
}
