package avl

import "math"

// order if operations inside matters
func rotateLeft(x *node) *node {
	if x == nil {
		return x
	}

	y := x.right
	x.right = y.left

	if y.left != nil {
		y.left.parent = x
	}

	y.parent = x.parent

	if x.parent == nil {
		y.parent = nil
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	x.parent = y
	y.left = x

	x.height = 1 + int(math.Max(float64(getHeight(x.left)), float64(getHeight(x.right))))
	y.height = 1 + int(math.Max(float64(getHeight(y.left)), float64(getHeight(y.right))))

	return y
}

// order if operations inside matters
func rotateRight(x *node) *node {
	if x == nil {
		return x
	}

	y := x.left
	x.left = y.right

	if y.right != nil {
		y.right.parent = x
	}

	y.parent = x.parent

	if x.parent == nil {
		y.parent = nil
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.right = x
	x.parent = y

	x.height = 1 + int(math.Max(float64(getHeight(x.left)), float64(getHeight(x.right))))
	y.height = 1 + int(math.Max(float64(getHeight(y.left)), float64(getHeight(y.right))))

	return y
}
