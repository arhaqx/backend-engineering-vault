package trie

import "testing"

func TestTrie(t *testing.T) {
	tr := New()
	tr.Insert("apple")
	if !tr.Search("apple") {
		t.Errorf("expected apple to be found")
	}
	if tr.Search("app") {
		t.Errorf("expected app not to be full word")
	}
	if !tr.StartsWith("app") {
		t.Errorf("expected prefix app to match")
	}
}
