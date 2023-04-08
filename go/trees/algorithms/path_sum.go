package algorithms

func pathSum(root *TreeNode, sum int) (result [][]int) {
	var dfs func(root *TreeNode, path []int, sum int)
	dfs = func(root *TreeNode, path []int, sum int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if root.Val == sum && root.Left == nil && root.Right == nil {
			result = append(result, path)
		}

		sum -= root.Val
		dfs(root.Left, path, sum)
		dfs(root.Right, path, sum)
	}

	dfs(root, nil, sum)

	return
}
