package algorithms

import (
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	testCases := []struct {
		l        *listNode
		left     int
		right    int
		expected []int
	}{
		{
			l:        fromSlice([]int{1, 2, 3, 4, 5}),
			left:     2,
			right:    4,
			expected: []int{1, 4, 3, 2, 5},
		},
		{
			l:        fromSlice([]int{1}),
			left:     1,
			right:    1,
			expected: []int{1},
		},
	}

	for _, tc := range testCases {
		r := reverseBetween(tc.l, tc.left, tc.right)
		if !reflect.DeepEqual(tc.expected, r.toSlice()) {
			t.Errorf("expected %v, got %v", tc.expected, r.toSlice())
		}
	}
}
