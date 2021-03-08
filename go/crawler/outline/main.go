package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		Outline(url)
	}
}

func Outline(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	forEachNode(doc, startElement, endElement)
}

func forEachNode(node *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(node)
	}

	for n := node.FirstChild; n != nil; n = n.NextSibling {
		fmt.Printf("%#v\n", n)
		forEachNode(n, pre, post)
	}

	if post != nil {
		post(node)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
