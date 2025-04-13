package main

import (
	"fmt"

	"gopl.io/ch7/eval"
)

func main() {
	s := "(1 + 2) * (3 + 4) / 5"
	expr, err := eval.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", expr)
	fmt.Printf("%#v\n", expr)
}
