package algorithms

import (
	"reflect"
	"testing"
)

func Test_printLots(t *testing.T) {
	testCases := []struct {
		l        *listNode
		s        *listNode
		expected []int
	}{
		{
			l:        fromSlice([]int{0, 1, 2, 3, 4}),
			s:        fromSlice([]int{2, 3, 5}),
			expected: []int{1, 2, 4},
		},
		{
			l:        fromSlice([]int{0, 1, 2, 3, 4}),
			s:        fromSlice([]int{}),
			expected: []int{},
		},
		{
			l:        fromSlice([]int{}),
			s:        fromSlice([]int{}),
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		result := printLots(tc.l, tc.s)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("expected %v, got %v", tc.expected, result)
		}
	}
}
