package algorithms

import (
	"reflect"
	"testing"
)

func Test_mergeLists(t *testing.T) {
	testCases := []struct {
		a *listNode
		b *listNode
		c []int
	}{
		{
			a: fromSlice([]int{1, 2, 3}),
			b: fromSlice([]int{4, 5, 6}),
			c: []int{1, 2, 3, 4, 5, 6},
		},
		{
			a: fromSlice([]int{1, 3, 5}),
			b: fromSlice([]int{2, 4, 6}),
			c: []int{1, 2, 3, 4, 5, 6},
		},
		{
			a: fromSlice([]int{1, 2, 3}),
			b: nil,
			c: []int{1, 2, 3},
		},
		{
			a: nil,
			b: fromSlice([]int{1, 2, 3}),
			c: []int{1, 2, 3},
		},
		{
			a: nil,
			b: nil,
			c: []int{},
		},
	}

	for _, tc := range testCases {
		c := mergeLists(tc.a, tc.b)
		if !reflect.DeepEqual(tc.c, c.toSlice()) {
			t.Errorf("expected %v, got %v", tc.c, c.toSlice())
		}
	}
}
