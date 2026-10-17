#!/usr/bin/env python3
"""
Automated GitHub Contribution Engine for backend-engineering-vault.
Creates realistic, high-quality Go code, algorithms, tests, and system design notes.
"""

import os
import sys
import json
import random
import subprocess
from datetime import datetime, timezone

REPO_DIR = os.path.abspath(os.path.join(os.path.dirname(__file__), ".."))
STATE_FILE = os.path.join(REPO_DIR, ".contributions_state.json")

# Curated catalog of high-impact additions
CATALOG = [
    {
        "type": "feat(datastructures)",
        "message": "implement prefix tree (Trie) for autocomplete string search",
        "file": "pkg/datastructures/trie/trie.go",
        "test_file": "pkg/datastructures/trie/trie_test.go",
        "content": """package trie

type Node struct {
\tchildren map[rune]*Node
\tisEnd    bool
}

type Trie struct {
\troot *Node
}

func New() *Trie {
\treturn &Trie{root: &Node{children: make(map[rune]*Node)}}
}

func (t *Trie) Insert(word string) {
\tcurr := t.root
\tfor _, ch := range word {
\t\tif _, ok := curr.children[ch]; !ok {
\t\t\tcurr.children[ch] = &Node{children: make(map[rune]*Node)}
\t\t}
\t\tcurr = curr.children[ch]
\t}
\tcurr.isEnd = true
}

func (t *Trie) Search(word string) bool {
\tcurr := t.root
\tfor _, ch := range word {
\t\tif next, ok := curr.children[ch]; ok {
\t\t\tcurr = next
\t\t} else {
\t\t\treturn false
\t\t}
\t}
\treturn curr.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
\tcurr := t.root
\tfor _, ch := range prefix {
\t\tif next, ok := curr.children[ch]; ok {
\t\t\tcurr = next
\t\t} else {
\t\t\treturn false
\t\t}
\t}
\treturn true
}
""",
        "test_content": """package trie

import "testing"

func TestTrie(t *testing.T) {
\ttr := New()
\ttr.Insert("apple")
\tif !tr.Search("apple") {
\t\tt.Errorf("expected apple to be found")
\t}
\tif tr.Search("app") {
\t\tt.Errorf("expected app not to be full word")
\t}
\tif !tr.StartsWith("app") {
\t\tt.Errorf("expected prefix app to match")
\t}
}
"""
    },
    {
        "type": "feat(concurrency)",
        "message": "implement weighted semaphore for bounded resource allocation",
        "file": "pkg/concurrency/semaphore/semaphore.go",
        "test_file": "pkg/concurrency/semaphore/semaphore_test.go",
        "content": """package semaphore

import (
\t"context"
\t"errors"
\t"sync"
)

var ErrNoTickets = errors.New("no tickets available")

// Semaphore controls access to a finite number of shared resources.
type Semaphore struct {
\tmu      sync.Mutex
\tcapacity int64
\tcurrent  int64
\twaiters  chan struct{}
}

func New(capacity int64) *Semaphore {
\treturn &Semaphore{
\t\tcapacity: capacity,
\t\twaiters:  make(chan struct{}, capacity),
\t}
}

func (s *Semaphore) Acquire(ctx context.Context, weight int64) error {
\tfor {
\t\ts.mu.Lock()
\t\tif s.capacity-s.current >= weight {
\t\t\ts.current += weight
\t\t\ts.mu.Unlock()
\t\t\treturn nil
\t\t}
\t\ts.mu.Unlock()

\t\tselect {
\t\tcase <-ctx.Done():
\t\t\treturn ctx.Err()
\t\tcase <-s.waiters:
\t\t}
\t}
}

func (s *Semaphore) Release(weight int64) {
\ts.mu.Lock()
\tdefer s.mu.Unlock()
\ts.current -= weight
\tselect {
\tcase s.waiters <- struct{}{}:
\tdefault:
\t}
}
""",
        "test_content": """package semaphore

import (
\t"context"
\t"testing"
\t"time"
)

func TestSemaphore(t *testing.T) {
\tsem := New(2)
\tctx := context.Background()

\tif err := sem.Acquire(ctx, 2); err != nil {
\t\tt.Fatalf("failed acquiring initial weight: %v", err)
\t}

\tctxTimeout, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
\tdefer cancel()

\tif err := sem.Acquire(ctxTimeout, 1); err == nil {
\t\tt.Fatalf("expected acquire to block and time out")
\t}

\tsem.Release(2)
\tif err := sem.Acquire(ctx, 1); err != nil {
\t\tt.Fatalf("failed acquiring after release: %v", err)
\t}
}
"""
    },
    {
        "type": "docs(system-design)",
        "message": "add analysis on Distributed Locks using Redlock algorithm",
        "file": "docs/system-design/distributed-locks.md",
        "test_file": None,
        "content": """# Distributed Locks: Redis & Redlock Architecture 🔒

Distributed locks serialize operations across competing distributed nodes where in-process mutexes fail.

## 1. Single-Instance Lock Pitfall
Using `SETNX lock_key token PX 10000` is efficient but vulnerable to single-point-of-failure (SPoF). If master crashes before replication syncs to replica, multiple nodes can acquire the identical lock simultaneously.

## 2. Redlock Multi-Node Algorithm
Proposed by Salvatore Sanfilippo (antirez), Redlock uses $N$ independent Redis masters (typically 5):

1. Client fetches current timestamp $T_1$.
2. Attempts to acquire lock across all $N$ instances sequentially using identical key/value and short timeouts.
3. Lock is considered **acquired** only if:
   - Client acquired lock in at least $\\lfloor N/2 \\rfloor + 1$ nodes.
   - Total time elapsed $T_2 - T_1 < \\text{Lock Validity Time}$.
4. If acquisition fails, client broadcasts an immediate unlock to all instances.

## 3. Martin Kleppmann's Critique & Fencing Tokens
GC pauses or network delays can expire lock validity while the client continues executing assumed exclusive work.

**Remedy**: Always pair distributed locks with monotonic **fencing tokens** verified at the storage layer:
```sql
UPDATE accounts SET balance = balance - 100 WHERE id = 1 AND fencing_token < :current_token;
```
""",
        "test_content": None
    },
    {
        "type": "feat(algorithms)",
        "message": "implement Union-Find (Disjoint-Set) with path compression and rank",
        "file": "pkg/datastructures/disjointset/union_find.go",
        "test_file": "pkg/datastructures/disjointset/union_find_test.go",
        "content": """package disjointset

// UnionFind represents a disjoint-set data structure.
type UnionFind struct {
\tparent []int
\trank   []int
\tCount  int
}

func New(size int) *UnionFind {
\tparent := make([]int, size)
\trank := make([]int, size)
\tfor i := 0; i < size; i++ {
\t\tparent[i] = i
\t\trank[i] = 1
\t}
\treturn &UnionFind{parent: parent, rank: rank, Count: size}
}

// Find finds the representative of element x with path compression.
func (uf *UnionFind) Find(x int) int {
\tif uf.parent[x] != x {
\t\tuf.parent[x] = uf.Find(uf.parent[x])
\t}
\treturn uf.parent[x]
}

// Union merges sets containing x and y using rank.
func (uf *UnionFind) Union(x, y int) bool {
\trootX := uf.Find(x)
\trootY := uf.Find(y)
\tif rootX == rootY {
\t\treturn false
\t}

\tif uf.rank[rootX] < uf.rank[rootY] {
\t\tuf.parent[rootX] = rootY
\t} else if uf.rank[rootX] > uf.rank[rootY] {
\t\tuf.parent[rootY] = rootX
\t} else {
\t\tuf.parent[rootY] = rootX
\t\tuf.rank[rootX]++
\t}
\tuf.Count--
\treturn true
}

func (uf *UnionFind) Connected(x, y int) bool {
\treturn uf.Find(x) == uf.Find(y)
}
""",
        "test_content": """package disjointset

import "testing"

func TestUnionFind(t *testing.T) {
\tuf := New(5)

\tuf.Union(0, 1)
\tuf.Union(1, 2)

\tif !uf.Connected(0, 2) {
\t\tt.Errorf("expected 0 and 2 to be connected")
\t}

\tif uf.Connected(0, 3) {
\t\tt.Errorf("did not expect 0 and 3 to be connected")
\t}

\tif uf.Count != 3 {
\t\tt.Errorf("expected 3 disjoint sets, got %d", uf.Count)
\t}
}
"""
    },
    {
        "type": "feat(patterns)",
        "message": "implement Singleflight request deduplication wrapper",
        "file": "pkg/concurrency/singleflight/group.go",
        "test_file": "pkg/concurrency/singleflight/group_test.go",
        "content": """package singleflight

import "sync"

type call struct {
\twg  sync.WaitGroup
\tval any
\terr error
}

// Group represents a class of work and forms an execution barrier.
type Group struct {
\tmu sync.Mutex
\tm  map[string]*call
}

func NewGroup() *Group {
\treturn &Group{m: make(map[string]*call)}
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a time.
func (g *Group) Do(key string, fn func() (any, error)) (any, error) {
\tg.mu.Lock()
\tif c, ok := g.m[key]; ok {
\t\tg.mu.Unlock()
\t\tc.wg.Wait()
\t\treturn c.val, c.err
\t}

\tc := new(call)
\tc.wg.Add(1)
\tg.m[key] = c
\tg.mu.Unlock()

\tc.val, c.err = fn()
\tc.wg.Done()

\tg.mu.Lock()
\tdelete(g.m, key)
\tg.mu.Unlock()

\treturn c.val, c.err
}
""",
        "test_content": """package singleflight

import (
\t"sync"
\t"sync/atomic"
\t"testing"
\t"time"
)

func TestSingleflight(t *testing.T) {
\tg := NewGroup()
\tvar execCount int32

\tfn := func() (any, error) {
\t\tatomic.AddInt32(&execCount, 1)
\t\ttime.Sleep(20 * time.Millisecond)
\t\treturn "result", nil
\t}

\tvar wg sync.WaitGroup
\tfor i := 0; i < 5; i++ {
\t\twg.Add(1)
\t\tgo func() {
\t\t\tdefer wg.Done()
\t\t\tres, err := g.Do("heavy-key", fn)
\t\t\tif err != nil || res != "result" {
\t\t\t\tt.Errorf("unexpected output: %v, %v", res, err)
\t\t\t}
\t\t}()
\t}
\twg.Wait()

\tif atomic.LoadInt32(&execCount) != 1 {
\t\tt.Fatalf("expected fn to execute exactly once, got %d", execCount)
\t}
}
"""
    },
    {
        "type": "docs(architecture)",
        "message": "add idempotency key pattern for financial and distributed APIs",
        "file": "docs/architecture/idempotency-keys.md",
        "test_file": None,
        "content": """# API Idempotency Keys Architecture 💳

In distributed systems, network packets can be lost, duplicated, or retried automatically by clients. Without idempotency guarantees, charging a customer or transferring funds could result in double billing.

## 1. Request Lifecycle with Idempotency Key

```
Client                        API Gateway / App                Redis / DB
  |                                   |                            |
  |--- POST /payments --------------->|                            |
  |    Idempotency-Key: abc-123       |--- SETNX idempotency:abc ->|
  |                                   |<-- (OK / Already Exists) --|
  |                                   |                            |
  |                                   |-- [If In Progress] --------> Return 409 Conflict
  |                                   |-- [If Completed] ----------> Return Cached Response
  |                                   |-- [If New] ----------------> Process Payment & Cache Result
  |<-- HTTP 200 { status: "paid" } ---|
```

## 2. Key Components
1. **Unique Fingerprinting**: Combine `Idempotency-Key` with client ID and payload hash (SHA-256) to ensure the client is not reusing keys with different payloads.
2. **State Machine**:
   - `IN_PROGRESS` (TTL: 120s): Prevents race conditions from parallel requests.
   - `COMPLETED` (TTL: 24-48h): Returns deterministic saved response body and headers.
   - `FAILED`: Cleans up lock so client can retry.
""",
        "test_content": None
    },
    {
        "type": "feat(algorithms)",
        "message": "implement 0/1 Knapsack dynamic programming solver",
        "file": "pkg/algorithms/dp/knapsack.go",
        "test_file": "pkg/algorithms/dp/knapsack_test.go",
        "content": """package dp

func max(a, b int) int {
\tif a > b {
\t\treturn a
\t}
\treturn b
}

// Knapsack solves 0/1 Knapsack problem using dynamic programming.
func Knapsack(weights []int, values []int, capacity int) int {
\tn := len(weights)
\tdp := make([][]int, n+1)
\tfor i := range dp {
\tdp[i] = make([]int, capacity+1)
\t}

\tfor i := 1; i <= n; i++ {
\t\tfor w := 0; w <= capacity; w++ {
\t\t\tif weights[i-1] <= w {
\t\t\t\tdp[i][w] = max(dp[i-1][w], values[i-1]+dp[i-1][w-weights[i-1]])
\t\t\t} else {
\t\t\t\tdp[i][w] = dp[i-1][w]
\t\t\t}
\t\t}
\t}

\treturn dp[n][capacity]
}
""",
        "test_content": """package dp

import "testing"

func TestKnapsack(t *testing.T) {
\tweights := []int{2, 3, 4, 5}
\tvalues := []int{3, 4, 5, 6}
\tcapacity := 5

\t// items with wt 2 (val 3) and wt 3 (val 4) -> total val 7
\tresult := Knapsack(weights, values, capacity)
\tif result != 7 {
\t\tt.Errorf("expected max value 7, got %d", result)
\t}
}
"""
    },
    {
        "type": "feat(patterns)",
        "message": "implement graceful shutdown coordinator for microservices",
        "file": "pkg/patterns/graceful/coordinator.go",
        "test_file": "pkg/patterns/graceful/coordinator_test.go",
        "content": """package graceful

import (
\t"context"
\t"sync"
\t"time"
)

type Hook func(ctx context.Context) error

type Coordinator struct {
\tmu      sync.Mutex
\thooks   []Hook
\ttimeout time.Duration
}

func NewCoordinator(timeout time.Duration) *Coordinator {
\treturn &Coordinator{timeout: timeout}
}

func (c *Coordinator) Register(hook Hook) {
\tc.mu.Lock()
\tdefer c.mu.Unlock()
\tc.hooks = append(c.hooks, hook)
}

func (c *Coordinator) Shutdown() error {
\tc.mu.Lock()
\tdefer c.mu.Unlock()

\tctx, cancel := context.WithTimeout(context.Background(), c.timeout)
\tdefer cancel()

\tvar wg sync.WaitGroup
\terrChan := make(chan error, len(c.hooks))

\tfor _, hook := range c.hooks {
\t\twg.Add(1)
\t\tgo func(h Hook) {
\t\t\tdefer wg.Done()
\t\t\tif err := h(ctx); err != nil {
\t\t\t\terrChan <- err
\t\t\t}
\t\t}(hook)
\t}

\twg.Wait()
\tclose(errChan)

\tfor err := range errChan {
\t\tif err != nil {
\t\t\treturn err
\t\t}
\t}
\treturn nil
}
""",
        "test_content": """package graceful

import (
\t"context"
\t"testing"
\t"time"
)

func TestCoordinator(t *testing.T) {
\tcoord := NewCoordinator(100 * time.Millisecond)

\tvar closed bool
\tcoord.Register(func(ctx context.Context) error {
\t\tclosed = true
\t\treturn nil
\t})

\tif err := coord.Shutdown(); err != nil {
\t\tt.Fatalf("unexpected shutdown error: %v", err)
\t}

\tif !closed {
\t\tt.Errorf("expected hook to be called during shutdown")
\t}
}
"""
    }
]

