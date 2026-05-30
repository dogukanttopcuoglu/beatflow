# BeatFlow — Sprint 1 Checklist

## Sprint

**Sprint 1 — Project Foundation**

## Sprint Goal

Create the initial foundation for a production-grade microservice backend project before implementing business features.

Sprint 1 focuses on:

* repository foundation
* documentation foundation
* local development workflow
* shared Go packages
* local infrastructure skeleton
* initial Auth Service skeleton
* standard operational endpoints

The goal is not to implement full authentication business logic yet. The goal is to prepare a clean, production-minded base for future service development.

---

## Definition of Done Reference

Before moving any Sprint 1 item to Done, check the project Definition of Done:

```txt
docs/process/definition-of-done.md
```

A work item is Done when:

* acceptance criteria are complete
* code or documentation is committed
* work is in the expected path
* formatting is clean
* tests pass or are explicitly not required
* no secrets are committed
* local run is verified when applicable
* README or documentation is updated when needed
* Plane item is moved to Done

---

## Documentation

* [ ] `docs: write project vision and scope`

  * Output: `docs/product/project-vision.md`

* [ ] `docs: write product requirements`

  * Output: `docs/product/product-requirements.md`

* [ ] `docs: define development methodology`

  * Output: `docs/process/development-methodology.md`

* [ ] `docs: write Definition of Done`

  * Output: `docs/process/definition-of-done.md`

* [ ] `docs: define microservice architecture decision`

  * Output: `docs/architecture/microservice-architecture.md`

* [ ] `docs: write system requirements`

  * Output: `docs/architecture/system-requirements.md`

* [ ] `docs: write non-functional requirements`

  * Output: `docs/architecture/non-functional-requirements.md`

* [ ] `docs: write Sprint 1 README checklist`

  * Output: `docs/process/sprint-1-checklist.md`

---

## Architecture Additions

* [ ] `architecture: document resilience patterns`

  * Output: `docs/architecture/resilience-patterns.md`

* [ ] `architecture: document event-driven architecture`

  * Output: `docs/architecture/event-driven-architecture.md`

* [ ] `architecture: document CQRS and consistency`

  * Output: `docs/architecture/cqrs-and-consistency.md`

* [ ] `architecture: document data storage strategy`

  * Output: `docs/architecture/data-storage-strategy.md`

---

## Observability / DevOps / Security

* [ ] `observability: write observability strategy`

  * Output: `docs/observability/observability-strategy.md`

* [ ] `deployment: write DevOps CI/CD GitOps strategy`

  * Output: `docs/deployment/devops-ci-cd-gitops.md`

* [ ] `deployment: write Kubernetes strategy`

  * Output: `docs/deployment/kubernetes-strategy.md`

* [ ] `security: write security strategy`

  * Output: `docs/security/security-strategy.md`

* [ ] `testing: write testing strategy`

  * Output: `docs/testing/testing-strategy.md`

---

## ADRs

* [ ] `adr: use microservices`

  * Output: `docs/adr/ADR-001-use-microservices.md`

* [ ] `adr: use PostgreSQL, Redis, Elasticsearch, and S3`

  * Output: `docs/adr/ADR-002-use-postgresql-redis-elasticsearch-s3.md`

* [ ] `adr: use RabbitMQ and document Kafka`

  * Output: `docs/adr/ADR-003-use-rabbitmq-and-document-kafka.md`

* [ ] `adr: add Couchbase as NoSQL learning track`

  * Output: `docs/adr/ADR-004-add-couchbase-as-nosql-learning-track.md`

* [ ] `adr: use Kubernetes and ArgoCD`

  * Output: `docs/adr/ADR-005-use-kubernetes-and-argocd.md`

---

## Project Foundation

* [x] `infra: initialize BeatFlow monorepo`

  * Created initial repository structure.
  * Added root README and `.gitignore`.
  * Established `apps/`, `packages/go/`, `docs/`, `deployments/`, and `scripts/`.

* [x] `infra: setup Go workspace`

  * Added root `go.work`.
  * Connected `apps/auth-service`, `packages/go/config`, and `packages/go/logger`.

* [x] `infra: add base Makefile`

  * Added common development commands:

    * `make help`
    * `make test`
    * `make lint`
    * `make fmt`
    * `make docker-up`
    * `make docker-down`
    * `make migrate-up`
    * `make migrate-down`

* [x] `infra: add Docker Compose skeleton`

  * Added local infrastructure skeleton.
  * Included PostgreSQL, Redis, RabbitMQ, Elasticsearch, and Nginx placeholder.
  * Added `.env.example`.

* [x] `shared: design Viper config package`

  * Added shared config package under `packages/go/config`.
  * Supports defaults, environment variable overrides, env prefixes, and service-specific config structs.

* [x] `shared: design Zap logger package`

  * Added shared logger package under `packages/go/logger`.
  * Supports development config, production JSON config, service name field, and request/correlation field pattern.

* [x] `auth: create auth-service skeleton`

  * Added Auth Service startup flow.
  * Wired shared config package.
  * Wired shared logger package.
  * Added Fiber server.
  * Implemented graceful shutdown.

* [x] `auth: add health check endpoint`

  * Added `GET /healthz`.
  * Added `GET /readyz`.
  * Added health handler structure.
  * Added generic HTTP adapter for typed handlers.
  * Registered health routes in Auth Service app layer.

---

## Updated Learning Scope

Sprint 1 covers the initial foundation for:

* Go
* Fiber
* Microservices
* Monorepo development
* Go workspaces
* Viper configuration
* Zap logging
* Docker Compose
* PostgreSQL
* Redis
* RabbitMQ
* Elasticsearch
* Nginx placeholder
* Health checks
* Readiness checks
* Graceful shutdown
* Documentation-driven development
* Plane-based work tracking
* GitHub-based progress tracking

Sprint 1 documentation also introduces future learning tracks:

* Kubernetes
* ArgoCD
* CI/CD
* RabbitMQ
* Kafka comparison
* Couchbase experiment
* AWS S3
* Prometheus
* Grafana
* OpenTelemetry
* Jaeger
* timeout management
* retry
* circuit breaker
* rate limiter
* context propagation
* static analysis
* security scanning
* load testing

---

## Progress Notes

This checklist is intended to track Sprint 1 progress in GitHub alongside Plane.

Plane remains the task board.

GitHub keeps a versioned project record.

When a work item is completed:

1. Update the checkbox in this document.
2. Commit the code or documentation.
3. Push to GitHub.
4. Move the corresponding Plane item to Done.

---

## Final Sprint 1 Completion Checklist

Before closing Sprint 1:

* [ ] All selected Sprint 1 work items are either Done or intentionally moved out.
* [ ] Root README links to this checklist.
* [ ] All documentation files are committed.
* [ ] All code compiles.
* [ ] `make test` passes.
* [ ] Docker Compose config validates.
* [ ] Auth Service starts locally.
* [ ] `/healthz` returns `200 OK`.
* [ ] `/readyz` returns `200 OK`.
* [ ] No secrets are committed.
* [ ] Plane Sprint 1 board is updated.
