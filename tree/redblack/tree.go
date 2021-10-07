package redblack

import (
	"fmt"
	"log"
)

type Tree interface {
	Insert(key int)
	Delete(key int)
	Search(key int) bool
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
	t.size++

	x := newNode(key)
	var y *node // going to be parent of x

	r := t.root

	for r != nil {
		y = r
		if x.key < r.key {
			r = r.left
		} else {
			r = r.right
		}
	}

	x.parent = y
	if y == nil {
		t.root = x
	} else if x.key < y.key {
		y.left = x
	} else {
		y.right = x
	}

	// if new node is a root node, set black color and simply return
	if x.parent == nil {
		x.color = colorBlack
		return
	}

	// if the grandparent is null, simply return
	if x.parent.parent == nil {
		return
	}

	t.fixInsert(x)
}

func (t *treeImpl) fixInsert(x *node) {
	for x.parent.isRed() {
		if x.parent == x.parent.parent.left {
			y := x.parent.parent.right // uncle
			if y != nil && y.isRed() {
				y.color = colorBlack
				x.parent.color = colorBlack
				x.parent.parent.color = colorRed // grandparent
				x = x.parent.parent
			} else { // Case: 3.2.4
				if x == x.parent.right {
					x = x.parent
					t.rotateLeft(x)
				}
				x.parent.color = colorBlack
				x.parent.parent.color = colorRed
				t.rotateRight(x.parent.parent)
			}
		} else {
			y := x.parent.parent.left  // uncle
			if y != nil && y.isRed() { // Case: 3.1
				y.color = colorBlack
				x.parent.color = colorBlack
				x.parent.parent.color = colorRed
				x = x.parent.parent
			} else { // Case: 3.2.2
				if x == x.parent.left {
					x = x.parent
					t.rotateRight(x)
				}
				x.parent.color = colorBlack
				x.parent.parent.color = colorRed
				t.rotateLeft(x.parent.parent)
			}
		}

		if x == t.root {
			break
		}
	}

	t.root.color = colorBlack
}

func (t *treeImpl) fixDelete(x *node) {
	if x == nil {
		return
	}

	for x != t.root && x.isBlack() {
		if x == x.parent.left {
			s := x.parent.right
			if s != nil && s.isRed() {
				s.color = colorBlack
				x.parent.color = colorRed
				t.rotateLeft(x.parent)
				s = x.parent.right
			}
			if (s.left == nil || s.left.isBlack()) && (s.right != nil || s.right.isBlack()) {
				s.color = colorRed
				x = x.parent
			} else {
				if s.right.isBlack() {
					s.left.color = colorBlack
					s.color = colorRed
					t.rotateRight(s)
					s = x.parent.right
				}

				s.color = x.parent.color
				s.right.color = colorBlack
				x.parent.color = colorBlack
				t.rotateLeft(x.parent)
				x = t.root
			}
		} else { // x is right child
			s := x.parent.left
			if s != nil && s.isRed() {
				s.color = colorBlack
				x.parent.color = colorRed
				t.rotateRight(x.parent)
				s = x.parent.left
			}
			if (s.left == nil || s.left.isBlack()) && (s.right == nil || s.right.isBlack()) {
				s.color = colorRed
				x = x.parent
			} else {
				if s.left.isBlack() {
					s.right.color = colorBlack
					s.color = colorRed
					t.rotateLeft(s)
					s = s.parent.left
				}

				s.color = x.parent.color
				x.parent.color = colorBlack
				s.left.color = colorBlack
				t.rotateRight(x.parent)
				x = t.root
			}
		}
	}

	x.color = colorBlack
}

func (t *treeImpl) Delete(key int) {
	log.Printf("deleting %d; root: %s", key, t.root)
	t.size--

	if t.root == nil {
		return
	}

	var z, x *node
	r := t.root

	for r != nil {
		if r.key == key {
			z = r
			break
		} else if key < r.key {
			r = r.left
		} else {
			r = r.right
		}
	}

	// a node containing the key is not found
	if z == nil {
		return
	}

	log.Printf("found node: %s", z)

	y := z
	yOriginalColor := y.color
	if z.left == nil {
		x = z.right
		t.rbTransplant(z, z.right)
	} else if z.right == nil {
		x = z.left
		t.rbTransplant(z, z.left)
	} else {
		y = getMin(y.right)
		yOriginalColor = y.color
		x = y.right

		if y.parent == z && x != nil {
			x.parent = y
		} else {
			t.rbTransplant(y, y.right) // (y.parent.left|right) --> (y.right); y is spliced out
			y.right = z.right
			if y.right != nil {
				y.right.parent = y
			}
		}

		t.rbTransplant(z, y)
		y.left = z.left
		if y.left != nil {
			y.left.parent = y
		}
		y.color = z.color
	}

	if yOriginalColor == colorBlack {
		t.fixDelete(x)
	}
}

func (t *treeImpl) Search(key int) bool {
	r := t.root

	for r != nil {
		if r.key == key {
			return true
		} else if key < r.key {
			r = r.left
		} else {
			r = r.right
		}
	}

	return false
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

func (t *treeImpl) rbTransplant(u, v *node) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	if v != nil {
		v.parent = u.parent
	}
}

func getMin(node *node) *node {
	for node.left != nil {
		node = node.left
	}

	return node
}
