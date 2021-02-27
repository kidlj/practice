package sorting

import (
	"fmt"
)

func ExampleHeapsort() {
	a := []int{3, 4, 7, 9, 2, 6, 5, 8, 1, 0}
	Heapsort(a)
	fmt.Println(a)
	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
}
