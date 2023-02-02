package binarysearchtree

type node struct {
	val   int
	left  *node
	right *node
}

type BinarySearchTree struct {
	root *node
	size int
}

func (t *BinarySearchTree) Size() int {
	return t.size
}

func (t *BinarySearchTree) Find(v int) *node {
	return t.find(t.root, v)
}

func (t *BinarySearchTree) find(n *node, v int) *node {
	if n == nil {
		return nil
	}

	if v > n.val {
		return t.find(n.right, v)
	} else if v < n.val {
		return t.find2(n.left, v)
	}

	return n
}

func (t *BinarySearchTree) find2(n *node, v int) *node {
	for n != nil {
		if n.val == v {
			return n
		} else if v > n.val {
			n = n.right
		} else if v < n.val {
			n = n.left
		}
	}

	return nil
}

func (t *BinarySearchTree) Insert(v int) {
	t.root = t.insert(t.root, v)
}

func (t *BinarySearchTree) insert(n *node, v int) *node {
	if n == nil {
		n = &node{val: v}
		t.size++
		return n
	}

	if v > n.val {
		n.right = t.insert(n.right, v)
	} else if v < n.val {
		n.left = t.insert(n.left, v)
	}

	return n
}

func (t *BinarySearchTree) Delete(v int) {
	t.delete(t.root, v)
}

func (t *BinarySearchTree) delete(n *node, v int) *node {
	if n == nil {
		return n
	}

	if v > n.val {
		n.right = t.delete(n.right, v)
	} else if v < n.val {
		n.left = t.delete(n.left, v)
	} else {
		if n.left != nil && n.right != nil {
			min := t.findMin(n.right)
			n.val = min.val
			n.right = t.delete(n.right, n.val)
		} else {
			if n.left == nil {
				n = n.right
				t.size--
			} else if n.right == nil {
				n = n.left
				t.size--
			}
		}
	}

	return n
}

func (t *BinarySearchTree) FindMin() *node {
	return t.findMin(t.root)
}

func (t *BinarySearchTree) findMin(n *node) *node {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}

	return n
}

func (t *BinarySearchTree) VisitPreOrder(f func(v int)) {
	t.visitPreOrder(t.root, f)
}

func (t *BinarySearchTree) visitPreOrder(n *node, f func(v int)) {
	if n == nil {
		return
	}

	f(n.val)

	t.visitPreOrder(n.left, f)
	t.visitPreOrder(n.right, f)
}

// MARK: Iterators

type Iterator interface {
	Reset()
	Done() bool
	Next() (int, bool)
}

type PreOrderIterator struct {
	stack []*node
}

func (t *BinarySearchTree) NewPreOrderIterator() Iterator {
	iter := &PreOrderIterator{stack: []*node{t.root}}
	return iter
}

func (iter *PreOrderIterator) Next() (int, bool) {
	if len(iter.stack) == 0 {
		return 0, false
	}

	var top *node
	iter.stack, top = iter.stack[:len(iter.stack)-1], iter.stack[len(iter.stack)-1]
	if top.right != nil {
		iter.stack = append(iter.stack, top.right)
	}
	if top.left != nil {
		iter.stack = append(iter.stack, top.left)
	}
	return top.val, true
}

func (iter *PreOrderIterator) Reset() {
	iter.stack = nil
}

func (iter *PreOrderIterator) Done() bool {
	return len(iter.stack) == 0
}

type InOrderIterator struct {
	stack []*node
}

func (t *BinarySearchTree) NewInOrderIterator() Iterator {
	iter := new(InOrderIterator)
	iter.push(t.root)
	return iter
}

func (iter *InOrderIterator) Next() (int, bool) {
	if len(iter.stack) == 0 {
		return 0, false
	}

	var top *node
	iter.stack, top = iter.stack[:len(iter.stack)-1], iter.stack[len(iter.stack)-1]
	iter.push(top.right)
	return top.val, true
}

func (iter *InOrderIterator) Reset() {
	iter.stack = nil
}

func (iter *InOrderIterator) Done() bool {
	return len(iter.stack) == 0
}

func (iter *InOrderIterator) push(n *node) {
	for n != nil {
		iter.stack = append(iter.stack, n)
		n = n.left
	}
}

// TODO
type PostOrderIterator struct {
}

func (iter *PostOrderIterator) Next() (int, bool) {
	return 0, false
}

type DepthOrderIterator struct {
	queue []*node
}

func (t *BinarySearchTree) NewDepthOrderIterator() Iterator {
	iter := &DepthOrderIterator{queue: []*node{t.root}}
	return iter
}

func (iter *DepthOrderIterator) Next() (int, bool) {
	if len(iter.queue) == 0 {
		return 0, false
	}

	var top *node
	iter.queue, top = iter.queue[1:], iter.queue[0]
	if top.left != nil {
		iter.queue = append(iter.queue, top.left)
	}
	if top.right != nil {
		iter.queue = append(iter.queue, top.right)
	}
	return top.val, true
}

func (iter *DepthOrderIterator) Reset() {
	iter.queue = nil
}

func (iter *DepthOrderIterator) Done() bool {
	return len(iter.queue) == 0
}
