package linearprobing

import (
	"fmt"
	"math"
	"strings"

	hashing "github.com/iDiyor/algo-ds/hashing"
)

type hashTableImpl struct {
	items    []string
	size     int
	capacity int
}

func NewHashTable(cap int) hashing.HashTable {
	return &hashTableImpl{
		items:    make([]string, cap),
		size:     0,
		capacity: cap,
	}
}

func (h *hashTableImpl) Insert(data string) bool {
	if h.size >= h.capacity {
		return false
	}

	index := h.hash(data)

	for h.items[index] != "" {
		index = (index + 1) % h.capacity
	}

	h.items[index] = data
	h.size++

	return true
}

func (h *hashTableImpl) Delete(data string) bool {
	index := h.hash(data)

	for h.items[index] != data {
		index = (index + 1) % h.capacity
	}

	if h.items[index] == data {
		h.items[index] = ""
		h.size--
		return true
	}

	return false
}

func (h *hashTableImpl) Search(data string) bool {
	index := h.hash(data)

	for h.items[index] != data {
		index = (index + 1) % h.capacity
	}

	return h.items[index] == data
}

func (h *hashTableImpl) Size() int {
	return h.size
}

func (h *hashTableImpl) String() string {
	var result strings.Builder

	result.WriteString("\n###### HASH TABLE ######\n")
	for idx, v := range h.items {
		result.WriteString(fmt.Sprintf("[%d]->", idx))
		if v != "" {
			result.WriteString(fmt.Sprintf("%s->", v))
		}
		result.WriteString("\n")
	}

	result.WriteString("###### END ######")

	return result.String()
}

func (h *hashTableImpl) hash(data string) int {
	var hash int
	for idx, char := range data {
		hash += int(char) + int(math.Pow(float64(26), float64(len(data)-idx+1)))
	}
	return hash % h.capacity
}
