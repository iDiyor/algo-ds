package sorting

import (
	"math"
)

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	middle := int(math.Floor(float64(len(array)) / 2))
	leftSubarray := array[:middle]
	rightSubarray := array[middle:]

	leftSubarray = MergeSort(leftSubarray)
	rightSubarray = MergeSort(rightSubarray)

	return merge(leftSubarray, rightSubarray)
}

func merge(leftSubarray []int, rightSubarray []int) []int {
	array := make([]int, len(leftSubarray)+len(rightSubarray))
	var i, j, k int

	for i < len(leftSubarray) && j < len(rightSubarray) {
		if leftSubarray[i] < rightSubarray[j] {
			array[k] = leftSubarray[i]
			i++
		} else {
			array[k] = rightSubarray[j]
			j++
		}
		k++
	}

	for i < len(leftSubarray) {
		array[k] = leftSubarray[i]
		i++
		k++
	}

	for j < len(rightSubarray) {
		array[k] = rightSubarray[j]
		j++
		k++
	}

	return array
}
