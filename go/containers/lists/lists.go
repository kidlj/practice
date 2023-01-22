package lists

import "github.com/kidlj/playground/go/containers"

// List is the interface for lists in the container hierarchy.
type List interface {
	containers.Collection
	Insert(i int, e interface{}) error
	Delete(i int) (interface{}, error)
	Get(i int) (interface{}, error)
	Put(i int, e interface{}) error
	Index(e interface{}) (int, bool)
	Slice(i, j int) (List, error)
	Equal(l List) bool
}
