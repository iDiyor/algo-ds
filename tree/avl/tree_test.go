package avl

import "testing"

func TestTree(t *testing.T) {
	tests := []int{
		13, 10, 5, 4, 6, 11, 15, 16,
	}

	tree := NewTree()

	for _, test := range tests {
		tree.Insert(test)
	}

	if tree.Size() != len(tests) {
		t.Errorf("expected size %d; got %d", len(tests), tree.Size())
	}

	if tree.Search(44) {
		t.Errorf("not expected to find %d; but found", 44)
	}

	for _, test := range tests {
		if !tree.Search(test) {
			t.Errorf("%d not found", test)
		}
	}

	for _, test := range tests {
		tree.Delete(test)
	}

	if tree.Size() != 0 {
		t.Errorf("expected size %d; got %d", 0, tree.Size())
	}
}
