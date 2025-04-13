package types

import (
	"fmt"
	"strconv"
)

func Example_Itoa() {
	port := 9014
	s := strconv.Itoa(port)
	fmt.Println(s)
	fmt.Println(port)
	// Output:
	// 9014
	// 9014
}

func Example_Atoi() {
	port := "hello"
	i, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(i)
	// Output:
	// error
	// 0
}

func Example_len() {
	s := "hello"
	fmt.Println(len(s))
	// Output:
	// 5
}
