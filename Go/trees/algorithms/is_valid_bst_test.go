package algorithms

import "testing"

func Test_isValidBST(t *testing.T) {
	testCases := []struct {
		t        *TreeNode
		expected bool
	}{
		{
			t:        FromSlice([]int{0, -3, 9, -10, 5}),
			expected: true,
		},
		{
			t:        buildInvalidBST(),
			expected: false,
		},
	}

	for _, tc := range testCases {
		r := isValidBST(tc.t)
		if r != tc.expected {
			t.Error("test failed")
		}
	}
}

func buildInvalidBST() *TreeNode {
	t := FromSlice([]int{0, -3, 9, -10, 5})
	return &TreeNode{Val: 3, Right: t}
}
