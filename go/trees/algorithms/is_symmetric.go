package algorithms

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

func compare(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if n1 != nil && n2 != nil {
		return compare(n1.Left, n2.Right) && compare(n1.Right, n2.Left)
	}

	return false
}
