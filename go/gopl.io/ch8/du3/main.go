package main

import (
	"fmt"
	"os"
	"sync"
)

var semaphore = make(chan struct{}, 10)

func main() {
	var sum int64 = 0
	ch := make(chan int64)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go sumSize(".", ch, wg)
	go func() {
		wg.Wait()
		close(ch)
	}()
	for size := range ch {
		sum += size
	}
	fmt.Printf("dir total size: %d\n", sum)
}

func sumSize(path string, ch chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	semaphore <- struct{}{}
	defer func() { <-semaphore }()
	file, err := os.Stat(path)
	if err != nil {
		return
	}
	if !file.IsDir() {
		ch <- file.Size()
		return
	}
	ch <- file.Size()
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("read dir error:", err)
		return
	}

	for _, f := range files {
		wg.Add(1)
		// can be optimized: control concurrency while listing entries; see du.
		go sumSize(path+"/"+f.Name(), ch, wg)
	}
}
