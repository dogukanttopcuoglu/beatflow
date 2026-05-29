# BeatFlow — Microservice Architecture Decision

## 1. Purpose

This document explains the decision to design BeatFlow as a microservice-based backend system using a monorepo structure.

The project is intentionally designed to practice production-ready backend concepts.

---

## 2. Decision

BeatFlow will be designed as:

```txt
Architecture: Microservice-based backend
Repository strategy: Monorepo
External API: REST over HTTP
Internal sync communication: REST initially, gRPC optional later
Async communication: RabbitMQ
Event streaming comparison: Kafka
Primary DB: PostgreSQL
Cache/temp data: Redis
Search read model: Elasticsearch
Object storage: AWS S3
NoSQL learning track: Couchbase
Local runtime: Docker Compose
Production-like orchestration: Kubernetes
GitOps: ArgoCD
CI/CD: GitHub Actions
Observability: Prometheus, Grafana, OpenTelemetry, Jaeger
```

---

## 3. Why Microservices?

Microservices allow learning:

- service boundaries
- independent deployability
- database ownership
- API contracts
- event-driven workflows
- distributed consistency
- async processing
- Docker Compose multi-service development
- Kubernetes deployment
- observability across services
- CI/CD for multiple services
- resilience patterns

---

## 4. Risks

Microservices introduce:

- boilerplate
- network failures
- distributed transactions
- debugging complexity
- observability requirements
- deployment complexity
- testing complexity

Mitigation:

- monorepo
- incremental implementation
- documentation-first architecture
- Docker Compose locally
- clear service ownership
- strong observability
- outbox/inbox patterns
- integration tests

---

## 5. Initial Services

- API Gateway
- Auth Service
- User Service
- Beat Service
- File Service
- Marketplace Service
- Order Service
- Payment Service
- Notification Service
- Analytics Service
- Messaging Service
- Admin Service
- Personalization Service optional for Couchbase
- Streaming/Analytics Worker optional for Kafka

---

## 6. API Gateway

Initial implementation:

- Nginx

Responsibilities:

- route external requests
- CORS
- request size limits
- reverse proxy
- gateway-level rate limiting

Future options:

- Kong
- Traefik
- Envoy
- Go Fiber BFF/Gateway

---

## 7. Service Communication

### Synchronous

Use REST initially.

Examples:

- Order Service → Beat Service for license validation
- File Service → Order Service for entitlement check
- Admin Service → Beat Service for moderation action

Rules:

- Use timeout.
- Propagate context.
- Add retry only for safe operations.
- Add circuit breaker for critical dependencies.

### Asynchronous

Use RabbitMQ.

Examples:

- beat.published → Marketplace Service
- payment.completed → Order Service
- order.paid → Notification and Analytics
- message.sent → Notification

### Streaming

Kafka may be added later for analytics streams.

---

## 8. Database Ownership

Rules:

- Each service owns its data.
- No service directly queries another service database.
- Cross-service access happens through API or events.
- Local development may use one PostgreSQL instance with multiple databases/schemas.

---

## 9. Source of Truth

- PostgreSQL: transactional source of truth
- Redis: temporary/cache data
- Elasticsearch: search projection
- S3: file storage
- Couchbase: NoSQL learning/experimental document storage
- RabbitMQ: message transport
- Kafka: event stream option

---

## 10. Resilience Requirements

Every service should eventually support:

- graceful shutdown
- timeout
- retry/backoff
- circuit breaker for critical outbound calls
- rate limiter where needed
- health check
- readiness check
- metrics endpoint
- request/correlation ID
- trace propagation

---

## 11. Deployment

Local:

- Docker Compose

Production-like:

- Kubernetes
- ArgoCD GitOps

---

## 12. Observability

Required:

- Zap logs
- Prometheus metrics
- Grafana dashboards
- OpenTelemetry traces
- Jaeger visualization

---

## 13. Security

Required:

- JWT
- RBAC
- ownership checks
- password hashing
- webhook signature verification
- S3 private access
- container hardening
- Kubernetes secrets
- gosec
- Trivy

---

## 14. Summary

BeatFlow is intentionally designed as a production-ready microservice learning project.

The architecture includes both core business functionality and operational concerns expected from modern backend systems.
