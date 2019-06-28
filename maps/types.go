package main

import "fmt"

type Annotations map[string]string

func main() {
	anno := Annotations{}
	anno["hello"] = "hello"
	anno["hello"] = "world"
	fmt.Println(anno["hello"])
}
