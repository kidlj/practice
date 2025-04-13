package algorithms

import (
	"reflect"
	"testing"
)

func Test_inOrderTraversal(t *testing.T) {
	testCases := []struct {
		t        *TreeNode
		expected []int
	}{
		{
			t:        FromSlice([]int{0, -3, 9, -10, 5}),
			expected: []int{-10, -3, 0, 5, 9},
		},
		{
			t:        nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		r := inOrderTraversal(tc.t)
		if !reflect.DeepEqual(tc.expected, r) {
			t.Errorf("expected %v, got %v", tc.expected, r)
		}
	}
}
