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
