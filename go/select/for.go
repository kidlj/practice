package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Second)
	for {
		select {
		case t := <-tick:
			fmt.Println(t)
			// break 'select' not 'for'
			break
		}
	}
	fmt.Println("end")
}
