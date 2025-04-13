package algorithms

func flatten(root *TreeNode) {
	var last *TreeNode

	var f func(*TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}

		f(root.Right) // Right first
		f(root.Left)

		root.Left = nil
		root.Right = last
		last = root
	}

	f(root)
}
