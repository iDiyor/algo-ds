package sorting

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	input := []int{9, 5, 1, 4, 3, 0}
	expectedOutput := []int{0, 1, 3, 4, 5, 9}

	SelectionSort(input)

	for idx := range input {
		if expectedOutput[idx] != input[idx] {
			t.Errorf("expected %d at %d; got %d", expectedOutput[idx], idx, input[idx])
		}
	}
}
