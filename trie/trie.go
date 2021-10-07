package trie

type Trie interface {
	Insert(key string)
	Search(key string) bool
}

type trieImpl struct {
	root *node
}

func NewTrie() Trie {
	return &trieImpl{
		root: newNode(),
	}
}

func (t *trieImpl) Insert(key string) {
	current := t.root
	for i := 0; i < len(key); i++ {
		index := key[i] - 'a'
		if current.children[index] == nil {
			current.children[index] = newNode()
		}
		current = current.children[index]
	}
	current.isEndOfWord = true
}

func (t *trieImpl) Search(key string) bool {
	current := t.root
	for i := 0; i < len(key); i++ {
		index := key[i] - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}

	return current.isEndOfWord
}
