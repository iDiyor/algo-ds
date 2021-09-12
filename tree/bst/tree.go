package bst

import "fmt"

type Tree interface {
	Insert(data int)
	Delete(data int)
	Search(data int) bool
	Size() int
	PreOrder() string
	InOrder() string
	PostOrder() string
}

type treeImpl struct {
	root   *Node
	height int
	size   int // nodes count
}

func NewTree() Tree {
	return &treeImpl{
		root:   nil,
		height: 0,
		size:   0,
	}
}

func (t *treeImpl) Insert(data int) {
	t.root = insertNode(t.root, data)
	t.size++
}

func insertNode(node *Node, data int) *Node {
	if node == nil {
		return NewNode(data)
	}

	if data < node.Data {
		node.Left = insertNode(node.Left, data)
	} else {
		node.Right = insertNode(node.Right, data)
	}

	return node
}

func (t *treeImpl) Delete(data int) {
	t.root = deleteNode(t.root, data)
	t.size--
}

func deleteNode(node *Node, data int) *Node {
	if node == nil {
		return node
	}

	// find the node to be deleted
	if data < node.Data {
		node.Left = deleteNode(node.Left, data)
	} else if data > node.Data {
		node.Right = deleteNode(node.Right, data)
	} else {
		// if key is same as node's
		// key, then This is the
		// node to be deleted
		// CASE 1: node is a leaf node
		if node.Left == nil && node.Right == nil {
			node = nil
		} else if node.Left != nil && node.Right != nil { // CASE 2: Node has two child nodes
			// if the node has two children
			// Place the inorder successor in position of
			// the node to be deleted
			// node with two children: Get the inorder
			// successor (smallest in the right subtree)
			node.Data = minValue(node.Right)
			node.Right = deleteNode(node.Right, node.Data)
		} else if node.Left == nil { // CASE 3: Node has one child node
			return node.Right
		} else {
			return node.Left
		}
	}

	return node
}

// find the inorder successor
func minValue(root *Node) int {
	minValue := root.Data
	for root.Left != nil {
		minValue = root.Left.Data
		root = root.Left
	}
	return minValue
}

func (t *treeImpl) Search(data int) bool {
	return searchNode(t.root, data)
}

func searchNode(root *Node, data int) bool {
	if root == nil {
		return false
	}

	if root.Data == data {
		return true
	}

	// search left subtree
	if data < root.Data {
		return searchNode(root.Left, data)
	} else { // search right subtree
		return searchNode(root.Right, data)
	}
}

func (t *treeImpl) Size() int {
	return t.size
}

func (t *treeImpl) PreOrder() string {
	return preOrder(t.root)
}

func preOrder(root *Node) string {
	if root == nil {
		return ""
	}

	var result string
	result += fmt.Sprintf("(%d)->", root.Data)
	result += preOrder(root.Left)
	result += preOrder(root.Right)
	return result
}

func (t *treeImpl) InOrder() string {
	return inOrder(t.root)
}

func inOrder(root *Node) string {
	if root == nil {
		return ""
	}

	var result string
	result += inOrder(root.Left)
	result += fmt.Sprintf("(%d)->", root.Data)
	result += inOrder(root.Right)
	return result
}

func (t *treeImpl) PostOrder() string {
	return postOrder(t.root)
}

func postOrder(root *Node) string {
	if root == nil {
		return ""
	}

	var result string
	result += postOrder(root.Left)
	result += postOrder(root.Right)
	result += fmt.Sprintf("(%d)->", root.Data)
	return result
}
