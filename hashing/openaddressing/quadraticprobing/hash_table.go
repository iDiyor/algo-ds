package quadraticprobing

import (
	"fmt"
	"log"
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
	index := h.hash(data)
	log.Printf("insert data %s at index %d", data, index)

	if h.items[index] == "" {
		h.items[index] = data
		h.size++
		return true
	}

	for i := 0; i < h.capacity; i++ {
		index = (index + i*i) % h.capacity
		if h.items[index] == "" {
			h.items[index] = data
			h.size++
			return true
		}
	}

	return false
}

func (h *hashTableImpl) Delete(data string) bool {
	index := h.hash(data)

	if h.items[index] == data {
		h.items[index] = ""
		h.size--
		return true
	}

	for i := 0; i < h.capacity; i++ {
		index = (index + i*i) % h.capacity
		if h.items[index] == data {
			h.items[index] = ""
			h.size--
			return true
		}
	}

	return false
}

func (h *hashTableImpl) Search(data string) bool {
	index := h.hash(data)

	if h.items[index] == data {
		return true
	}

	for i := 0; i < h.capacity; i++ {
		index = (index + i*i) % h.capacity
		if h.items[index] == data {
			return true
		}
	}

	return false
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
