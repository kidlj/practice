package stacks

import "fmt"

func ExampleIsBalancedStack() {
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
		fmt.Println(IsBalancedStack(s))
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
