package searching

func BinarySearchIterative(array []int, x int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2
		if x == array[mid] {
			return mid
		} else if x > array[mid] {
			low = mid + 1
		} else if x < array[mid] {
			high = mid - 1
		}
	}

	return -1
}

func BinarySearchRecursive(array []int, x int, low int, high int) int {
	if low >= high {
		return -1
	}

	mid := (low + high) / 2

	if x == array[mid] {
		return mid
	}

	if x > array[mid] {
		return BinarySearchRecursive(array, x, mid+1, high)
	} else {
		return BinarySearchRecursive(array, x, low, mid-1)
	}
}
