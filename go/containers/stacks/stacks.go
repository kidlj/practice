package stacks

import "github.com/kidlj/playground/go/containers"

type Stack interface {
	containers.Container
	Push(e any)
	Pop() (any, error)
	Top() (any, error)
}
