# BeatFlow — Testing Strategy

## 1. Purpose

This document defines how BeatFlow will be tested.

Testing is required because BeatFlow includes critical flows such as authentication, payment, order creation, exclusive license sale, and secure file downloads.

---

## 2. Test Types

### 2.1 Unit Tests

Used for pure business logic.

Examples:

- password validation
- price calculation
- trending score
- role checks
- license rule validation

### 2.2 Handler Tests

Used for API behavior.

Examples:

- invalid request body returns 400
- unauthorized request returns 401
- forbidden request returns 403
- successful request returns expected JSON

### 2.3 Service Tests

Used for business use cases.

Examples:

- register user
- create beat
- submit beat for review
- create order
- process payment completed

### 2.4 Repository Tests

Used for database queries.

Use Testcontainers with PostgreSQL.

### 2.5 Integration Tests

Used for multiple components.

Examples:

- Auth Service + PostgreSQL
- Order Service + PostgreSQL + RabbitMQ
- Marketplace Service + Elasticsearch
- Rate limiter + Redis

### 2.6 Worker Tests

Used for event consumers.

Examples:

- beat.published indexes Elasticsearch
- payment.completed marks order paid
- message.sent creates notification

### 2.7 Contract Tests

Used for service-to-service APIs.

Options:

- OpenAPI contract tests
- Pact
- gRPC/proto tests later

### 2.8 E2E Tests

Used for critical flows.

Examples:

- register → create producer → create beat → approve → search
- cart → order → payment webhook → download URL
- message producer → notification

### 2.9 Load Tests

Tool:

- k6

Scenarios:

- search endpoint load
- login burst
- order creation load
- download URL generation
- marketplace listing

---

## 3. Tools

- Go test
- httptest
- Fiber test utilities
- Testcontainers
- gomock or mockery
- k6
- golangci-lint
- go vet
- gosec
- Trivy

---

## 4. Critical Flows Requiring Tests

- Register
- Login
- Refresh token
- RBAC checks
- Producer profile creation
- Create beat
- Submit beat for review
- Admin approve/reject
- License creation
- Exclusive license purchase
- Cart validation
- Order creation
- Payment webhook idempotency
- Download authorization
- Search indexing
- Rate limiter
- Retry/circuit breaker behavior

---

## 5. Testcontainers

Use Testcontainers for:

- PostgreSQL
- Redis
- RabbitMQ
- Elasticsearch
- Couchbase experiment
- LocalStack for S3 experiment if needed

---

## 6. CI Testing

CI should run:

- gofmt check
- go vet
- golangci-lint
- unit tests
- selected integration tests
- gosec
- Trivy
- Docker build

---

## 7. Load Testing with k6

Example scenarios:

```txt
GET /api/v1/beats/search
POST /api/v1/auth/login
POST /api/v1/orders
GET /api/v1/downloads/{id}
```

Metrics:

- p95 latency
- p99 latency
- error rate
- throughput
- saturation

---

## 8. Plane Work Items

```txt
testing: setup Go unit test structure
testing: add Testcontainers PostgreSQL setup
testing: add Redis integration tests
testing: add RabbitMQ consumer tests
testing: add Elasticsearch integration tests
testing: add payment webhook idempotency tests
testing: add k6 load testing plan
testing: add CI test workflow
```
