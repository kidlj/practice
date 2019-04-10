package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "a-nonexistent-file"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("file does not exist")
	}
}
