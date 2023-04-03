package algorithms

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	testCases := []struct {
		l        *listNode
		expected []int
	}{
		{
			l:        fromSlice([]int{0, 1, 2, 3, 4}),
			expected: []int{4, 3, 2, 1, 0},
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
		r := reverseList(tc.l)
		if !reflect.DeepEqual(tc.expected, r.toSlice()) {
			t.Errorf("expected %v, got %v", tc.expected, r.toSlice())
		}
	}
}
