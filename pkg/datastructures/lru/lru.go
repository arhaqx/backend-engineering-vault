package lru

import (
	"container/list"
	"sync"
)

type entry[K comparable, V any] struct {
	key   K
	value V
}

// Cache is a thread-safe LRU cache using generics.
type Cache[K comparable, V any] struct {
	mu       sync.RWMutex
	capacity int
	items    map[K]*list.Element
	evictList *list.List
}

// New creates a new Cache with the given capacity.
func New[K comparable, V any](capacity int) *Cache[K, V] {
	return &Cache[K, V]{
		capacity:  capacity,
		items:     make(map[K]*list.Element),
		evictList: list.New(),
	}
}

// Get looks up a key's value from the cache.
func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		c.evictList.MoveToFront(elem)
		return elem.Value.(*entry[K, V]).value, true
	}

	var zero V
	return zero, false
}

// Put adds or updates a key-value pair in the cache.
func (c *Cache[K, V]) Put(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		c.evictList.MoveToFront(elem)
		elem.Value.(*entry[K, V]).value = value
		return
	}

	if c.evictList.Len() >= c.capacity {
		c.evictOldest()
	}

	ent := &entry[K, V]{key: key, value: value}
	elem := c.evictList.PushFront(ent)
	c.items[key] = elem
}

func (c *Cache[K, V]) evictOldest() {
	elem := c.evictList.Back()
	if elem != nil {
		c.evictList.Remove(elem)
		kv := elem.Value.(*entry[K, V])
		delete(c.items, kv.key)
	}
}

// Len returns the current number of items.
func (c *Cache[K, V]) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}
