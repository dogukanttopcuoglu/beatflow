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

Contains project documentation.

Examples:

* product documentation
* architecture documentation
* development process documentation
* ADRs
* security strategy
* testing strategy
* observability strategy
* deployment strategy

### `deployments/`

Contains deployment-related files.

Examples:

* Docker Compose files
* Kubernetes manifests
* ArgoCD GitOps manifests
* environment-specific deployment configuration

### `scripts/`

Contains helper scripts for local development, testing, automation, and operations.

Examples:

* local setup scripts
* migration scripts
* test scripts
* formatting/linting scripts
* CI helper scripts

## Current Sprint

Sprint 1 focuses on creating the initial project foundation before implementing business features.

Initial foundation goals:

* initialize Git repository
* create root README
* add `.gitignore`
* create base monorepo folder structure
* create application/service directory
* create shared package directory
* create documentation directory
* create deployment directory
* create scripts directory
* document repository structure

```
```
