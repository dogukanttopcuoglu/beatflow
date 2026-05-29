# BeatFlow — System Requirements

## 1. Purpose

This document defines the technical requirements for BeatFlow.

Product requirements describe what the system should do. System requirements describe how the system should behave technically.

---

## 2. Functional System Requirements

### 2.1 API

- All external APIs should be RESTful.
- API versioning should start with `/api/v1`.
- JSON should be the default request/response format.
- Error responses should be consistent across services.
- Pagination, filtering, and sorting should be standardized.

### 2.2 Authentication

- JWT access tokens are used for API authentication.
- Refresh tokens are used for session continuation.
- Passwords are hashed.
- Sensitive auth operations should be rate-limited.

### 2.3 Authorization

- RBAC is required.
- Ownership checks are required.
- Admin actions are audited.

### 2.4 Data Storage

- PostgreSQL stores transactional data.
- Redis stores temporary/cache data.
- Elasticsearch stores search projections.
- S3 stores files.
- Couchbase is used only for NoSQL learning/experimental domain unless later promoted.

### 2.5 Messaging

- RabbitMQ handles async workflows.
- DLQ must exist for critical queues.
- Event handlers should be idempotent.
- Kafka is documented for stream processing and can be added later.

### 2.6 Search

- Search must use Elasticsearch.
- Search index is derived from events.
- PostgreSQL remains source of truth.

### 2.7 File Delivery

- Private files are never directly public.
- Downloads use presigned S3 URLs.
- Uploads use presigned S3 URLs.
- Download authorization is checked by backend.

---

## 3. Non-Functional System Requirements

### 3.1 Availability

- Core read APIs should remain available when non-critical async workers are down.
- Payment and order consistency is more important than availability.
- Search can degrade if Elasticsearch is down.

### 3.2 Consistency

- Payment/order/download data requires strong consistency.
- Search and analytics can be eventually consistent.
- Notifications can be eventually consistent.

### 3.3 Performance

Target response times:

```txt
Auth endpoints: < 200ms typical
Beat detail: < 200ms typical
Search: < 300ms typical
Order creation: < 500ms typical
Download URL generation: < 300ms typical
```

### 3.4 Scalability

- Stateless services should scale horizontally.
- Kubernetes HPA should be used for selected services.
- Redis and Elasticsearch should be monitored for bottlenecks.
- RabbitMQ queue lag should be monitored.

### 3.5 Reliability

- Critical operations must be idempotent.
- Retry should be used for transient failures.
- Circuit breakers should protect unstable dependencies.
- Timeouts must be used on outbound calls.
- Graceful shutdown must prevent dropped in-flight requests.

### 3.6 Security

- Secrets must not be committed.
- Services should run with least privilege.
- Containers should be hardened.
- Inputs must be validated.
- Webhook signatures must be verified.
- Rate limiting must protect sensitive endpoints.

### 3.7 Observability

Each service should expose:

```txt
GET /healthz
GET /readyz
GET /metrics
```

Each service should support:

- structured logs
- request ID
- correlation ID
- trace ID
- metrics
- traces

---

## 4. Scale Assumptions

Initial learning-scale assumptions:

```txt
Users: 10k
Producers: 1k
Beats: 100k
Daily searches: 100k
Daily plays: 500k
Daily purchases: 1k
```

These are not production guarantees. They are design assumptions for learning.

---

## 5. Critical Flows

### 5.1 Register/Login

Must be secure, rate-limited, logged, and tested.

### 5.2 Beat Publish

Must update:

- PostgreSQL
- event stream
- Elasticsearch projection
- notification system

### 5.3 Checkout/Payment

Must be idempotent and consistent.

### 5.4 Download

Must enforce:

- auth
- ownership
- paid order
- license file access
- download limit

### 5.5 Search Indexing

Must handle:

- duplicate events
- stale events
- retries
- Elasticsearch temporary failure

---

## 6. Failure Scenarios

The system should consider:

- PostgreSQL unavailable
- Redis unavailable
- RabbitMQ unavailable
- Elasticsearch unavailable
- S3 unavailable
- Payment webhook duplicate
- Payment webhook out of order
- Event consumer crash
- Service restart during request
- Kubernetes pod termination
- Queue message poison event
- Network timeout
- Downstream service returns 500
- Downstream service is slow

---

## 7. Required Engineering Concepts

BeatFlow must cover:

- REST API
- Microservices
- PostgreSQL
- Redis
- RabbitMQ
- Elasticsearch
- S3
- Couchbase experiment
- Kafka comparison
- Docker
- Kubernetes
- ArgoCD
- CI/CD
- Observability
- Resilience patterns
- Security
- Testing
- Static analysis
- Load testing
- Documentation
