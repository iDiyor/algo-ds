package bst

type Tree interface {
	Insert(data int)
	Delete(data int)
	Search(data int) *Node
	Size() int
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

func deleteNode(root *Node, data int) *Node {
	if root == nil {
		return root
	}

	// find the node to be deleted
	if data < root.Data {
		root.Left = deleteNode(root.Left, data)
	} else if data > root.Data {
		root.Right = deleteNode(root.Right, data)
	} else {
		// if key is same as root's
		// key, then This is the
		// node to be deleted
		// CASE 1: node has two children
		if root.Left != nil && root.Right != nil {
			// if the node has two children
			// Place the inorder successor in position of
			// the node to be deleted
			// node with two children: Get the inorder
			// successor (smallest in the right subtree)
			root.Data = minValue(root.Right)
			// delete the inorder successor
			root.Right = deleteNode(root.Right, root.Data)
		} else {
			// CASE 2: node is a leaf node
			if root.Left == nil && root.Right == nil {
				root = nil
			} else if root.Left == nil { // CASE 3: node has one or no child
				root = root.Right
			} else {
				root = root.Left
			}
		}
	}

	return root
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

func (t *treeImpl) Search(data int) *Node {
	return searchNode(t.root, data)
}

func searchNode(root *Node, data int) *Node {
	if root == nil {
		return root
	}

	if root.Data == data {
		return root
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
