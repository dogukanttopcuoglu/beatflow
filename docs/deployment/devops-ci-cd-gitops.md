# BeatFlow — DevOps, CI/CD, and GitOps

## 1. Purpose

DevOps is treated as automation and culture.

BeatFlow should teach how code moves from local development to production-like deployment through automated steps.

---

## 2. DevOps Principle

Core principle:

```txt
If a task can be automated, it should not be manual forever.
```

BeatFlow automation areas:

- formatting
- linting
- tests
- static analysis
- build
- Docker image creation
- image scanning
- deployment
- rollback
- Kubernetes sync
- monitoring setup

---

## 3. CI Pipeline

Tool:

- GitHub Actions

Pipeline steps:

1. Checkout
2. Setup Go
3. Cache dependencies
4. Run gofmt check
5. Run go vet
6. Run golangci-lint
7. Run unit tests
8. Run integration tests when needed
9. Run static analysis
10. Run security checks
11. Build binaries
12. Build Docker images
13. Scan Docker images
14. Push images if branch/tag allows

---

## 4. Static Analysis

Tools:

- go vet
- golangci-lint
- gosec
- CodeQL
- SonarQube optional

Static analysis should catch:

- unused code
- unchecked errors
- suspicious constructs
- security smells
- overly complex code
- potential bugs

---

## 5. Security Scanning

Tools:

- gosec
- Trivy
- GitHub Dependabot
- CodeQL

Scan:

- Go source code
- dependencies
- Docker images
- secrets
- container vulnerabilities

---

## 6. CD Pipeline

Initial CD strategy:

- build Docker image
- push to registry
- update Kubernetes manifest image tag manually or through automation

Advanced CD strategy:

- ArgoCD watches Git repository
- Kubernetes manifests are updated in Git
- ArgoCD syncs cluster to desired state

---

## 7. ArgoCD / GitOps

GitOps principle:

```txt
Git is the source of truth for deployment state.
```

Flow:

1. Developer merges code.
2. CI builds and pushes image.
3. Deployment manifest image tag is updated.
4. Git repo contains desired Kubernetes state.
5. ArgoCD detects change.
6. ArgoCD syncs Kubernetes cluster.
7. Rollback happens by reverting Git commit.

Benefits:

- audit trail
- reproducible deployment
- easier rollback
- less manual kubectl usage
- clear desired state

---

## 8. Jenkins / GitHub Actions / GitLab CI

BeatFlow will use GitHub Actions by default.

Other CI/CD tools to understand:

- Jenkins
- GitLab CI
- Travis CI

Learning goal:

- Understand that tools differ but pipeline concepts are similar:
  - test
  - lint
  - build
  - scan
  - package
  - deploy

---

## 9. Docker Image Strategy

Requirements:

- multi-stage builds
- small final image
- no unnecessary tools
- non-root user when possible
- pinned base image versions
- healthcheck where useful
- image scanning

Example:

```txt
builder image → compile Go binary
runtime image → only binary + certs
```

---

## 10. Deployment Environments

Environments:

- local
- development
- staging
- production-like

Local:

- Docker Compose

Production-like:

- Kubernetes
- ArgoCD

---

## 11. Rollback Strategy

Rollback options:

- Kubernetes rollout undo
- ArgoCD sync previous Git commit
- revert manifest change
- redeploy previous image tag

---

## 12. Plane Work Items

```txt
ci-cd: add GitHub Actions test pipeline
ci-cd: add golangci-lint pipeline
ci-cd: add gosec security scan
ci-cd: add Trivy Docker image scan
ci-cd: build Docker images in CI
deployment: document ArgoCD GitOps flow
deployment: create ArgoCD application manifest
deployment: document rollback strategy
```
