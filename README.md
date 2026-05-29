# BeatFlow

BeatFlow is a production-grade, microservice-based Go backend system for a beat marketplace platform.

The project is designed as a long-term backend engineering portfolio project. It focuses on Go, Fiber, microservices, PostgreSQL, Redis, RabbitMQ, Elasticsearch, AWS S3, Docker, Kubernetes, CI/CD, observability, security, testing, and resilience patterns.

## Repository Structure

```txt
beatflow/
├── apps/
│   └── auth-service/
├── packages/
│   └── go/
│       ├── config/
│       └── logger/
├── docs/
├── deployments/
├── scripts/
├── README.md
└── .gitignore
```

## Directory Overview

### `apps/`

Contains deployable backend applications and microservices.

In this repository, `apps/` acts as the services directory.

Initial service:

* `auth-service`: handles authentication and identity-related features.

Future services may include:

* `user-service`
* `beat-service`
* `file-service`
* `marketplace-service`
* `order-service`
* `payment-service`
* `notification-service`
* `analytics-service`
* `messaging-service`
* `admin-service`

### `packages/go/`

Contains shared Go packages used by multiple services.

In this repository, `packages/go/` acts as the shared Go package directory.

Initial shared packages:

* `config`: shared configuration loading utilities.
* `logger`: shared structured logging utilities.

Shared packages should only contain technical concerns such as configuration, logging, middleware, errors, validation helpers, and observability helpers.

Domain-specific business logic should stay inside the owning service.

### `docs/`

Contains project documentation, including product, architecture, process, security, testing, observability, deployment, and ADR documents.

### `deployments/`

Contains deployment-related files such as Docker Compose, Kubernetes, and GitOps manifests.

### `scripts/`

Contains helper scripts for local development, testing, automation, and operations.

## Documentation

Project documentation lives under `docs/`.

Recommended starting points:

* `docs/product/project-vision.md`
* `docs/product/product-requirements.md`
* `docs/process/development-methodology.md`
* `docs/process/definition-of-done.md`
* `docs/process/sprint-1-checklist.md`
* `docs/architecture/microservice-architecture.md`
* `docs/architecture/system-requirements.md`
* `docs/architecture/non-functional-requirements.md`
* `docs/architecture/resilience-patterns.md`
* `docs/architecture/event-driven-architecture.md`
* `docs/architecture/cqrs-and-consistency.md`
* `docs/architecture/data-storage-strategy.md`
* `docs/observability/observability-strategy.md`
* `docs/deployment/devops-ci-cd-gitops.md`
* `docs/deployment/kubernetes-strategy.md`
* `docs/security/security-strategy.md`
* `docs/testing/testing-strategy.md`
* `docs/adr/ADR-001-use-microservices.md`
* `docs/adr/ADR-002-use-postgresql-redis-elasticsearch-s3.md`
* `docs/adr/ADR-003-use-rabbitmq-and-document-kafka.md`
* `docs/adr/ADR-004-add-couchbase-as-nosql-learning-track.md`
* `docs/adr/ADR-005-use-kubernetes-and-argocd.md`