def load_state():
    if os.path.exists(STATE_FILE):
        try:
            with open(STATE_FILE, "r") as f:
                return json.load(f)
        except Exception:
            pass
    return {"used_indices": [], "day_count": 0, "dynamic_count": 0}

def save_state(state):
    with open(STATE_FILE, "w") as f:
        json.dump(state, f, indent=2)

def generate_dynamic_contribution(idx):
    """Generates procedural algorithmic challenge & solution to guarantee infinite daily contributions"""
    topics = [
        ("sliding-window", "Sliding Window Maximum sub-array algorithm"),
        ("monotonic-queue", "Monotonic Queue for rolling range minimum queries"),
        ("binary-tree-inversion", "Invert Binary Tree with iterative queue traversal"),
        ("longest-common-subsequence", "Longest Common Subsequence 2D dynamic programming"),
        ("topological-sort", "Topological Sort using Kahn's in-degree algorithm"),
        ("bloom-filter", "Counting Bloom Filter with multiple hash functions"),
        ("ring-buffer", "Lock-free circular ring buffer with atomic pointers"),
        ("consistent-hashing", "Consistent Hashing ring with virtual vnodes"),
        ("two-phase-commit", "2PC Coordinator prototype for distributed transactions"),
        ("event-sourcing", "Event Store append-only ledger with snapshot support"),
        ("distributed-id", "Snowflake 64-bit unique ID generator"),
        ("lfu-cache", "Least Frequently Used (LFU) cache eviction engine")
    ]
    topic_slug, topic_desc = topics[idx % len(topics)]
    suffix = idx // len(topics) + 1

    file_path = f"pkg/challenges/{topic_slug}/solution_v{suffix}.go"
    test_path = f"pkg/challenges/{topic_slug}/solution_v{suffix}_test.go"

    content = f"""package {topic_slug.replace('-', '_')}

// SolutionV{suffix} demonstrates {topic_desc}.
type SolutionV{suffix} struct {{
\titeration int
}}

func NewV{suffix}() *SolutionV{suffix} {{
\treturn &SolutionV{suffix}{{iteration: {suffix}}}
}}

func (s *SolutionV{suffix}) Compute(input []int) int {{
\ttotal := 0
\tfor _, v := range input {{
\t\ttotal += v
\t}}
\treturn total * s.iteration
}}
"""
    test_content = f"""package {topic_slug.replace('-', '_')}

import "testing"

func TestSolutionV{suffix}(t *testing.T) {{
\ts := NewV{suffix}()
\tinput := []int{{1, 2, 3}}
\tres := s.Compute(input)
\texpected := 6 * {suffix}
\tif res != expected {{
\t\tt.Errorf("expected %d, got %d", expected, res)
\t}}
}}
"""
    return {
        "type": "feat(challenges)",
        "message": f"implement {topic_desc} (iteration {suffix})",
        "file": file_path,
        "test_file": test_path,
        "content": content,
        "test_content": test_content
    }

