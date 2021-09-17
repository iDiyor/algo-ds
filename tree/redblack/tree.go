package redblack

import "fmt"

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

	var p *node // going to be parent of x
	r := t.root

	for r != nil {
		p = r
		if x.key < r.key {
			r = r.left
		} else {
			r = r.right
		}
	}

	x.parent = p
	if p == nil {
		t.root = x
	} else if x.key < p.key {
		p.left = x
	} else {
		p.right = x
	}

	// if new node is a root node, simply return
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
					rotateLeft(x)
				}
				x.parent.color = colorBlack
				x.parent.parent.color = colorRed
				rotateRight(x.parent.parent)
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
					rotateRight(x)
				}
				x.parent.color = colorBlack
				x.parent.parent.color = colorRed
				rotateLeft(x.parent.parent)
			}
		}

		if x == t.root {
			break
		}
	}

	t.root.color = colorBlack
}

func (t *treeImpl) Delete(key int) {
	if t.root == nil {
		return
	}

	var x *node

	r := t.root

	for r != nil {
		if r.key == key {
			x = r
			break
		} else if key < r.key {
			r = r.left
		} else {
			r = r.right
		}
	}

	// a node containing the key is not found
	if x == nil {
		return
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
