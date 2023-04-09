package algorithms

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := root.Left
	right := root.Right
	root.Left = invertTree(right)
	root.Right = invertTree(left)
	return root
}
