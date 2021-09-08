package bst

import "testing"

func TestBinarySearchTree(t *testing.T) {
	tests := []int{
		5, 10, 13, 9, 6, 20, 17, 11,
	}

	tree := NewTree()

	for _, test := range tests {
		tree.Insert(test)
	}

	t.Logf("pre-order: %s", tree.PreOrder())
	t.Logf("in-order: %s", tree.InOrder())
	t.Logf("post-order: %s", tree.PostOrder())

	if tree.Size() != len(tests) {
		t.Errorf("expected size %d; got %d", len(tests), tree.Size())
	}

	for _, test := range tests {
		node := tree.Search(test)
		if node != nil && node.Data != test {
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
