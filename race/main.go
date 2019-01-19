package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Foo struct {
	Content string `json:"content"`
}

type FooSlice []*Foo

func updateFooSlice(fooSlice FooSlice) {
	for {
		foo := &Foo{Content: "new"}
		(fooSlice)[0] = foo
		fmt.Printf("(fooSlice[0]).Content: %s", (fooSlice[0]).Content)
		time.Sleep(time.Second)
	}
}

func installHttpHandler(fooSlice FooSlice) {
	router := gin.New()
	handler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"items": fooSlice,
		})
	}
	router.GET("/", handler)

	router.Run(":8080")
}

func main() {
	foo1 := &Foo{Content: "hey"}
	foo2 := &Foo{Content: "yo"}
	fooSlice := FooSlice{foo1, foo2}

	go updateFooSlice(fooSlice)

	installHttpHandler(fooSlice)

}
