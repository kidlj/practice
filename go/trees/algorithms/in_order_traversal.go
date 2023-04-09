package algorithms

func inOrderTraversal(root *TreeNode) (result []int) {
	if root == nil {
		return
	}

	result = append(result, inOrderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inOrderTraversal(root.Right)...)

	return
}
