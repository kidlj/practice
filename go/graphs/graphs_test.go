package graphs

import (
	"math/rand"
	"testing"
	"time"
)

func TestUnweightedGraph(t *testing.T) {
	testSearch(t, "LinkedGraph", NewGraph(20))
	testUnweightedAlgorithms(t, "LinkedGraph", NewGraph(10))
}

func TestDAG(t *testing.T) {
	testSearch(t, "DAG", NewDAG(20))
	testDAGAlgorithms(t, "DAG", NewDAG(10))
}

func TestWeightedGraph(t *testing.T) {
	testSearch(t, "WeightedGraph", NewWeightedGraph(20))
	testWeightedAlgorithms(t, "WeightedGraph", NewWeightedGraph(10))
}

func testSearch(t *testing.T, name string, g Graph) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 100; i++ {
		_ = g.AddEdge(&Edge{vertex(rand.Intn(20)), vertex(rand.Intn(20)), 0})
	}
	/* Add this code to guarantee that the graph is connected. Usually it will be by chance.
	for i := 0; i < g.Vertices(); i++ {
		g.AddEdge(i,(i+12)%20)
	}
	*/
	counts := make([]int, g.Vertices())
	f := func(g Graph, e *Edge) {
		counts[e.to]++
	}
	g.DFS(0, f)
	for i := 0; i < g.Vertices(); i++ {
		if counts[i] == 0 {
			t.Errorf(name+": DFS did not visit vertex %v", i)
		} else if counts[i] > 1 {
			t.Errorf(name+": DFS visited vertex %v %v times", i, counts[i])
		}
	}

	counts = make([]int, g.Vertices())
	g.BFS(0, f)
	for i := 0; i < g.Vertices(); i++ {
		if counts[i] == 0 {
			t.Errorf(name+": BFS did not visit vertex %v", i)
		} else if counts[i] > 1 {
			t.Errorf(name+": BFS visited vertex %v %v times", i, counts[i])
		}
	}
}

func testUnweightedAlgorithms(t *testing.T, name string, g Graph) {
	g.AddEdge(&Edge{0, 1, 1})
	g.AddEdge(&Edge{0, 3, 1})
	g.AddEdge(&Edge{2, 3, 1})
	g.AddEdge(&Edge{4, 3, 1})
	g.AddEdge(&Edge{3, 5, 1})
	g.AddEdge(&Edge{4, 5, 1})
	g.AddEdge(&Edge{6, 7, 1})
	g.AddEdge(&Edge{8, 7, 1})
	g.AddEdge(&Edge{8, 9, 1})
	if g.IsPath(3, 6) || g.IsPath(6, 3) {
		t.Errorf(name + ": there is no path from 3 to 6")
	}
	if !g.IsPath(5, 1) && !g.IsPath(1, 5) {
		t.Errorf(name + ": there is a path from 5 to 1")
	}
	if g.IsConnected() {
		t.Errorf(name + ": graph is not connected, but IsConnected says it is")
	}
	g.AddEdge(&Edge{2, 8, 1})
	g.AddEdge(&Edge{3, 6, 1})
	if !g.IsConnected() {
		t.Errorf(name + ": graph is connected, but IsConnected says it is not")
	}

	// test shortest path generation
	path, _ := g.ShortestPath(0, 9)
	if !samePath(path, []vertex{0, 3, 2, 8, 9}) {
		t.Errorf(name + ": failed to find the shortes path from 0 to 9")
	}
	path, _ = g.ShortestPath(5, 7)
	if !samePath(path, []vertex{5, 3, 6, 7}) {
		t.Errorf(name + ": failed to find the shortes path from 5 to 7")
	}

	// test spanning tree generation
	// h, err := g.SpanningTree()
	// if err != nil {
	// 	t.Errorf(name + ": spanning tree generation failed")
	// }
	// if g.Vertices() != h.Vertices() || !h.IsConnected() || h.Edges() != h.Vertices()-1 {
	// 	t.Errorf(name + ": spanning tree generation failed with a bad spanning tree")
	// }
}

