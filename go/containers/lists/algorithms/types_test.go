package algorithms

import (
	"reflect"
	"testing"
)

func Test_helpers(t *testing.T) {
	testCases := [][]int{
		{0, 1, 2, 3, 4},
		{1},
		{},
	}
	for _, tc := range testCases {
		l := fromSlice(tc)
		result := l.toSlice()
		if !reflect.DeepEqual(tc, result) {
			t.Errorf("expected %v, got %v", tc, result)
		}
	}
}
