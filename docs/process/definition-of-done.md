# BeatFlow — Definition of Done

## 1. Purpose

This document defines what "done" means for BeatFlow.

Done means the task satisfies its acceptance criteria and meets project quality standards.

---

## 2. General Done Criteria

A work item is Done when:

- Acceptance criteria are completed.
- Code or documentation is committed.
- Work is in the correct path.
- Formatting is clean.
- Tests are added or explicitly not required.
- No known blocker remains.
- No secrets are committed.
- Plane item is moved to Done.

---

## 3. Documentation Task Done Criteria

- Document exists in expected path.
- Purpose is clear.
- Sections are organized.
- Markdown formatting is clean.
- The document can be understood without external context.
- Related README links are updated when applicable.

---

## 4. Code Task Done Criteria

- Code compiles.
- Code is formatted with gofmt/goimports.
- Error handling exists.
- Logs are useful and safe.
- Config values are not hardcoded.
- No secrets are committed.
- Tests are added when applicable.
- Acceptance criteria are satisfied.

---

## 5. API Task Done Criteria

- Route is registered.
- Request parsing is safe.
- Validation exists.
- Auth is enforced if needed.
- Authorization is enforced if needed.
- Correct status codes are returned.
- Error shape is consistent.
- Handler delegates business logic to service layer.
- API docs are updated when available.

---

## 6. Database Task Done Criteria

- Goose migration exists.
- Up and Down sections exist.
- Constraints are defined.
- Indexes are added for query patterns.
- Transaction boundaries are clear.
- Rollback is tested when feasible.

---

## 7. Microservice Task Done Criteria

- Service has clear folder structure.
- Service starts locally.
- Config is loaded with Viper/shared config.
- Logger uses Zap/shared logger.
- `/healthz` exists.
- `/readyz` exists.
- Graceful shutdown is implemented or planned.
- Service does not directly access another service's database.

---

## 8. Resilience Done Criteria

When a task involves outbound calls or async processing:

- Timeouts are defined.
- Retry strategy is considered.
- Circuit breaker is considered for critical dependencies.
- Operation is idempotent if retried.
- Context is propagated.
- Errors are observable.
- DLQ exists for critical async consumers.

---

## 9. Observability Done Criteria

For service tasks:

- Structured logs exist.
- Request/correlation ID is included when available.
- Metrics are added when meaningful.
- Trace propagation is considered.
- Health/readiness endpoints are updated if dependencies changed.

---

## 10. Security Done Criteria

- Authentication is required where needed.
- Authorization and ownership checks exist.
- Inputs are validated.
- Secrets are not committed.
- Sensitive data is not logged.
- Webhook signatures are verified when applicable.
- Private files are not public.
- Container does not include unnecessary tools when production image exists.

---

## 11. Testing Done Criteria

Tests are required for critical logic:

- auth
- RBAC
- payment webhook
- exclusive license
- order creation
- download authorization
- rate limiter
- event consumers
- retry/circuit breaker logic

If tests are not added, the reason must be clear.

---

## 12. Docker/Kubernetes Done Criteria

For infrastructure tasks:

- Dockerfile builds.
- Docker Compose works locally.
- Kubernetes manifest is valid.
- ConfigMap/Secret strategy is clear.
- Probes are configured.
- Resource requests/limits exist.
- Container runs as non-root when possible.

---

## 13. CI/CD Done Criteria

- Pipeline runs.
- Tests run.
- Lint runs.
- Security scan runs when configured.
- Docker build works when configured.
- Failure output is understandable.

---

## 14. Review Checklist

Before Done:

- [ ] Acceptance criteria complete
- [ ] Code/docs in expected path
- [ ] Formatting clean
- [ ] Tests pass or not required
- [ ] No secrets committed
- [ ] Error handling acceptable
- [ ] Logs safe
- [ ] Docs updated
- [ ] Docker/local run checked if applicable
- [ ] CI green if applicable
- [ ] Plane item updated
