package nil

import "fmt"

func Example_nil_error() {
	var err error
	// Pointer of nil is legal and not nil
	fmt.Println(&err)

	var m map[int]string
	fmt.Println(&m)
	// Output:
	// good
}
