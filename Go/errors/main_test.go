package main

import (
	"errors"
	"fmt"
	"os"
)

func Example_As() {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *os.PathError
		pp := &pathError
		fmt.Printf("pathError is nil: %t\n", pathError == nil)
		fmt.Printf("pp is nil: %t\n", pp == nil)
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
	// Output:
	// pathError is nil: true
	// pp is nil: false
	// Failed at path: non-existing
}

type ExampleError string

func (e ExampleError) Error() string {
	return string(e)
}

func Example_Is() {
	err := ExampleError("test error")
	err2 := fmt.Errorf("%w", err)
	if errors.Is(err, err) {
		fmt.Println("true")
	}
	if errors.Is(err2, err) {
		fmt.Println("true")
	}
	// Output:
	// true
	// true
}
