package types

import "fmt"

import "reflect"

func Example_Pointer_Equal() {
	var a *int
	var b *int
	fmt.Println(a == b)
	// fmt.Println(*a == *b) // panic
	fmt.Println(reflect.DeepEqual(a, b))

	s := 3
	q := 3
	a = &s
	b = &q
	fmt.Println(reflect.DeepEqual(a, b))
	fmt.Println(a == b)
	// Output:
	// true
	// true
	// true
	// false
}
