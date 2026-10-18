package disjointset

// UnionFind represents a disjoint-set data structure.
type UnionFind struct {
	parent []int
	rank   []int
	Count  int
}

func New(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return &UnionFind{parent: parent, rank: rank, Count: size}
}

// Find finds the representative of element x with path compression.
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union merges sets containing x and y using rank.
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX == rootY {
		return false
	}

	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
	uf.Count--
	return true
}

func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
