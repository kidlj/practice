package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println("Hello, playground", i)
		updateTime := fmt.Sprintf("%d", time.Now().UnixNano())
		fmt.Println(updateTime)
	}
}
