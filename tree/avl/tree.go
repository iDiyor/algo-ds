package avl

import (
	"fmt"
	"math"
)

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
	root *node
	size int
}

func NewTree() Tree {
	return &treeImpl{
		root: nil,
		size: 0,
	}
}

func (t *treeImpl) Insert(key int) {
	t.root = insertNode(t.root, key)
	t.size++
}

func insertNode(node *node, key int) *node {
	if node == nil {
		return newNode(key)
	}

	if key < node.key {
		node.left = insertNode(node.left, key)
	} else {
		node.right = insertNode(node.right, key)
	}

	node.height = 1 + int(math.Max(float64(getHeight(node.left)), float64(getHeight(node.right))))
	node.balanceFactor = getHeight(node.left) - getHeight(node.right)

	if node.balanceFactor > 1 { // height of left subtree is greater than that of right subtree
		if key < node.left.key {
			return rotateRight(node)
		} else {
			node.left = rotateLeft(node.left)
			return rotateRight(node)
		}
	} else if node.balanceFactor < -1 { // height of right subtree is greater than that of left subtree
		if key > node.right.key {
			return rotateLeft(node)
		} else {
			node.right = rotateRight(node.right)
			return rotateLeft(node)
		}
	}

	return node
}

func getHeight(node *node) int {
	if node == nil {
		return 0
	}
	return node.height
}

func (t *treeImpl) Delete(key int) {
	t.root = deleteNode(t.root, key)
	t.size--
}

func deleteNode(node *node, key int) *node {
	if node == nil {
		return node
	}

	if key < node.key {
		node.left = deleteNode(node.left, key)
	} else if key > node.key {
		node.right = deleteNode(node.right, key)
	} else {
		// CASE 1: leaf node
		if node.left == nil && node.right == nil {
			node = nil
		} else if node.left != nil && node.right != nil { // CASE 2: a node has 2 child nodes
			node.key = getMinKey(node.right)
			node.right = deleteNode(node.right, node.key)
		} else if node.left == nil { // CASE 3: a node has 1 child node
			node = node.right
		} else {
			node = node.left
		}
	}

	if node == nil {
		return node
	}

	node.height = 1 + int(math.Max(float64(getHeight(node.left)), float64(getHeight(node.right))))
	node.balanceFactor = getHeight(node.left) - getHeight(node.right)

	if node.balanceFactor > 1 { // height of left subtree is greater than that of right subtree
		if getBalanceFactor(node.left) >= 0 { // left side is heavier
			return rotateRight(node)
		} else { // right side is heavier
			node.left = rotateLeft(node.left)
			return rotateRight(node)
		}
	} else if node.balanceFactor < -1 { // height of right subtree is greater then that of left subtree
		if getBalanceFactor(node.right) <= 0 { // right side is heavier
			return rotateLeft(node)
		} else { // left side is heavier
			node.right = rotateRight(node.right)
			return rotateLeft(node)
		}
	}

	return node
}

func (t *treeImpl) Search(key int) bool {
	return searchNode(t.root, key)
}

func searchNode(root *node, key int) bool {
	if root == nil {
		return false
	}
	if root.key == key {
		return true
	}

	if key < root.key {
		return searchNode(root.left, key) // search in left subtree
	}

	return searchNode(root.right, key) // search in right subtree
}

func (t *treeImpl) Size() int {
	return t.size
}

func (t *treeImpl) PreOrder() string {
	return preOrder(t.root)
}

func preOrder(root *node) string {
	if root == nil {
		return ""
	}

	var result string
	result += fmt.Sprintf("(%d)->", root.key)
	result += inOrder(root.left)
	result += inOrder(root.right)

	return result
}

func (t *treeImpl) InOrder() string {
	return inOrder(t.root)
}

func inOrder(root *node) string {
	if root == nil {
		return ""
	}

	var result string
	result += inOrder(root.left)
	result += fmt.Sprintf("(%d)->", root.key)
	result += inOrder(root.right)

	return result
}

func (t *treeImpl) PostOrder() string {
	return postOrder(t.root)
}

func postOrder(root *node) string {
	if root == nil {
		return ""
	}

	var result string
	result += inOrder(root.left)
	result += inOrder(root.right)
	result += fmt.Sprintf("(%d)->", root.key)

	return result
}

func getMinKey(node *node) int {
	if node == nil {
		return 0
	}

	for node.left != nil {
		node = node.left
	}

	return node.key
}

func getBalanceFactor(node *node) int {
	if node == nil {
		return 0
	}

	return getHeight(node.left) - getHeight(node.right)
}
