package chaining

import (
	"log"
	"testing"
)

func TestStringHashTable(t *testing.T) {
	tests := []string{
		"Go", "Python", "Swift", "Javascript", "Go 2", "Python 3", "Swift 3", "Swift 4", "Swift 5", "Python 2",
	}

	hashTable := NewHashTable(10)

	for _, test := range tests {
		hashTable.Insert(test)
	}

	log.Println(hashTable)

	for _, test := range tests {
		if !hashTable.Search(test) {
			t.Errorf("%s not found", test)
		}
	}

	for _, test := range tests {
		if !hashTable.Delete(test) {
			t.Errorf("%s not found and deleted", test)
		}
	}

	log.Println(hashTable)
}
