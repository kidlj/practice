package queue

import (
	"fmt"
	"log"
)

func ExampleQueue() {
	var q ArrayQueue
	var a []interface{}
	for i := 0; i < 10; i++ {
		q.Enter(i)
	}

	for i := 0; i < 5; i++ {
		v, err := q.Leave()
		if err != nil {
			log.Fatal(err)
		}

		a = append(a, v)
	}

	for i := 10; i < 20; i++ {
		q.Enter(i)
	}

	for i := 5; i < 20; i++ {
		v, err := q.Leave()
		if err != nil {
			log.Fatal(err)
		}

		a = append(a, v)
	}

	fmt.Println(a)
	// Output:
	// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
}
