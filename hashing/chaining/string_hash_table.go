package chaining

import (
	"fmt"
	"log"
	"math"
	"strings"
)

type Node struct {
	Next *Node
	Prev *Node
	Data string
}

func NewNode(data string) *Node {
	return &Node{
		Next: nil,
		Prev: nil,
		Data: data,
	}
}

type HashTable struct {
	hashTable map[int]*Node
	size      int
	capacity  int
}

func NewHashTable(cap int) *HashTable {
	return &HashTable{
		hashTable: make(map[int]*Node, cap),
		size:      0,
		capacity:  cap,
	}
}

func (h *HashTable) Insert(data string) {
	if h.size >= h.capacity {
		return
	}

	index := h.hash(data)
	h.size++

	log.Printf("insert %s at %d index", data, index)

	if h.hashTable[index] == nil {
		h.hashTable[index] = NewNode(data)
		return
	}

	head := h.hashTable[index]
	newNode := NewNode(data)
	newNode.Next = head
	head.Prev = newNode
	h.hashTable[index] = newNode
}

func (h *HashTable) Delete(data string) bool {
	index := h.hash(data)

	if h.hashTable[index] == nil {
		return false
	}

	log.Printf("delete %s at %d index", data, index)

	head := h.hashTable[index]

	if head.Data == data {
		h.hashTable[index] = head.Next
		return true
	}

	current := head

	for current != nil {
		if current.Data == data {
			break
		}
		current = current.Next
	}

	if current == head {
		h.hashTable[index] = nil
		return true
	}

	if current.Prev != nil {
		current.Prev.Next = current.Next
	}

	if current.Next != nil {
		current.Next.Prev = current.Prev
	}

	h.size--

	return true
}

func (h *HashTable) Search(data string) bool {
	index := h.hash(data)

	if h.hashTable[index] == nil {
		return false
	}

	head := h.hashTable[index]

	current := head
	for current != nil {
		if current.Data == data {
			return true
		}

		current = current.Next
	}

	return false
}

func (h *HashTable) hash(s string) int {
	var hash int
	for idx, char := range s {
		hash += int(char) + int(math.Pow(float64(26), float64(len(s)-idx+1)))
	}
	return hash % h.capacity
}

func (h *HashTable) String() string {
	var result strings.Builder

	result.WriteString("\n###### HASH TABLE ######\n")
	for k, v := range h.hashTable {
		result.WriteString(fmt.Sprintf("[%d]->", k))
		if v != nil {
			i := v
			for i != nil {
				result.WriteString(fmt.Sprintf("%s->", i.Data))
				i = i.Next
			}
		}
		result.WriteString("\n")
	}

	result.WriteString("###### END ######")

	return result.String()
}
