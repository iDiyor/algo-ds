package btree

import "testing"

func TestTree(t *testing.T) {
	tests := []int{
		30, 40, 50, 60, 70, 80, 90, 100,
	}

	tree := NewTree()
	for _, test := range tests {
		tree.Insert(test)
	}
}
