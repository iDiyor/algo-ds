package quadraticprobing

import (
	"log"
	"testing"
)

func TestHashTable(t *testing.T) {
	tests := []string{
		"Golang", "Python", "Java", "Javascript", "Rust", "OCaml", "C#", "Swift", "Kotlin", "C++",
	}

	hashTable := NewHashTable(17)

	for _, test := range tests {
		if !hashTable.Insert(test) {
			t.Errorf("no empty slot found for %s", test)
		}
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
