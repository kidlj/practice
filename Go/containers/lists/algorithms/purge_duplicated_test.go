package algorithms

import (
	"reflect"
	"testing"
)

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
func Test_purgeDuplicated(t *testing.T) {
	testCases := []struct {
		l        *listNode
		expected []int
	}{
		{
			l:        fromSlice([]int{0, 0, 1, 2, 2, 2, 3, 3, 3}),
			expected: []int{1},
		},
		{
			l:        fromSlice([]int{1}),
			expected: []int{1},
		},
		{
			l:        nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		r := purgeDuplicated(tc.l)
		if !reflect.DeepEqual(tc.expected, r.toSlice()) {
			t.Errorf("expected %v, got %v", tc.expected, r.toSlice())
		}
	}
}
