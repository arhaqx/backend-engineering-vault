# Distributed Locks: Redis & Redlock Architecture 🔒

Distributed locks serialize operations across competing distributed nodes where in-process mutexes fail.

## 1. Single-Instance Lock Pitfall
Using `SETNX lock_key token PX 10000` is efficient but vulnerable to single-point-of-failure (SPoF). If master crashes before replication syncs to replica, multiple nodes can acquire the identical lock simultaneously.

## 2. Redlock Multi-Node Algorithm
Proposed by Salvatore Sanfilippo (antirez), Redlock uses $N$ independent Redis masters (typically 5):

1. Client fetches current timestamp $T_1$.
2. Attempts to acquire lock across all $N$ instances sequentially using identical key/value and short timeouts.
3. Lock is considered **acquired** only if:
   - Client acquired lock in at least $\lfloor N/2 \rfloor + 1$ nodes.
   - Total time elapsed $T_2 - T_1 < \text{Lock Validity Time}$.
4. If acquisition fails, client broadcasts an immediate unlock to all instances.

## 3. Martin Kleppmann's Critique & Fencing Tokens
GC pauses or network delays can expire lock validity while the client continues executing assumed exclusive work.

**Remedy**: Always pair distributed locks with monotonic **fencing tokens** verified at the storage layer:
```sql
UPDATE accounts SET balance = balance - 100 WHERE id = 1 AND fencing_token < :current_token;
```
