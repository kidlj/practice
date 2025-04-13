package main

import (
	"fmt"

	"github.com/labstack/gommon/random"
)

func main() {
	code := random.New().String(6, random.Numeric)
	fmt.Println(code)
}
