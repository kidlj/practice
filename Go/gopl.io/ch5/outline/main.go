package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(doc, nil)
}

func outline(n *html.Node, tags []string) {
	if n.Type == html.ElementNode {
		tags = append(tags, n.Data)
		fmt.Println(tags)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(c, tags)
	}
}