func samePath(p, q []vertex) bool {
	if len(p) != len(q) {
		return false
	}
	for i := 0; i < len(p); i++ {
		if p[i] != q[i] {
			return false
		}
	}
	return true
}

func testDAGAlgorithms(t *testing.T, name string, g DAG) {
	/*
			2		3

		0		4		6

			1		5
	*/
	g.AddEdge(&Edge{0, 1, 1})
	g.AddEdge(&Edge{0, 2, 1})
	g.AddEdge(&Edge{1, 4, 1})
	g.AddEdge(&Edge{1, 5, 1})
	g.AddEdge(&Edge{2, 4, 1})
	g.AddEdge(&Edge{2, 3, 1})
	g.AddEdge(&Edge{3, 4, 1})
	g.AddEdge(&Edge{3, 6, 1})
	g.AddEdge(&Edge{4, 5, 1})
	g.AddEdge(&Edge{4, 6, 1})
	g.AddEdge(&Edge{5, 6, 1})
	if !g.IsPath(0, 5) {
		t.Errorf(name + ": there is a path from 0 to 5")
	}
	if g.IsPath(6, 3) {
		t.Errorf(name + ": there is no path from 6 to 3")
	}
	if g.IsConnected() {
		t.Errorf(name + ": graph is not connected, but IsConnected says it is")
	}

	// test shortest path generation
	path, _ := g.ShortestPath(0, 5)
	expected := []vertex{0, 1, 5}
	if !samePath(path, expected) {
		t.Errorf(name+": failed to find the shortes path from 0 to 5, expected %v, got %v", expected, path)
	}
	path, _ = g.ShortestPath(2, 5)
	expected = []vertex{2, 4, 5}
	if !samePath(path, expected) {
		t.Errorf(name+": failed to find the shortes path from 2 to 5, expected %v, got %v", expected, path)
	}

	// test topological sort
	topo, _ := g.TopologicalSort()
	if !samePath(topo, []vertex{0, 7, 8, 9, 2, 1, 3, 4, 5, 6}) {
		t.Errorf(name+": failed to topo sort the graph: got %v", topo)
	}
}

func testWeightedAlgorithms(t *testing.T, name string, g Graph) {
	/*
			2		3

		0		4		6

			1		5
	*/
	g.AddEdge(&Edge{0, 1, 8})
	g.AddEdge(&Edge{0, 2, 1})
	g.AddEdge(&Edge{1, 4, 1})
	g.AddEdge(&Edge{1, 5, 1})
	g.AddEdge(&Edge{2, 3, 2})
	g.AddEdge(&Edge{2, 4, 10})
	g.AddEdge(&Edge{3, 4, 3})
	g.AddEdge(&Edge{3, 6, 1})
	g.AddEdge(&Edge{4, 5, 2})
	g.AddEdge(&Edge{4, 6, 5})
	g.AddEdge(&Edge{5, 6, 8})
	if !g.IsPath(0, 5) {
		t.Errorf(name + ": there is a path from 0 to 5")
	}
	if !g.IsPath(2, 5) {
		t.Errorf(name + ": there is a path from 2 to 5")
	}
	// 7,8,9 are not connected
	if g.IsConnected() {
		t.Errorf(name + ": graph is not connected, but IsConnected says it is")
	}

	// test shortest path generation
	path, _ := g.ShortestPath(0, 5)
	expected := []vertex{0, 2, 3, 4, 5}
	if !samePath(path, expected) {
		t.Errorf(name+": failed to find the shortes path from 0 to 5, expected %v, got %v", expected, path)
	}
	path, _ = g.ShortestPath(2, 5)
	expected = []vertex{2, 3, 4, 5}
	if !samePath(path, expected) {
		t.Errorf(name+": failed to find the shortes path from 2 to 5, expected %v, got %v", expected, path)
	}
	path, _ = g.ShortestPath(1, 6)
	expected = []vertex{1, 4, 3, 6}
	if !samePath(path, expected) {
		t.Errorf(name+": failed to find the shortes path from 1 to 6, expected %v, got %v", expected, path)
	}
}
