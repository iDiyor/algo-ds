package btree

const (
	T = 3 // order
)

type Tree interface {
	Search(key int) bool
	Insert(key int)
}

type treeImpl struct {
	root *node
}

func NewTree() Tree {
	r := newNode(T)
	r.leaf = true

	return &treeImpl{
		root: r,
	}
}

func (t *treeImpl) Search(key int) bool {
	return t.search(t.root, key) != nil
}

func (t *treeImpl) search(node *node, key int) *node {
	if node == nil {
		return node
	}

	i := 0

	for i < len(node.keys) && key < node.keys[i] {
		i++
	}

	if i < len(node.keys) && key == node.keys[i] {
		return node
	} else if node.leaf {
		return nil
	}

	return t.search(node.children[i], key)
}

func (t *treeImpl) Insert(key int) {
	r := t.root
	if r.n == 2*T-1 {
		x := newNode(T)
		t.root = x
		x.children[0] = x
		t.split(0, x, r)
		t.insertNonFull(x, key)
	} else {
		t.insertNonFull(r, key)
	}
}

func (t *treeImpl) split(i int, x, y *node) {
	z := newNode(T)
	z.leaf = y.leaf
	z.n = T - 1

	// copy [T:] keys into z
	for j := 0; j < T-1; j++ {
		z.keys[j] = y.keys[j+T]
	}

	// copy [T:] children into z
	if !y.leaf {
		for j := 0; j < T; j++ {
			z.children[j] = y.children[j+T]
		}
	}

	y.n = T - 1

	// allocate a spot for z
	for j := x.n; j >= i+1; j-- {
		x.children[j+1] = x.children[j]
	}
	x.children[i+1] = z

	// allocate a spot for key in y at T-1
	for j := x.n - 1; j >= i; j-- {
		x.keys[j+1] = x.keys[j]
	}

	// set the key
	x.keys[i] = y.keys[T-1]
	// increase keys count
	x.n++
}

func (t *treeImpl) insertNonFull(x *node, key int) {
	i := x.n - 1
	if x.leaf {
		for i >= 0 && x.keys[i] > key {
			x.keys[i+1] = x.keys[i]
			i--
		}
		i++
		x.keys[i] = key
		x.n++
	} else {
		for i >= 0 && x.keys[i] > key {
			i--
		}
		i++
		if x.children[i].n == 2*T-1 {
			t.split(i, x, x.children[i])
			if key > x.keys[i] {
				i++
			}
		}
		t.insertNonFull(x.children[i], key)
	}
}
