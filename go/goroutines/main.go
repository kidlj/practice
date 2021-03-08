package main

import (
	"fmt"
	"time"
)

func main() {
	chA := make(chan int)
	chB := make(chan int)

	go func() { chB <- 0 }()

	for i := 0; i < 1; i++ {
		go func() {
			for a := range chA {
				time.Sleep(time.Second * 1) // simulate time consumption
				fmt.Println("A:", a)
				go func(a int) { // goroutine accumulation
					chB <- a
				}(a)
			}
		}()
	}

	for b := range chB { // consume one
		fmt.Println("B:", b)
		for i := 0; i < 10; i++ { // produce many
			chA <- i
		}
	}
}
