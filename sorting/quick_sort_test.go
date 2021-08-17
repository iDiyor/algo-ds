package sorting

import "testing"

func TestQuickSort(t *testing.T) {
	input := []int{9, 5, 1, 4, 3, 0}
	expectedOutput := []int{0, 1, 3, 4, 5, 9}

	QuickSort(input, 0, len(input)-1)

	for idx := range input {
		if expectedOutput[idx] != input[idx] {
			t.Errorf("expected %d at %d; got %d", expectedOutput[idx], idx, input[idx])
		}
	}
}
