# BeatFlow — Development Methodology

## 1. Purpose

This document defines how BeatFlow will be planned, developed, tested, documented, deployed, and improved.

BeatFlow is a large learning-oriented backend project. It includes microservices, databases, caching, messaging, search, storage, DevOps, observability, security, testing, and deployment.

---

## 2. Methodology Summary

BeatFlow uses an:

**Agile-inspired, Kanban-driven, documentation-guided, DevOps-oriented development process.**

Meaning:

- Work is split into small iterations.
- Work items are tracked in Plane.
- Documentation is written before or alongside implementation.
- Every service is treated as production-bound.
- Automation is preferred over manual work.

---

## 3. Workflow

Plane states:

- Backlog
- Todo
- In Progress
- Done
- Cancelled

Rules:

- Keep maximum 1-2 items in progress.
- Every item needs acceptance criteria.
- Done means Definition of Done is satisfied.
- Labels are useful but Module + Cycle + State are the core tracking fields.

---

## 4. Cycles / Sprints

Recommended sprint length:

```txt
1 week
```

Each sprint has:

- sprint goal
- selected work items
- completion criteria
- review
- retrospective notes

---

## 5. Work Item Format

```txt
Title:
architecture: design circuit breaker strategy

Description:
Document when and how BeatFlow services should use circuit breakers for downstream dependencies.

Acceptance Criteria:
- Defines circuit breaker purpose
- Defines when to use it
- Defines when not to use it
- Includes Go package option
- Includes config fields
- Includes observability requirements
- Document is saved under docs/architecture/resilience-patterns.md

Module:
Documentation or Observability

Labels:
type: docs
area: architecture

Priority:
High

State:
Todo

Cycle:
Sprint X
```

---

## 6. Branch Naming

Format:

```txt
<type>/<scope>-<short-description>
```

Examples:

```txt
docs/project-vision
infra/init-monorepo
feat/auth-service-skeleton
test/payment-webhook-idempotency
security/container-hardening
observability/prometheus-metrics
```

---

## 7. Commit Convention

Use Conventional Commits.

Examples:

```txt
docs(product): update project vision
docs(architecture): add resilience patterns
feat(auth): add health check endpoint
infra(k8s): add auth service deployment
ci(actions): add lint and test workflow
security(container): run app as non-root
```

---

## 8. DevOps Practice

BeatFlow treats DevOps as automation.

Automate:

- formatting
- linting
- static analysis
- tests
- builds
- Docker image creation
- security scans
- deployment
- rollback
- observability setup

---

## 9. Documentation-Driven Development

Major decisions require docs.

Docs include:

- product requirements
- system requirements
- architecture decisions
- service boundaries
- event catalog
- resilience patterns
- security strategy
- observability strategy
- deployment strategy

---

## 10. ADRs

Architecture Decision Records are required for major decisions.

Examples:

- Use microservices
- Use PostgreSQL as source of truth
- Use RabbitMQ for async workflows
- Add Couchbase as NoSQL learning track
- Use Kubernetes and ArgoCD

---

## 11. Quality Gates

Before moving work to Done:

- acceptance criteria complete
- tests pass or not required
- lint passes where applicable
- docs updated
- no secrets committed
- local run verified where applicable
- CI green where applicable

---

## 12. Learning Tracks

BeatFlow has core and advanced learning tracks.

Core:

- Go
- Fiber
- PostgreSQL
- Redis
- RabbitMQ
- Elasticsearch
- S3
- Docker
- CI
- testing

Advanced:

- Kubernetes
- ArgoCD
- OpenTelemetry
- Jaeger
- Prometheus/Grafana
- Couchbase
- Kafka
- Circuit breaker
- Load testing
- Security scanning
- GitOps
- Service mesh later

---

## 13. Release Strategy

Milestone versions:

```txt
v0.1.0 documentation and foundation
v0.2.0 auth service
v0.3.0 user and producer service
v0.4.0 beat service
v0.5.0 file service
v0.6.0 marketplace search
v0.7.0 order and payment flow
v0.8.0 event-driven workers
v0.9.0 observability and CI
v1.0.0 MVP complete
v1.1.0 Kubernetes and ArgoCD
v1.2.0 NoSQL/Kafka experiments
```

---

## 14. Summary

BeatFlow will be developed slowly, intentionally, and with production-grade discipline.

The goal is to learn modern backend engineering by implementing real concepts in a realistic system.
