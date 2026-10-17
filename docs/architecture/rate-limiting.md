# Rate Limiting Strategies for High-Throughput APIs 🛡️

Rate limiting controls the frequency of incoming traffic to protect downstream dependencies, prevent cascading failures, and enforce fair resource usage among clients.

## 1. Classical Algorithms Comparison

| Algorithm | Pros | Cons | Ideal Use Case |
|---|---|---|---|
| **Token Bucket** | Handles bursts gracefully; memory efficient ($O(1)$) | Token calculation overhead under extreme contention | General public APIs, webhook consumers |
| **Leaky Bucket** | Guarantees smooth constant egress rate | Drops requests aggressively when buffer overflows | Egress traffic shaping, message queue publishers |
| **Fixed Window** | Trivial memory usage ($O(1)$) | Edge-of-window traffic spike vulnerability | Coarse tier-limits (e.g., daily API quota) |
| **Sliding Window Log** | Highly accurate; zero boundary burst anomalies | High memory footprint ($O(N)$ requests in window) | Low volume, high precision security throttling |
| **Sliding Window Counter** | Approximates sliding log with $O(1)$ memory | Approximate calculation around boundaries | Microservice gateway rate limiting |

## 2. Distributed Architecture Implementation

When scaling horizontally across multiple compute nodes:

```
[Client] ---> [Load Balancer] ---> [API Gateway / Envoy]
                                         |
                                (Redis Cluster via Lua Script)
                                         |
                            +------------+------------+
                            |                         |
                       [Allowed]                 [HTTP 429]
                           v
                   [Internal Services]
```

- **Redis Lua Scripts**: Atomicity is critical. Calculating tokens and setting TTLs must happen in a single non-blocking Lua evaluation to prevent race conditions across parallel ingress instances.
- **Fail-Open vs Fail-Closed**: In high-availability architectures, if the rate limiting cache (Redis) becomes unreachable, services should typically fail-open with local in-memory fallback to preserve core business availability.
