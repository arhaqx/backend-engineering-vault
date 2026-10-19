# API Idempotency Keys Architecture 💳

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
