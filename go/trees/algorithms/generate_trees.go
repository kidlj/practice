package algorithms

func generateTrees(n int) []*TreeNode {
	if n <= 0 {
		return nil
	}

	return generate(1, n)
}

func generate(left, right int) (result []*TreeNode) {
	if left > right {
		result = append(result, nil)
		return
	}

	for i := left; i <= right; i++ {
		leftNodes := generate(left, i-1)
		rightNodes := generate(i+1, right)
		for _, left := range leftNodes {
			for _, right := range rightNodes {
				n := &TreeNode{Val: i}
				n.Left = left
				n.Right = right
				result = append(result, n)
			}
		}
	}

	return
}
