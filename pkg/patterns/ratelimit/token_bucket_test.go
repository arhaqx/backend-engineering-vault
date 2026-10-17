package ratelimit

import (
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	// 5 tokens capacity, 2 tokens per second refill
	tb := NewTokenBucket(5, 2)

	// Consume all 5 tokens
	for i := 0; i < 5; i++ {
		if !tb.Allow() {
			t.Fatalf("expected token %d to be allowed", i)
		}
	}

	// 6th immediate request should be denied
	if tb.Allow() {
		t.Fatalf("expected request to be rate limited")
	}

	// Wait 600ms (refills ~1.2 tokens)
	time.Sleep(600 * time.Millisecond)

	if !tb.Allow() {
		t.Fatalf("expected request to be allowed after token refill")
	}
}
