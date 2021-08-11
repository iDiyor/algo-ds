package main

import (
	"fmt"

	"github.com/iDiyor/algo-ds/singly_linked_list"
)

func main() {
	singlyLinkedListRoot := singly_linked_list.NewList()
	singlyLinkedListRoot.Append(0)
	singlyLinkedListRoot.Append(1)
	singlyLinkedListRoot.Append(2)
	singlyLinkedListRoot.Append(3)
	singlyLinkedListRoot.Append(4)
	fmt.Printf("Singly Linked List - %s\n", singlyLinkedListRoot)

	singlyLinkedListRoot.RemoveTail()
	fmt.Printf("Singly Linked List - %s\n", singlyLinkedListRoot)

	singlyLinkedListRoot.RemoveHead()
	fmt.Printf("Singly Linked List - %s\n", singlyLinkedListRoot)
}
