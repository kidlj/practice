package graphs

import (
	"fmt"

	"github.com/kidlj/practice/go/containers"
	"github.com/kidlj/practice/go/containers/lists"
)

type adjacentLists struct {
	numEdges int
	adjacent []lists.List
}

func (al *adjacentLists) edges() int {
	return al.numEdges
}

func (al *adjacentLists) vertices() int {
	return len(al.adjacent)
}

func (al *adjacentLists) addEdge(e *Edge, directed bool) error {
	from, to, weight := e.from, e.to, e.weight
	if from == to {
		return fmt.Errorf("invalid edge: from %d, to %d", from, to)
	}
	if from < 0 || al.vertices() <= int(from) || to < 0 || al.vertices() <= int(to) {
		return fmt.Errorf("invalid edge: from %d, to %d", from, to)
	}
	if al.isEdge(e) {
		return nil
	}

	al.adjacent[from].Insert(0, e)
	if !directed {
		al.adjacent[to].Insert(0, &Edge{from: to, to: from, weight: weight})
	}
	al.numEdges++
	return nil
}

func (al *adjacentLists) isEdge(e *Edge) bool {
	from, to := e.from, e.to
	if from == to {
		return false
	}
	if from < 0 || al.vertices() <= int(from) || to < 0 || al.vertices() <= int(to) {
		return false
	}

	isEdge := false
	al.adjacent[from].Apply(func(e interface{}) {
		if e.(*Edge).to == to {
			isEdge = true
		}
	})
	return isEdge
}

type adjacentListsIterator struct {
	lists    *adjacentLists
	source   vertex
	listIter containers.Iterator
}

func (al *adjacentLists) newIterator(source vertex) (Iterator, error) {
	if source < 0 || al.vertices() <= int(source) {
		return nil, fmt.Errorf("The source vertex is not in the graph")
	}
	result := new(adjacentListsIterator)
	result.lists = al
	result.source = source
	result.Reset()
	return result, nil
}

func (iter *adjacentListsIterator) Reset() {
	iter.listIter = iter.lists.adjacent[iter.source].NewIterator()
}

func (iter *adjacentListsIterator) Done() bool {
	return iter.listIter.Done()
}

func (iter *adjacentListsIterator) Next() (*Edge, bool) {
	e, ok := iter.listIter.Next()
	if !ok {
		return nil, false
	}
	return e.(*Edge), true
}
