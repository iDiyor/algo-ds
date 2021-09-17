package redblack

import (
	"log"
	"testing"
)

func TestTree(t *testing.T) {
	tests := []int{
		50, 30, 60, 65, 35, 25, 45, 70, 100, 150, 120, 130, 110, 200, 500, 1000, 600, 555, 444, 333, 220,
	}

	tree := NewTree()

	for _, test := range tests {
		tree.Insert(test)
	}

	log.Printf("AFTER_INSERT: %v", tree.InOrder())

	if tree.Size() != len(tests) {
		t.Errorf("expected size %d; got %d", len(tests), tree.Size())
	}

	if !tree.Search(tests[0]) {
		t.Errorf("%d not found", tests[0])
	}

	if tree.Search(100) {
		t.Errorf("found non-existing key")
	}
}
