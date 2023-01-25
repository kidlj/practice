package strings

import (
	"fmt"
	"strconv"
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

func Example_Split() {
	var str = "div.comment"
	s := strings.Split(str, " ")
	fmt.Println(s)
	s = strings.Split(str, "或")
	fmt.Println(s)
	// Output:
	// [div.comment]
	// [div.comment]
}

func Example_Replace() {
	var str = "Avoiding Memory Leak in\n\t\t\t\t\tGolang API"
	str = strings.ReplaceAll(str, "\n", " ")
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\r", "")
	fmt.Println(str)
	// Output:
	// Avoiding Memory Leak in Golang API
}

func Example_Atoi() {
	s := " 345"
	i, err := strconv.Atoi(s)
	fmt.Println(i)
	fmt.Println(err)
	// Output:
	// 345
}
