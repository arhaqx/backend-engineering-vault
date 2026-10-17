package ratelimit

import (
	"sync"
	"time"
)

// TokenBucket limits operations based on a replenish rate and capacity.
type TokenBucket struct {
	mu         sync.Mutex
	capacity   float64
	tokens     float64
	refillRate float64 // tokens per second
	lastRefill time.Time
}

// NewTokenBucket creates a token bucket with max capacity and refill rate (per second).
func NewTokenBucket(capacity, refillRate float64) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

// Allow checks if a single token can be consumed.
func (tb *TokenBucket) Allow() bool {
	return tb.AllowN(1.0)
}

// AllowN checks if n tokens can be consumed.
func (tb *TokenBucket) AllowN(n float64) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tb.lastRefill = now

	// Refill tokens according to elapsed duration
	tb.tokens += elapsed * tb.refillRate
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	if tb.tokens >= n {
		tb.tokens -= n
		return true
	}

	return false
}
