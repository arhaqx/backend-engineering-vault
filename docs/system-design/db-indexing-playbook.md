# Database Indexing & Query Optimization Playbook 🗄️

Effective index design is the single highest leverage activity in relational database performance tuning.

## 1. B-Tree vs LSM-Tree Storage Engines

| Feature | B-Tree (PostgreSQL, MySQL InnoDB) | LSM-Tree (RocksDB, Cassandra, Scylla) |
|---|---|---|
| **Read Latency** | Constant $O(\log N)$, minimal disk hops | Variable (MemTable -> Bloom Filter -> SSTables) |
| **Write Throughput**| Random I/O updates in-place | Append-only sequential writes (very high throughput) |
| **Write Amplification** | Higher due to page splits | Managed through background compaction |
| **Ideal Workload** | High OLTP, transactional consistency | Write-heavy telemetry, time-series, log ingestion |

## 2. Composite Index Cardinality & Column Ordering

When creating compound indexes `INDEX (col_a, col_b, col_c)`:

1. **Equality First**: Place columns evaluated with equality (`=`) at the leading edge.
2. **Cardinality Priority**: Higher cardinality (more distinct values) provides faster selectivity filtering.
3. **Range Last**: The moment a range operator (`<`, `>`, `BETWEEN`, `LIKE 'prefix%'`) is evaluated on column $K$, subsequent columns in the composite index cannot be utilized for index filtering.

## 3. Covering Indexes (Index-Only Scans)

By including all requested `SELECT` columns inside the index leaf node (or via PostgreSQL `INCLUDE` clause):

```sql
CREATE INDEX idx_orders_customer_status_covering 
ON orders (customer_id, status) 
INCLUDE (total_amount, created_at);
```

The database query engine fulfills the query strictly from RAM-cached index buffers without fetching heap tuples from disk, avoiding expensive table random access.
