package graph

import (
	"testing"
)

func TestDijkstra(t *testing.T) {
	g := NewGraph(5)
	g.AddEdge(0, 1, 4)
	g.AddEdge(0, 2, 2)
	g.AddEdge(2, 1, 1)
	g.AddEdge(1, 3, 5)
	g.AddEdge(2, 3, 8)
	g.AddEdge(2, 4, 10)
	g.AddEdge(3, 4, 2)

	distances := g.ShortestPath(0)

	expected := map[int]int{
		0: 0,
		1: 3, // 0 -> 2 (2) -> 1 (1) = 3
		2: 2,
		3: 8, // 0 -> 2 -> 1 -> 3 = 8
		4: 10, // 0 -> 2 -> 1 -> 3 -> 4 = 10
	}

	for node, expDist := range expected {
		if distances[node] != expDist {
			t.Errorf("node %d: expected distance %d, got %d", node, expDist, distances[node])
		}
	}
}
