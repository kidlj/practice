package types

import (
	"fmt"
	"reflect"
)

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

func Example_nil() {
	var a = new(int)
	fmt.Println(*a)
	fmt.Println(a == nil)
	var b *int
	// panic !!!
	*b = 3
	// Output:
	// 0
	// false
}
