package lists

import "github.com/kidlj/practice/go/containers"

// List is the interface for lists in the container hierarchy.
type List interface {
	containers.Collection
	Insert(i int, e any) error
	Delete(i int) (any, error)
	Get(i int) (any, error)
	Put(i int, e any) error
	Index(e any) (int, bool)
	Slice(i, j int) (List, error)
	Equal(l List) bool
}
