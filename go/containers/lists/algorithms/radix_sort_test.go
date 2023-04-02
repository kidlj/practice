package algorithms

import "fmt"

func ExampleRadixSort() {
	s := []int{2, 1, 3, 4, 6, 5, 7, 0, 9, 8, 14, 13, 15, 12, 11, 10, 146}
	o := RadixSort(s)
	fmt.Println(o)
	// Output:
	// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 146]
}
