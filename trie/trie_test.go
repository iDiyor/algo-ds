package trie

import "testing"

func TestTrie(t *testing.T) {
	tests := []string{
		"abba", "aba", "abc", "aa", "bb", "cc", "abcc", "abbccdd",
	}

	trie := NewTrie()

	for _, test := range tests {
		trie.Insert(test)
	}

	for _, test := range tests {
		if !trie.Search(test) {
			t.Errorf("%s not found", test)
		}
	}
}
