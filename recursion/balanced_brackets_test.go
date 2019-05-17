package recursion

import "fmt"

func ExampleIsBalancedRecursive() {
	a := []string{
		"",
		"[]",
		"[][]",
		"[[]]",
		"[[]][]",
		"[[[[][]]][]]",
		"[",
		"]",
		"[[",
		"][",
		"]]",
		"[]]",
		"[a]",
	}

	for _, s := range a {
		fmt.Println(IsBalancedRecursive(s))
	}
	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// false
	// false
	// false
	// false
	// false
	// false
	// false
}
