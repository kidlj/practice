package algorithms

import "testing"

func Test_isSymmetric(t *testing.T) {
	testCases := []struct {
		t        *TreeNode
		expected bool
	}{
		{
			t:        buildSymmetricTree(),
			expected: true,
		},
		{
			t:        FromSlice([]int{0, -3, 9, -10, 5}),
			expected: false,
		},
	}

	for _, tc := range testCases {
		r := isSymmetric(tc.t)
		if r != tc.expected {
			t.Errorf("test failed")
		}
	}
}

func buildSymmetricTree() *TreeNode {
	t1 := FromSlice([]int{0, -3, 9, -10, 5})
	t2 := FromSlice([]int{0, -3, 9, -10, 5})
	t2 = invertTree(t2)
	return &TreeNode{Left: t1, Right: t2}
}