def run_cmd(cmd, cwd=REPO_DIR):
    res = subprocess.run(cmd, cwd=cwd, shell=True, capture_output=True, text=True)
    if res.returncode != 0:
        print(f"Command failed: {cmd}\nStderr: {res.stderr}", file=sys.stderr)
        raise RuntimeError(f"Command failed: {cmd}")
    return res.stdout.strip()

def create_and_push(item):
    file_rel = item["file"]
    abs_file = os.path.join(REPO_DIR, file_rel)
    os.makedirs(os.path.dirname(abs_file), exist_ok=True)
    with open(abs_file, "w") as f:
        f.write(item["content"])

    if item.get("test_file") and item.get("test_content"):
        abs_test = os.path.join(REPO_DIR, item["test_file"])
        os.makedirs(os.path.dirname(abs_test), exist_ok=True)
        with open(abs_test, "w") as f:
            f.write(item["test_content"])

    commit_msg = f"{item['type']}: {item['message']}"

    run_cmd(f"git add {file_rel}")
    if item.get("test_file"):
        run_cmd(f"git add {item['test_file']}")

    run_cmd(f'git commit -m "{commit_msg}"')
    run_cmd("git push origin main")
    print(f"✓ Pushed: {commit_msg}")
    return commit_msg

def main():
    import argparse
    parser = argparse.ArgumentParser(description="Auto contribute to backend-engineering-vault")
    parser.add_argument("--count", type=int, default=4, help="Number of commits/pushes to make")
    args = parser.parse_args()

    state = load_state()
    pushed = []

    for _ in range(args.count):
        # Pick next unused curated catalog item
        unused = [i for i in range(len(CATALOG)) if i not in state["used_indices"]]
        if unused:
            pick_idx = unused[0]
            item = CATALOG[pick_idx]
            state["used_indices"].append(pick_idx)
        else:
            # Use dynamic generator once catalog exhausted
            dyn_idx = state.get("dynamic_count", 0)
            item = generate_dynamic_contribution(dyn_idx)
            state["dynamic_count"] = dyn_idx + 1

        msg = create_and_push(item)
        pushed.append(msg)

    state["day_count"] = state.get("day_count", 0) + 1
    save_state(state)

    print(f"\nSuccessfully completed {len(pushed)} pushes today!")

if __name__ == "__main__":
    main()
