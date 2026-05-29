# BeatFlow Documentation Pack — Updated

beatflow/
├── apps/
│   └── auth-service/
│       ├── cmd/
│       │   └── api/
│       │       └── main.go
│       ├── internal/
│       │   ├── config/
│       │   ├── handler/
│       │   ├── service/
│       │   └── repository/
│       └── go.mod
│
├── packages/
│   └── go/
│       ├── config/
│       │   ├── config.go
│       │   └── go.mod
│       └── logger/
│           ├── logger.go
│           └── go.mod
│
├── deployments/
│   ├── docker-compose/
│   │   └── docker-compose.yml
│   └── k8s/
│
├── docs/
│
├── scripts/
│
├── Makefile
├── go.work
├── .gitignore
└── README.md
  

This package contains the updated Markdown documentation for BeatFlow.

  

The scope now includes the full modern backend / production-ready microservice learning track:

  

- Go + Fiber REST APIs

- Microservices

- PostgreSQL + pgx + Goose

- Redis

- RabbitMQ

- Kafka comparison / optional event streaming track

- Couchbase / NoSQL experiment track

- Elasticsearch

- AWS S3

- Docker

- Docker Compose

- Kubernetes

- ArgoCD / GitOps

- CI/CD

- Prometheus

- Grafana

- OpenTelemetry

- Jaeger

- Retry / timeout / circuit breaker / rate limiting

- Context propagation

- Graceful shutdown

- Static analysis

- Security scanning

- Container hardening

- Testing

- Load testing

- Architecture Decision Records

  

## Files

  

### Product

  

- `docs/product/project-vision.md`

- `docs/product/product-requirements.md`

  

### Process

  

- `docs/process/development-methodology.md`

- `docs/process/definition-of-done.md`

- `docs/process/sprint-1-checklist.md`

- `docs/process/backlog-updated.md`

  

### Architecture

  

- `docs/architecture/microservice-architecture.md`

- `docs/architecture/system-requirements.md`

- `docs/architecture/non-functional-requirements.md`

- `docs/architecture/resilience-patterns.md`

- `docs/architecture/event-driven-architecture.md`

- `docs/architecture/cqrs-and-consistency.md`

- `docs/architecture/data-storage-strategy.md`

  

### Operations

  

- `docs/observability/observability-strategy.md`

- `docs/deployment/devops-ci-cd-gitops.md`

- `docs/deployment/kubernetes-strategy.md`

- `docs/security/security-strategy.md`

- `docs/testing/testing-strategy.md`

  

### ADR

  

- `docs/adr/ADR-001-use-microservices.md`

- `docs/adr/ADR-002-use-postgresql-redis-elasticsearch-s3.md`

- `docs/adr/ADR-003-use-rabbitmq-and-document-kafka.md`

- `docs/adr/ADR-004-add-couchbase-as-nosql-learning-track.md`

- `docs/adr/ADR-005-use-kubernetes-and-argocd.md`