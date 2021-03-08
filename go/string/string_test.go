package string

import (
	"fmt"
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
	var str = ".cell"
	fmt.Println(strings.Replace(str, ".", "", 1))
	// Output:
	// cell
}
