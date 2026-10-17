# Distributed Caching Strategies & Consistency Patterns ⚡

Caching reduces read latency and offloads primary database engines. Selecting the correct invalidation and consistency strategy depends heavily on data criticality and update frequency.

## 1. Core Invalidation Patterns

### Cache-Aside (Lazy Loading)
- Application reads cache first. On miss, it reads database, populates cache, and returns.
- **Pros**: Only requested data is cached; cache node failure does not halt application writes.
- **Cons**: Cache miss penalty on initial read; stale data window if database updates without cache eviction.

### Read-Through & Write-Through
- Application treats cache as the primary data store. The cache layer itself loads from and writes to the DB synchronously.
- **Pros**: Simplified application logic; strict consistency guaranteed on write.
- **Cons**: Write latency is increased because data must persist in both cache and DB before returning HTTP 200.

### Write-Behind (Write-Back)
- Application writes to cache immediately. The cache asynchronously batches updates to the persistent database.
- **Pros**: Ultra-low write latency; absorbs massive write spikes.
- **Cons**: Potential data loss if cache instance dies before flushing dirty entries to disk.

## 2. Preventing Cache Stampede & Thundering Herd

When a hot cache key expires under 10,000+ RPS, every incoming worker concurrently queries the DB.

### Solutions:
1. **Mutex / Singleflight (`golang.org/x/sync/singleflight`)**:
   Deduplicates concurrent requests for the identical key, ensuring only 1 DB query executes while other goroutines wait for the shared result.
2. **Probabilistic Early Expiration (XFetch algorithm)**:
   Recomputes and refreshes the key in the background before it reaches physical TTL:
   $$\Delta \cdot \beta \cdot \ln(\text{rand}()) > \text{time\_to\_expire}$$
