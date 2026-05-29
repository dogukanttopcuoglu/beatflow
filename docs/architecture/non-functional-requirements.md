# BeatFlow — Non-Functional Requirements

## 1. Purpose

Non-functional requirements define how the system behaves in production-like conditions.

A service is not production-ready only because its business logic works. It must also be observable, reliable, secure, scalable, testable, and operable.

---

## 2. Availability

Requirements:

- Public marketplace browsing should remain available even if notification workers are down.
- Order/payment flows should prefer correctness over availability.
- Search can degrade gracefully if Elasticsearch is unavailable.
- Admin workflows can tolerate lower availability than checkout.

Design notes:

- Use readiness checks to remove unhealthy pods from traffic.
- Use liveness checks to restart stuck services.
- Use graceful shutdown to avoid cutting active requests.

---

## 3. Reliability

Requirements:

- Critical operations must be idempotent.
- Event consumers must be safe against duplicate messages.
- Webhooks must be safe against duplicate provider events.
- Retries must use backoff.
- Poison messages must go to DLQ.
- Outbound calls must have timeouts.

---

## 4. Performance

Requirements:

- Avoid blocking API requests on slow async work.
- Use Redis for hot data.
- Use Elasticsearch for search instead of heavy PostgreSQL full-text queries.
- Use S3 direct upload through presigned URLs to avoid large file uploads through API servers.
- Use pagination everywhere.

---

## 5. Scalability

Requirements:

- Services should be stateless unless explicitly designed otherwise.
- Stateless services should scale horizontally.
- Kubernetes HPA should be used for selected services.
- Message consumers should scale based on queue depth when possible.
- Search and analytics should be decoupled from write path.

---

## 6. Security

Requirements:

- Passwords must be hashed.
- JWT and refresh tokens must be handled securely.
- Sensitive endpoints must be rate-limited.
- Webhook signatures must be verified.
- S3 private files must not be publicly accessible.
- Admin actions must be audited.
- Containers must not run as root when possible.
- Images should be minimal and scanned.

---

## 7. Observability

Requirements:

- Structured logs
- Request ID
- Correlation ID
- Trace ID
- Prometheus metrics
- OpenTelemetry traces
- Grafana dashboards
- Jaeger trace visualization
- Queue lag metrics
- Error rate metrics
- Latency histograms

---

## 8. Operability

Requirements:

- Docker Compose for local development.
- Kubernetes manifests for production-like deployment.
- ArgoCD GitOps documentation.
- Health and readiness probes.
- ConfigMap and Secret usage.
- Resource requests and limits.
- Runbooks for common failures.

---

## 9. Maintainability

Requirements:

- Clear service boundaries.
- No cross-service database access.
- Shared packages only for technical concerns.
- Domain logic stays inside owning service.
- Documentation updated with architecture decisions.
- ADRs are used for major decisions.

---

## 10. Testability

Requirements:

- Business logic should be unit-testable.
- Repositories should be integration-testable.
- Handlers should be tested with httptest or Fiber test tools.
- Critical flows should have integration/E2E tests.
- Testcontainers should be used for PostgreSQL, Redis, RabbitMQ, Elasticsearch, and Couchbase where needed.

---

## 11. Production-Ready Checklist

A service is production-ready when it has:

- Config loading
- Structured logging
- Graceful shutdown
- `/healthz`
- `/readyz`
- `/metrics`
- Timeouts
- Retry policy where needed
- Circuit breaker for critical outbound calls
- Rate limiting where needed
- Input validation
- Error handling
- Unit/integration tests
- Dockerfile
- Kubernetes manifest
- CI checks
- Observability integration
