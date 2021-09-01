package hashing

type HashTable interface {
	Insert(data string) bool
	Delete(data string) bool
	Search(data string) bool
	Size() int
}
