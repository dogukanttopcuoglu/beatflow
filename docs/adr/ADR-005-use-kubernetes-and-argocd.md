# ADR-005 — Use Kubernetes and ArgoCD

## Status

Accepted

## Context

BeatFlow is a microservice system and should practice production-like deployment.

## Decision

Use Kubernetes for orchestration and ArgoCD for GitOps deployment learning.

## Consequences

Positive:

- Learn deployments, services, probes, scaling
- Learn GitOps workflow
- Learn rollback and desired state management

Negative:

- More manifests
- More local setup
- More debugging complexity
