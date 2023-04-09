package algorithms

func numTrees(n int) int {
	if n == 0 {
		return 0
	}

	nums := make([]int, n+1)
	nums[0], nums[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			nums[i] += nums[j-1] * nums[i-j]
		}
	}

	return nums[n]
}
