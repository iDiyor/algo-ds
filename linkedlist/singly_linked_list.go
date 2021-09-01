package linkedlist

import (
	"fmt"
)

type SLLNode struct {
	Next *SLLNode
	Data int
}

func (head *SLLNode) Size() int {
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

func (head *SLLNode) String() string {
	var result string

	current := head
	result += fmt.Sprintf("[%d]->", current.Data)

	for current.Next != nil {
		current = current.Next
		result += fmt.Sprintf("[%d]->", current.Data)
	}

	return result
}

func NewSLLNode(data int) *SLLNode {
	return &SLLNode{
		Next: nil,
		Data: data,
	}
}

type SLLList struct {
	head *SLLNode
}

func NewSLLList() *SLLList {
	return &SLLList{
		head: nil,
	}
}

func (l *SLLList) InsertAtHead(data int) {
	if l.head == nil {
		l.head = NewSLLNode(data)
		return
	}

	newNode := NewSLLNode(data)
	newNode.Next = l.head
	l.head = newNode
}

func (l *SLLList) InsertAtTail(data int) {
	if l.head == nil {
		l.head = NewSLLNode(data)
		return
	}

	tail := l.head

	for tail.Next != nil {
		tail = tail.Next
	}

	tail.Next = NewSLLNode(data)
}

func (l *SLLList) InsertAtIndex(index, val int) {
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
		newNode := NewSLLNode(val)
		newNode.Next = current.Next
		current.Next = newNode
	}
}

func (l *SLLList) Delete(data int) {
	current := l.head

	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
}

func (l *SLLList) DeleteHead() {
	l.head = l.head.Next
}

func (l *SLLList) DeleteTail() {
	preTail := l.head

	for preTail.Next.Next != nil {
		preTail = preTail.Next
	}

	preTail.Next = nil
}

func (l *SLLList) DeleteAtIndex(index int) {
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

func (l *SLLList) Get(index int) int {
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

func (l *SLLList) Search(data int) *SLLNode {
	current := l.head
	for current.Next != nil {
		if current.Data == data {
			return current
		}
		current = current.Next
	}

	return nil
}

func (l *SLLList) Sort() {
	bubbleSort(l.head)
}

// BubbleSort
func bubbleSort(head *SLLNode) {
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
func mergeSort(head *SLLNode) *SLLNode {
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

func merge(left, right *SLLNode) *SLLNode {

	var result *SLLNode

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

func getMiddleNode(head *SLLNode) *SLLNode {
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

func (l *SLLList) String() string {
	var result string

	current := l.head
	result += fmt.Sprintf("[%d]->", current.Data)

	for current.Next != nil {
		current = current.Next
		result += fmt.Sprintf("[%d]->", current.Data)
	}

	return result
}

func (l *SLLList) size() int {
	current := l.head
	var size int

	for current != nil {
		size++
		current = current.Next
	}

	return size
}
