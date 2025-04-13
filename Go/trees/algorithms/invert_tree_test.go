package algorithms

import (
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	testCases := []struct {
		t        *TreeNode
		expected []int // in-order traversal result of inverted tree
	}{
		{
			t:        FromSlice([]int{0, -3, 9, -10, 5}),
			expected: []int{9, 5, 0, -3, -10},
		},
	}

	for _, tc := range testCases {
		r := invertTree(tc.t)
		s := inOrderTraversal(r)
		if !reflect.DeepEqual(tc.expected, s) {
			t.Errorf("expected %v, got %v", tc.expected, s)
		}
	}
}
