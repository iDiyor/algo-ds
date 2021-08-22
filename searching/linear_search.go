package searching

// Linear Search Complexities
// Time Complexity: O(n)
// Space Complexity: O(1)
// For searching operations in smaller arrays (<100 items).
func LinearSearch(array []int, x int) int {
	for idx, el := range array {
		if el == x {
			return idx
		}
	}
	return -1
}
