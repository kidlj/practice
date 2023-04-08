package algorithms

import (
	"reflect"
	"testing"
)

func Test_pathSum(t *testing.T) {
	testCases := []struct {
		t        *TreeNode
		sum      int
		expected [][]int
	}{
		{
			t:   FromSlice([]int{9, 6, 5, 3, 8}),
			sum: 23,
			expected: [][]int{
				{9, 6, 5, 3},
				{9, 6, 8},
			},
		},
		{
			t:   FromSlice([]int{1}),
			sum: 1,
			expected: [][]int{
				{1},
			},
		},
		{
			t:        nil,
			sum:      1, // whatever
			expected: nil,
		},
	}

	for _, tc := range testCases {
		r := pathSum(tc.t, tc.sum)
		if !reflect.DeepEqual(tc.expected, r) {
			t.Errorf("expected %v, got %v", tc.expected, r)
		}
	}
}
