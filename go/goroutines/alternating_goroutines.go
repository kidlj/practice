package goroutines

import "fmt"

func alternatingGoroutinesABBA() {
	ch := make(chan struct{})

	go func() {
		for i := 0; i < 5; i++ {
			<-ch
			fmt.Print("A")
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- struct{}{}
		fmt.Print("B")
	}
	fmt.Println()
}

func alternatingGoroutinesABAB() {
	chA := make(chan struct{})
	chB := make(chan struct{})

	go func() {
		for i := 0; i < 5; i++ {
			<-chA
			fmt.Print("A")
			chB <- struct{}{}
		}
	}()

	for i := 0; i < 5; i++ {
		chA <- struct{}{}
		fmt.Print("B")
		<-chB
	}
	fmt.Println()
}

func alternatingGoroutinesABAB_Buffered() {
	chA := make(chan struct{}, 1)
	chB := make(chan struct{}, 1)
	chA <- struct{}{}

	go func() {
		for i := 0; i < 5; i++ {
			<-chA
			fmt.Print("A")
			chB <- struct{}{}
		}
	}()

	for i := 0; i < 5; i++ {
		<-chB
		fmt.Print("B")
		chA <- struct{}{}
	}
	fmt.Println()
}
