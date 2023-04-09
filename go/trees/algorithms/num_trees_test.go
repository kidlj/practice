package algorithms

import (
	"reflect"
	"testing"
)

func Test_numTress(t *testing.T) {
	testCases := []int{0, 1, 3}
	expected := []int{0, 1, 5}

	var result []int
	for _, n := range testCases {
		result = append(result, numTrees(n))
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
