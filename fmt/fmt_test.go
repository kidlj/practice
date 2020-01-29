package fmt

import "fmt"

func Example_quote() {
	s := "(/|$)(.*)"
	fmt.Println(quote(s))
	// Output:
	// "(/|$)(.*)"
}
