package sorting

func SelectionSort(array []int) {
	for i := 0; i < len(array); i++ {
		minIdx := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIdx] {
				minIdx = j
			}
		}
		temp := array[i]
		array[i] = array[minIdx]
		array[minIdx] = temp
	}
}
