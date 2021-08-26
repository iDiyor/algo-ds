package linkedlist

import (
	"fmt"
)

type SinglyLinkedListNode struct {
	Next *SinglyLinkedListNode
	Data int
}

func (head *SinglyLinkedListNode) Size() int {
	if head == nil {
		return 0
	}

	current := head
	var size int

	for current != nil {
		current = current.Next
		size++
	}

	return size
}

func (head *SinglyLinkedListNode) String() string {
	var result string

	current := head
	result += fmt.Sprintf("[%d]->", current.Data)

	for current.Next != nil {
		current = current.Next
		result += fmt.Sprintf("[%d]->", current.Data)
	}

	return result
}

func NewSinglyLinkedListNode(data int) *SinglyLinkedListNode {
	return &SinglyLinkedListNode{
		Next: nil,
		Data: data,
	}
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head: nil,
	}
}

func (l *SinglyLinkedList) InsertAtHead(data int) {
	if l.head == nil {
		l.head = NewSinglyLinkedListNode(data)
		return
	}

	newNode := NewSinglyLinkedListNode(data)
	newNode.Next = l.head
	l.head = newNode
}

func (l *SinglyLinkedList) InsertAtTail(data int) {
	if l.head == nil {
		l.head = NewSinglyLinkedListNode(data)
		return
	}

	tail := l.head

	for tail.Next != nil {
		tail = tail.Next
	}

	tail.Next = NewSinglyLinkedListNode(data)
}

func (l *SinglyLinkedList) InsertAtIndex(index, val int) {
	if index < 0 || index > l.size() {
		return
	}

	if index == 0 {
		l.InsertAtHead(val)
		return
	}

	if index == l.size() {
		l.InsertAtTail(val)
		return
	}

	current := l.head
	var currentIdx int

	for current != nil && currentIdx < index-1 {
		current = current.Next
		currentIdx++
	}

	if currentIdx == index-1 && current.Next != nil {
		newNode := NewSinglyLinkedListNode(val)
		newNode.Next = current.Next
		current.Next = newNode
	}
}

func (l *SinglyLinkedList) Delete(data int) {
	current := l.head

	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
}

func (l *SinglyLinkedList) DeleteHead() {
	l.head = l.head.Next
}

func (l *SinglyLinkedList) DeleteTail() {
	preTail := l.head

	for preTail.Next.Next != nil {
		preTail = preTail.Next
	}

	preTail.Next = nil
}

func (l *SinglyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > l.size() {
		return
	}

	if index == 0 {
		l.DeleteHead()
		return
	}

	if index == l.size()-1 {
		l.DeleteTail()
		return
	}

	current := l.head
	var currentIdx int

	for current.Next != nil && currentIdx < index-1 {
		current = current.Next
		currentIdx++
	}

	if currentIdx == index-1 && current.Next != nil {
		current.Next = current.Next.Next
	}
}

func (l *SinglyLinkedList) Get(index int) int {
	if l.head == nil {
		return -1
	}

	if index == 0 {
		return l.head.Data
	}

	current := l.head
	var currentIdx int

	for current != nil && currentIdx < index {
		current = current.Next
		currentIdx++
	}

	return current.Data
}

func (l *SinglyLinkedList) Search(data int) *SinglyLinkedListNode {
	current := l.head
	for current.Next != nil {
		if current.Data == data {
			return current
		}
		current = current.Next
	}

	return nil
}

func (l *SinglyLinkedList) Sort() {
	bubbleSort(l.head)
}

// BubbleSort
func bubbleSort(head *SinglyLinkedListNode) {
	if head == nil {
		return
	}

	i := head

	for i != nil {
		j := i.Next

		for j != nil {
			if i.Data > j.Data {
				i.Data, j.Data = j.Data, i.Data
			}
			j = j.Next
		}
		i = i.Next
	}
}

// MergeSort
func mergeSort(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := getMiddleNode(head)
	left := head
	right := mid.Next
	mid.Next = nil

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func merge(left, right *SinglyLinkedListNode) *SinglyLinkedListNode {

	var result *SinglyLinkedListNode

	if left == nil {
		return right
	} else if right == nil {
		return left
	}

	if left.Data < right.Data {
		result = left
		result.Next = merge(left.Next, right)
	} else {
		result = right
		result.Next = merge(left, right.Next)
	}

	return result
}

func getMiddleNode(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	if head == nil {
		return head
	}

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func (l *SinglyLinkedList) String() string {
	var result string

	current := l.head
	result += fmt.Sprintf("[%d]->", current.Data)

	for current.Next != nil {
		current = current.Next
		result += fmt.Sprintf("[%d]->", current.Data)
	}

	return result
}

func (l *SinglyLinkedList) size() int {
	current := l.head
	var size int

	for current != nil {
		size++
		current = current.Next
	}

	return size
}
