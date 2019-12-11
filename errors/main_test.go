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
