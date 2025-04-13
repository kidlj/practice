package types

import (
	"fmt"
	"strconv"
)

func Example_Map() {
	var m1 map[string]string
	// m1 = make(map[string]string)
	fmt.Println(m1 == nil)
	_, ok := m1["hello"]
	fmt.Println(ok)

	// accessing a nil map won't panic
	hello := m1["hello"]
	fmt.Println(hello)
	i, err := strconv.Atoi(hello)
	if err != nil {
		fmt.Println(i)
		fmt.Println(err)
	}

	// panic: assignment to entry in nil map
	// m1["hello"] = "world"
	// fmt.Println(hello)

	// Output:
	// true
	// false
	//
	// 0
	// strconv.Atoi: parsing "": invalid syntax
}

func Example_Map_init() {
	var m1 = map[string]string{}
	m1["a"] = "aa"
	fmt.Println(m1)
	// Output:
	// map[a:aa]
}

func Example_Map_access() {
	roles := make(map[string][]string)
	role := roles["test"]
	fmt.Println(role)

	roles["hello"] = append(roles["hello"], "world")
	fmt.Println(roles)
	// Output:
	// []
	// map[hello:[world]]
}
