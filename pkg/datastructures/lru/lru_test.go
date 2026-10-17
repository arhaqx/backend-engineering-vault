package lru

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := New[string, int](2)

	cache.Put("a", 1)
	cache.Put("b", 2)

	val, found := cache.Get("a")
	if !found || val != 1 {
		t.Fatalf("expected 1, got %v", val)
	}

	// This should evict "b" because "a" was recently accessed
	cache.Put("c", 3)

	_, found = cache.Get("b")
	if found {
		t.Errorf("expected 'b' to be evicted")
	}

	val, found = cache.Get("c")
	if !found || val != 3 {
		t.Errorf("expected 'c' to have value 3")
	}
}
