package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gopl.io/ch8/thumbnail"
)

func main() {
	ch := make(chan string)
	filenames := []string{
		"main.go",
		"IV.JPG",
	}

	go func() {
		for _, filename := range filenames {
			ch <- filename
		}
		close(ch)
	}()

	size := makeThumbnails(ch)
	fmt.Println(size)
}

func makeThumbnails(filenames chan string) int64 {
	var wg sync.WaitGroup
	sizes := make(chan int64)
	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}
