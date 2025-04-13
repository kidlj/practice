package types

import (
	"fmt"
)

func Example_Int() {
	var uweight uint
	uweight = 0
	fmt.Println(uweight)

	var weight int
	weight = -1
	fmt.Println(weight)
	// Output:
	// 0
	// -1
}

func ExampleIntDivision() {
	fmt.Println(-1 / 1)
	fmt.Println(-1 / 2)
	fmt.Println(-1 / -2)
	fmt.Println(-2 / 4)
	fmt.Println(-10 / 2)
	// Output:
	// -1
	// 0
	// 0
	// 0
	// -5
}
