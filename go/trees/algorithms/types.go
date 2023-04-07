package algorithms

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// FromSlice inserts values from a slice into a binary search tree in order.
func FromSlice(s []int) *TreeNode {
	var root *TreeNode
	for _, v := range s {
		root = root.Insert(v)
	}
	return root
}

// Binary search tree insertion
func (root *TreeNode) Insert(v int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: v}
	}

	if v > root.Val {
		root.Right = root.Right.Insert(v)
	} else if v < root.Val {
		root.Left = root.Left.Insert(v)
	}

	return root
}

func (t1 *TreeNode) Equal(t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 != nil && t2 != nil {
		return t1.Val == t2.Val && t1.Left.Equal(t2.Left) && t1.Right.Equal(t2.Right)
	}

	return false
}

// transfer right-most node values to a slice, for testing purpose.
func (root *TreeNode) RightMost() (s []int) {
	for root != nil {
		s = append(s, root.Val)
		root = root.Right
	}

	return
}
