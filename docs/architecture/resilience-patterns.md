# BeatFlow — Resilience Patterns

## 1. Purpose

This document defines the resilience patterns used in BeatFlow.

Microservices fail in many ways:

- Network timeout
- Downstream service unavailable
- Slow dependency
- Database connection issue
- Queue consumer crash
- Duplicate message
- Partial failure
- Pod termination

Resilience patterns make these failures controlled and observable.

---

## 2. Timeout Management

Every outbound call must have a timeout.

Examples:

- HTTP call from Order Service to Beat Service
- HTTP call from File Service to Order Service
- PostgreSQL query
- Redis command
- Elasticsearch query
- S3 request

Go pattern:

```go
ctx, cancel := context.WithTimeout(parentCtx, 3*time.Second)
defer cancel()
```

Rules:

- No unbounded outbound calls.
- Timeout values must be configurable.
- Long-running jobs should use separate job-level timeout.
- Timeout errors should be logged with dependency name.

---

## 3. Context Propagation

Context should carry:

- request_id
- correlation_id
- trace context
- deadline
- cancellation signal

HTTP boundaries:

- Incoming request extracts request/correlation/trace headers.
- Outgoing request forwards correlation and trace headers.

Async boundaries:

- RabbitMQ messages should include correlation ID.
- Event headers should include trace context when possible.
- Consumers should continue trace if trace context exists.

---

## 4. Retry with Backoff

Retries are useful for transient failures.

Use retries for:

- temporary network errors
- 502/503/504 responses
- temporary S3 errors
- temporary Elasticsearch errors
- temporary RabbitMQ publish errors

Avoid retries for:

- validation errors
- authorization errors
- insufficient balance
- duplicate email
- business rule failures

Rules:

- Use exponential backoff.
- Add jitter.
- Limit max attempts.
- Make operation idempotent before retrying.
- Log final failure.

Example policy:

```txt
max_attempts: 3
initial_delay: 100ms
max_delay: 2s
jitter: true
```

---

## 5. Circuit Breaker

Circuit breaker prevents continuously calling a failing dependency.

States:

- closed
- open
- half-open

Use circuit breaker for:

- payment provider
- search service
- external email provider
- internal service calls that can fail repeatedly

Rules:

- Do not wrap every tiny call blindly.
- Use circuit breaker for critical and failure-prone dependencies.
- Configure failure threshold.
- Configure open duration.
- Configure half-open trial requests.
- Business stakeholders should understand behavior for critical flows like payment.

Go package option:

```txt
sony/gobreaker
```

---

## 6. Rate Limiting

Rate limiting protects services from abuse and overload.

Use cases:

- login
- register
- refresh token
- payment webhook
- search
- download URL generation
- message sending

Redis-based algorithms:

- fixed window
- sliding window
- token bucket
- leaky bucket

Recommended:

- Token bucket for general API rate limiting.
- Sliding window for sensitive auth endpoints.

Rate limiting can exist at multiple layers:

- application
- Nginx/API gateway
- Kubernetes ingress
- service mesh later

---

## 7. Bulkhead

Bulkhead isolates resources.

Examples:

- Separate worker pools for payment and notifications.
- Separate database connection pools per service.
- Separate queues for high-priority and low-priority jobs.

Goal:

- Failure in notification delivery should not block payment processing.

---

## 8. Graceful Shutdown

Every Go service should handle shutdown signals.

Requirements:

- Stop accepting new requests.
- Finish in-flight requests within timeout.
- Close database pool.
- Close Redis client.
- Stop RabbitMQ consumers safely.
- Flush logs.
- Exit cleanly.

Kubernetes needs graceful shutdown because pods can be terminated during rollout.

---

## 9. Idempotency

Required for:

- payment webhook
- order creation
- event consumers
- download token generation
- search indexing
- notification sending

Patterns:

- Idempotency key table
- provider_event_id unique constraint
- processed_events table
- Redis idempotency key with TTL

---

## 10. Dead-Letter Queue

DLQ stores messages that cannot be processed.

Use DLQ for:

- invalid event payload
- repeated consumer failure
- poison messages
- unavailable dependency after retries

DLQ message should include:

- original event
- failure reason
- retry count
- failed_at
- consumer name

---

## 11. Transactional Outbox

Problem:

- Service writes to database.
- Service tries to publish event.
- Service crashes before publishing.

Solution:

- Write business data and outbox event in same DB transaction.
- Separate publisher reads outbox table.
- Publisher sends message to RabbitMQ.
- Publisher marks event as published.

Use cases:

- beat.published
- order.created
- order.paid
- payment.completed
- message.sent

---

## 12. Inbox / Idempotent Consumer

Problem:

- Broker can deliver same message more than once.
- Consumer can crash after processing but before ack.

Solution:

- Store processed event IDs.
- Skip duplicates.
- Make consumer side effects idempotent.

Table:

```txt
processed_events
- event_id
- consumer_name
- processed_at
```

---

## 13. Fallback

Fallback is an alternative behavior when dependency fails.

Examples:

- If Elasticsearch is down, return service_unavailable or fallback to PostgreSQL basic listing.
- If notification service is down, keep order flow successful and retry notification later.
- If analytics is down, drop or buffer event depending on priority.

---

## 14. Resilience Work Items

Plane work items:

```txt
architecture: design timeout and retry strategy
architecture: design circuit breaker strategy
security: design Redis-based rate limiter
architecture: design transactional outbox pattern
architecture: design idempotent consumer pattern
observability: propagate correlation IDs across async messages
infra: add DLQ configuration for RabbitMQ
```
