package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

const base = "http://test.domain.com"

func main() {
	urls := []string{
		"/header",
		"/hello",
		"/weight",
		"/400",
		"/404",
		"/500",
	}
	for {
		i := rand.Intn(6)
		url := urls[i]

		resp, err := http.Get(base + url)
		if err != nil {
			log.Println(err)
		}
		log.Printf("getting %s: %s", url, resp.Status)
		resp.Body.Close()
		time.Sleep(500 * time.Millisecond)
	}
}
