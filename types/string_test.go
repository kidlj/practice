package types

import (
	"fmt"
	"strconv"
)

func Example_string() {
	port := 9014
	s := strconv.Itoa(port)
	fmt.Println(s)
	fmt.Println(string(port))
	// Output:
	// 9014
	// 9014
}
