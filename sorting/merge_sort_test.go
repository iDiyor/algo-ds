package sorting

import "testing"

func TestMergeSort(t *testing.T) {
	input := []int{9, 5, 1, 4, 3, 0}
	expectedOutput := []int{0, 1, 3, 4, 5, 9}

	output := MergeSort(input)

	for idx := range output {
		if expectedOutput[idx] != output[idx] {
			t.Errorf("expected %d at %d; got %d", expectedOutput[idx], idx, output[idx])
		}
	}
}
