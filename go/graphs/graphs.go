package graphs

import (
	"fmt"
	"math"

	"github.com/kidlj/practice/go/containers/lists"
	"github.com/kidlj/practice/go/heaps"
	"github.com/kidlj/practice/go/queues"
)

const (
	NotAVertex = -1
)

// Graph is the interface for undirected graphs.
type Graph interface {
	// basic graph builders and accessors
	Edges() int                             // return the number of edges in the graph
	Vertices() int                          // return the number of vertices in the graph
	AddEdge(edge *Edge) error               // add an edge between vertices v and w
	IsEdge(edge *Edge) bool                 // true iff there is an edge between v and w
	NewIterator(v vertex) (Iterator, error) // make an iterator over edges adjacent to v

	// derived graph algorithms
	DFS(v vertex, f func(Graph, *Edge)) // resursive depth-first search from vertex v
	// StackDFS(v vertex, f func(Graph, *Edge))    // stack-based depth-first search from vertex v
	BFS(v vertex, f func(Graph, *Edge))         // breadth-first search from vertex v
	IsPath(v, w vertex) bool                    // true ff there is a path between v and w
	ShortestPath(v, w vertex) ([]vertex, error) // the vertices on the shortest path from v to w
	IsConnected() bool                          // true if the graph is connected
	// SpanningTree() (Graph, error)               // construct a spanning tree for this graph
	String() string
}

type DAG interface {
	Graph
	TopologicalSort() ([]vertex, error)
}

type Iterator interface {
	Next() (*Edge, bool)
	Done() bool
	Reset()
}

type Edge struct {
	from   vertex
	to     vertex
	weight int
}

type vertex int

// graphRepresentation is the interface for the methods essential for an internal
// representations of a graph: this is the method set for graph representations.
type graphRepresentation interface {
	edges() int                                  // return the number of edges in the graph
	vertices() int                               // return the number of vertices in the graph
	addEdge(edge *Edge, directed bool) error     // add an edge between vertices v and w
	isEdge(edge *Edge) bool                      // true if there is an edge between v and w
	newIterator(source vertex) (Iterator, error) // make an iterator over edges adjacent to v
}

type unweightedGraph struct {
	representation graphRepresentation
}

type weightedGraph struct {
	unweightedGraph
}

type weightedNegativeGraph struct {
	unweightedGraph
}

type dag struct {
	unweightedGraph
}

// Implementations

func newAdjacentListsRepresentation(n int) graphRepresentation {
	rep := new(adjacentLists)
	if n < 0 {
		n = 1
	}
	rep.adjacent = make([]lists.List, n)
	for i := 0; i < n; i++ {
		rep.adjacent[i] = lists.NewLinkedList()
	}
	return rep
}

func NewGraph(n int) Graph {
	result := new(unweightedGraph)
	result.representation = newAdjacentListsRepresentation(n)
	return result
}

func (g *unweightedGraph) NewIterator(source vertex) (Iterator, error) {
	return g.representation.newIterator(source)
}

func (g *unweightedGraph) Vertices() int {
	return g.representation.vertices()
}

func (g *unweightedGraph) Edges() int {
	return g.representation.edges()
}

func (g *unweightedGraph) IsEdge(e *Edge) bool {
	return g.representation.isEdge(e)
}

func (g *unweightedGraph) AddEdge(edge *Edge) error {
	return g.representation.addEdge(edge, false)
}

func (g *unweightedGraph) DFS(from vertex, visit func(g Graph, e *Edge)) {
	visited := make([]bool, g.Vertices())
	var dfs func(g Graph, e *Edge)
	dfs = func(g Graph, e *Edge) {
		if visited[e.to] {
			return
		}
		visited[e.to] = true
		visit(g, e)
		iter, _ := g.NewIterator(e.to)
		for e, ok := iter.Next(); ok; e, ok = iter.Next() {
			dfs(g, e)
		}
	}

	se := &Edge{from: vertex(NotAVertex), to: from}
	dfs(g, se)
}

// O(|V|+|E|)
// Level order traversal
func (g *unweightedGraph) BFS(from vertex, visit func(g Graph, e *Edge)) {
	q := queues.NewLinkedQueue()
	q.Enter(&Edge{from: vertex(NotAVertex), to: from})
	visited := make([]bool, g.Vertices())
	for !q.IsEmpty() {
		e, _ := q.Leave()
		// Each vertex is processed at most once.
		if visited[e.(*Edge).to] {
			continue
		}
		visited[e.(*Edge).to] = true
		visit(g, e.(*Edge))
		// Each edge is processed at most once.
		iter, _ := g.NewIterator(e.(*Edge).to)
		for e, ok := iter.Next(); ok; e, ok = iter.Next() {
			q.Enter(e)
		}
	}
}

func (g *unweightedGraph) IsPath(from, to vertex) bool {
	if from < 0 || g.Vertices() <= int(from) || to < 0 || g.Vertices() <= int(to) {
		return false
	}
	result := false
	visit := func(g Graph, e *Edge) {
		if e.to == to {
			result = true
		}
	}
	g.DFS(from, visit)
	return result
}

