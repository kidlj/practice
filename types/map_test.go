package types

import "fmt"

func Example_Map() {
	var m1 map[string]string
	m1 = make(map[string]string)
	fmt.Println(m1)
	// Output:
	// map[]
}

func Example_Map_init() {
	var m1 = map[string]string{}
	m1["a"] = "aa"
	fmt.Println(m1)
	// Output:
	// map[a:aa]
}
