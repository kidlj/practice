package algorithms

import "testing"

func Test_hasPathSum(t *testing.T) {
	testCases := []struct {
		t        *TreeNode
		sum      int
		expected bool
	}{
		{
			t:        FromSlice([]int{0, -3, 9, -10, 5}),
			sum:      14,
			expected: true,
		},
		{
			t:        FromSlice([]int{0, -3, 9, -10, 5}),
			sum:      9, // sum from root to non-leaf node
			expected: false,
		},
		{
			t:        FromSlice([]int{1}),
			sum:      1,
			expected: true,
		},
		{
			t:        nil,
			sum:      0,
			expected: false,
		},
	}

	for i, tc := range testCases {
		r := hasPathSum(tc.t, tc.sum)
		if r != tc.expected {
			t.Errorf("test case %d failed", i)
		}
	}
}
