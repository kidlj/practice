package string

import (
	"fmt"
	"strings"
)

func Example_contains_empty() {
	var empty string
	var s = "hello"
	b := strings.Contains(s, empty)
	fmt.Println(b)
	// Output:
	// true
}
