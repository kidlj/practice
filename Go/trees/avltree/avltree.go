package avltree

type node struct {
	val    int
	left   *node
	right  *node
	height int
}

type AVLTree struct {
	root *node
}

func (t *AVLTree) Size() int {
	return t.size(t.root)
}

func (t *AVLTree) size(n *node) int {
	if n == nil {
		return 0
	}

	left := t.size(n.left)
	right := t.size(n.right)

	return left + right + 1
}

func (t *AVLTree) Height() int {
	if t.root == nil {
		return -1
	}
	return t.root.height
}

func (t *AVLTree) Insert(v int) {
	t.root = t.insert(t.root, v)
}

func (t *AVLTree) insert(n *node, x int) *node {
	if n == nil {
		n = &node{val: x, height: 0}
		return n
	}

	if x > n.val {
		n.right = t.insert(n.right, x)
		if height(n.right)-height(n.left) == 2 {
			if x > n.right.val {
				n = singleRotateWithRight(n)
			} else {
				n = doubleRotateWithRight(n)
			}
		}
	} else if x < n.val {
		n.left = t.insert(n.left, x)
		if height(n.left)-height(n.right) == 2 {
			if x < n.left.val {
				n = singleRotateWithLeft(n)
			} else {
				n = doubleRotateWithLeft(n)
			}
		}
	}

	n.height = max(height(n.left), height(n.right)) + 1
	return n
}

func singleRotateWithRight(n *node) *node {
	root := n.right
	n.right = root.left
	root.left = n

	n.height = max(height(n.left), height(n.right)) + 1
	root.height = max(height(root.right), n.height) + 1

	return root
}

func singleRotateWithLeft(n *node) *node {
	root := n.left
	n.left = root.right
	root.right = n

	n.height = max(height(n.left), height(n.right)) + 1
	root.height = max(height(root.left), n.height) + 1

	return root
}

func doubleRotateWithRight(n *node) *node {
	n.right = singleRotateWithLeft(n.right)
	n = singleRotateWithRight(n)
	return n
}

func doubleRotateWithLeft(n *node) *node {
	n.left = singleRotateWithRight(n.left)
	n = singleRotateWithLeft(n)
	return n
}

func height(n *node) int {
	if n == nil {
		return -1
	}
	return n.height
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func (t *AVLTree) VisitDepthOrder(f func(v int)) {
	queue := []*node{t.root}
	for len(queue) > 0 {
		var top *node
		queue, top = queue[1:], queue[0]
		if top.left != nil {
			queue = append(queue, top.left)
		}
		if top.right != nil {
			queue = append(queue, top.right)
		}
		f(top.val)
	}
}
