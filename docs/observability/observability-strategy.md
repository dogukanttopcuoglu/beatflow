# BeatFlow — Observability Strategy

## 1. Purpose

Observability helps answer:

- Is the system healthy?
- Why is it slow?
- Where did the request fail?
- Which dependency is failing?
- Are queues lagging?
- Are deployments safe?

BeatFlow will include logs, metrics, and traces.

---

## 2. Three Pillars

### 2.1 Logs

Tool:

- Zap

Requirements:

- structured JSON logs in production
- human-readable logs in development
- request_id
- correlation_id
- service name
- error context
- no secrets

### 2.2 Metrics

Tools:

- Prometheus
- Grafana

Metrics:

- request count
- request duration
- error count
- dependency latency
- DB query latency
- queue depth
- consumer lag
- cache hit/miss
- payment webhook failures
- search indexing failures

### 2.3 Traces

Tools:

- OpenTelemetry
- Jaeger

Trace should show:

```txt
API Gateway
  → Auth/Order/Beat Service
  → PostgreSQL
  → RabbitMQ
  → Consumer
  → Elasticsearch/S3
```

---

## 3. Standard Endpoints

Every service should expose:

```txt
GET /healthz
GET /readyz
GET /metrics
```

### /healthz

Answers:

- Is the process alive?

Should not perform heavy dependency checks.

### /readyz

Answers:

- Is the service ready to receive traffic?

May check:

- database connection
- Redis connection if required
- RabbitMQ connection if required
- Elasticsearch connection if required

### /metrics

Prometheus scrape endpoint.

---

## 4. Request Correlation

Every request should have:

- request_id
- correlation_id
- trace_id

Rules:

- If request ID exists, propagate it.
- If not, generate one.
- Include correlation ID in logs.
- Include trace context in outbound HTTP calls.
- Include trace context in RabbitMQ message headers.

---

## 5. Async Tracing

RabbitMQ messages should carry:

- correlation_id
- traceparent
- tracestate

Consumer should extract trace context and continue span.

This allows tracing across async flows.

---

## 6. Prometheus

Prometheus scrapes `/metrics`.

Local Docker Compose should include:

- prometheus
- grafana
- app services

Prometheus config should include service targets.

Example metrics:

```txt
http_requests_total
http_request_duration_seconds
database_query_duration_seconds
rabbitmq_messages_published_total
rabbitmq_messages_consumed_total
search_index_failures_total
payment_webhook_processed_total
```

---

## 7. Grafana

Grafana dashboards should include:

- API latency
- API error rate
- Requests per second
- Service health
- DB latency
- Redis operations
- RabbitMQ queue depth
- Elasticsearch errors
- Payment webhook success/failure
- Kubernetes pod status

---

## 8. Jaeger

Jaeger visualizes distributed traces.

Useful for:

- finding slow dependencies
- seeing request path
- debugging async workflows
- understanding context propagation

---

## 9. Logging Guidelines

Good log:

```go
logger.Info("payment webhook processed",
    zap.String("payment_id", paymentID),
    zap.String("provider_event_id", eventID),
    zap.String("correlation_id", correlationID),
)
```

Bad log:

```go
logger.Info("done")
```

Do not log:

- passwords
- JWT secrets
- refresh tokens
- AWS secrets
- full private S3 URLs
- payment secrets

---

## 10. Alerts

Future alert examples:

- API 5xx error rate > threshold
- Payment webhook failures > threshold
- RabbitMQ DLQ messages > 0
- Elasticsearch indexing failures > threshold
- DB latency high
- Pod crash loop
- High memory usage
- S3 failures

---

## 11. Plane Work Items

```txt
observability: add standard health and readiness endpoints
observability: add Prometheus metrics endpoint
observability: add Grafana dashboard
observability: add OpenTelemetry tracing
observability: add Jaeger local setup
observability: propagate trace context through RabbitMQ
observability: define logging standards
```