// O(|V|+|E|)
func (g *unweightedGraph) ShortestPath(from, to vertex) ([]vertex, error) {
	if !g.IsPath(from, to) {
		return nil, fmt.Errorf("The vertices are not connected")
	}
	toEdge := make([]vertex, g.Vertices())
	visit := func(g Graph, e *Edge) {
		toEdge[e.to] = e.from
	}
	g.BFS(from, visit)
	result := make([]vertex, 0, g.Vertices())
	for to != NotAVertex {
		result = append([]vertex{to}, result...)
		to = toEdge[to]
	}
	return result, nil
}

func (g *unweightedGraph) IsConnected() bool {
	visited := make([]bool, g.Vertices())
	visit := func(g Graph, e *Edge) {
		visited[e.to] = true
	}
	g.DFS(0, visit)
	result := true
	for _, v := range visited {
		if !v {
			result = false
			break
		}
	}
	return result
}

// String produces a string representation of a graph.
func (g *unweightedGraph) String() string {
	result := ""
	for i := 0; i < g.Vertices(); i++ {
		result += fmt.Sprintf("%d:", i)
		iter, _ := g.NewIterator(vertex(i))
		for v, ok := iter.Next(); ok; v, ok = iter.Next() {
			result += fmt.Sprintf(" %d", v)
		}
		result += "\n"
	}
	return result
}

// DAG

func NewDAG(n int) DAG {
	result := new(dag)
	result.representation = newAdjacentListsRepresentation(n)
	return result
}

func (g *dag) AddEdge(edge *Edge) error {
	return g.representation.addEdge(edge, true)
}

// DAG only
func (g *dag) TopologicalSort() ([]vertex, error) {
	result := make([]vertex, 0, g.Vertices())
	q := queues.NewLinkedQueue()
	indgrees := g.indgrees()
	for i, d := range indgrees {
		if d == 0 {
			q.Enter(vertex(i))
		}
	}
	for !q.IsEmpty() {
		v, _ := q.Leave()
		result = append(result, v.(vertex))

		iter, _ := g.NewIterator(v.(vertex))
		for e, ok := iter.Next(); ok; e, ok = iter.Next() {
			indgrees[e.to]--
			if indgrees[e.to] == 0 {
				q.Enter(vertex(e.to))
			}
		}
	}

	if len(result) != g.Vertices() {
		return nil, fmt.Errorf("Graph has a cycle!")
	}

	return result, nil
}

func (g *dag) indgrees() []int {
	result := make([]int, g.Vertices())
	for i := 0; i < g.Vertices(); i++ {
		iter, _ := g.NewIterator(vertex(i))
		for e, ok := iter.Next(); ok; e, ok = iter.Next() {
			result[e.to]++
		}
	}
	return result
}

// Weighted Graph

func NewWeightedGraph(n int) Graph {
	result := new(weightedGraph)
	result.representation = newAdjacentListsRepresentation(n)
	return result
}

// Dijkstra algorithm
func (g *weightedGraph) ShortestPath(from, to vertex) ([]vertex, error) {
	if !g.IsPath(from, to) {
		return nil, fmt.Errorf("The vertices are not connected")
	}

	pq := heaps.NewGenericBinaryHeap(func(a, b distanceInfo) int { return a.dis - b.dis })
	pq.Insert(distanceInfo{v: from, dis: 0})

	t := initVertexTable(g.Vertices(), from)

	for !pq.IsEmpty() {
		i, _ := pq.DeleteMin()
		v := i.v
		if t[v].known {
			continue
		}
		t[v].known = true
		iter, _ := g.NewIterator(v)
		for e, ok := iter.Next(); ok; e, ok = iter.Next() {
			w, dis := e.to, e.weight
			if !t[w].known {
				newDist := t[v].dis + dis
				if newDist < t[w].dis {
					t[w].dis = newDist
					t[w].path = v
					pq.Insert(distanceInfo{v: w, dis: newDist})
				}
			}
		}
	}

	result := make([]vertex, 0, g.Vertices())
	for to != NotAVertex {
		result = append([]vertex{to}, result...)
		to = t[to].path
	}
	return result, nil
}

type distanceInfo struct {
	v   vertex
	dis int
}

type vertexInfo struct {
	dis   int
	known bool
	path  vertex
}

func initVertexTable(n int, i vertex) []vertexInfo {
	t := make([]vertexInfo, n)
	for i := 0; i < n; i++ {
		t[i].dis = math.MaxInt
		t[i].known = false
	}
	t[i].dis = 0
	t[i].path = NotAVertex
	return t
}

// Weighted Negative Graph

func NewWeightedNegativeGraph(n int) Graph {
	result := new(weightedNegativeGraph)
	result.representation = newAdjacentListsRepresentation(n)
	return result
}
