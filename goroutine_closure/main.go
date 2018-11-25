package main

import (
	"fmt"
	"net/http"
	"time"
)

func mirroredQueryClosure() string {
	responses := make(chan string, 3)
	go func() {
		responses <- request("http://kidlj.com")
		fmt.Println("kidlj.com done")
	}()
	go func() {
		responses <- request("http://www.ideavirgin.com")
		fmt.Println("ideavirgin.com done")
	}()
	go func() {
		responses <- request("https://douban.com")
		fmt.Println("douban.com done")
	}()

	return <-responses
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	go query("http://kidlj.com", responses)
	go query("http://www.ideavirgin.com", responses)
	go query("http://douban.com", responses)

	return <-responses
}

func query(hostname string, responses chan string) {
	responses <- request(hostname)
	fmt.Println(hostname, "done")
}

func request(hostname string) string {
	_, err := http.Get(hostname)
	if err != nil {
		return fmt.Sprint(err)
	}
	return hostname
}

func main() {
	first := mirroredQuery()
	fmt.Println("first:", first)
	time.Sleep(time.Minute * 10)
}
