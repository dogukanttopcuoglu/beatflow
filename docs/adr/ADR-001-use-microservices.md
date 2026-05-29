# ADR-001 — Use Microservices

## Status

Accepted

## Context

BeatFlow is a learning project intended to practice modern backend engineering.

A modular monolith would be simpler, but it would not expose enough distributed system concepts.

## Decision

Use a microservice-based architecture in a monorepo.

## Consequences

Positive:

- Learn service boundaries
- Learn event-driven architecture
- Learn database ownership
- Learn Kubernetes and CI/CD
- Learn observability

Negative:

- More complexity
- More infrastructure
- More testing burden
- More debugging difficulty
