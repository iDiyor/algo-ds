package trie

const (
	ALBHABET_SIZE = 26 // total characters in english alphabet
)

type node struct {
	children    []*node
	isEndOfWord bool
}

func newNode() *node {
	return &node{
		children:    make([]*node, ALBHABET_SIZE),
		isEndOfWord: false,
	}
}
