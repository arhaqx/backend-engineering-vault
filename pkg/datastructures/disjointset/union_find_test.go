package disjointset

import "testing"

func TestUnionFind(t *testing.T) {
	uf := New(5)

	uf.Union(0, 1)
	uf.Union(1, 2)

	if !uf.Connected(0, 2) {
		t.Errorf("expected 0 and 2 to be connected")
	}

	if uf.Connected(0, 3) {
		t.Errorf("did not expect 0 and 3 to be connected")
	}

	if uf.Count != 3 {
		t.Errorf("expected 3 disjoint sets, got %d", uf.Count)
	}
}
