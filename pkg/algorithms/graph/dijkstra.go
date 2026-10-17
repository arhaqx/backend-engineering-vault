package graph

import (
	"container/heap"
	"math"
)

type Edge struct {
	To     int
	Weight int
}

type Graph struct {
	Nodes int
	Edges map[int][]Edge
}

func NewGraph(nodes int) *Graph {
	return &Graph{
		Nodes: nodes,
		Edges: make(map[int][]Edge),
	}
}

func (g *Graph) AddEdge(u, v, weight int) {
	g.Edges[u] = append(g.Edges[u], Edge{To: v, Weight: weight})
}

type Item struct {
	node     int
	distance int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].distance < pq[j].distance }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i]; pq[i].index = i; pq[j].index = j }
func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// ShortestPath computes minimum distances from start node to all nodes using Dijkstra's algorithm.
func (g *Graph) ShortestPath(start int) map[int]int {
	dist := make(map[int]int)
	for i := 0; i < g.Nodes; i++ {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{node: start, distance: 0})

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*Item)
		u := curr.node

		if curr.distance > dist[u] {
			continue
		}

		for _, edge := range g.Edges[u] {
			v := edge.To
			w := edge.Weight

			if dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				heap.Push(&pq, &Item{node: v, distance: dist[v]})
			}
		}
	}

	return dist
}
