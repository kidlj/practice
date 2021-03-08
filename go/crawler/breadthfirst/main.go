package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kidlj/demo/go/crawler/links"
)

// CONCURRENCY sets the concurrency level
const CONCURRENCY = 40

func main() {
	worklist := make(chan []string)
	unseenlinks := make(chan string)
	contentsAll := [][]byte{}

	go func() {
		worklist <- os.Args[1:]
	}()

	for i := 0; i < CONCURRENCY; i++ {
		go func() {
			for link := range unseenlinks {
				foundLinks, contents := crawl(link)
				contentsAll = append(contentsAll, contents)
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenlinks <- link
			}
		}
	}
}

func crawl(url string) ([]string, []byte) {
	fmt.Println(url)
	list, contents, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list, contents
}
