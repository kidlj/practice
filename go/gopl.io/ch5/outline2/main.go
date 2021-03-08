package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	outline()
}

func outline() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal("parse doc: ", err)
	}

	var test = "sss"
	var f func(n *html.Node, prev, next func(*html.Node))
	f = func(n *html.Node, prev, next func(*html.Node)) {
		log.Println(test)
		if prev != nil {
			prev(n)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, prev, next)
		}

		if next != nil {
			next(n)
		}
	}
	f(doc, startElement, endElement)
}

func forEachNode(n *html.Node, prev func(*html.Node), next func(*html.Node)) {
	if prev != nil {
		prev(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, prev, next)
	}

	if next != nil {
		next(n)
	}
}

var indent = 0

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", indent*2, " ", n.Data)
	}
	indent++
}

func endElement(n *html.Node) {
	indent--
	if n.Type == html.ElementNode {
		fmt.Printf("%*s</%s>\n", indent*2, " ", n.Data)
	}
}
