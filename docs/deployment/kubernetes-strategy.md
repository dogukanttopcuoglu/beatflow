# BeatFlow — Kubernetes Strategy

## 1. Purpose

Kubernetes will be used as the production-like orchestration layer for BeatFlow.

Docker runs containers. Kubernetes manages containers at scale.

---

## 2. Why Kubernetes?

BeatFlow has multiple services.

Kubernetes helps with:

- service discovery
- scaling
- self-healing
- rolling updates
- rollbacks
- config management
- secret management
- health checks
- resource limits
- scheduling
- ingress routing

---

## 3. Local vs Production-like

Local development:

```txt
Docker Compose
```

Production-like deployment:

```txt
Kubernetes
```

Local Kubernetes options:

- kind
- minikube

Recommended:

```txt
kind
```

---

## 4. Kubernetes Resources

BeatFlow will use:

- Namespace
- Deployment
- Service
- ConfigMap
- Secret
- Ingress
- Liveness probe
- Readiness probe
- Resource requests
- Resource limits
- Horizontal Pod Autoscaler

---

## 5. Deployment

Each service should have a Deployment.

Example services:

- auth-service
- user-service
- beat-service
- marketplace-service
- order-service
- payment-service
- file-service
- notification-service
- analytics-service
- messaging-service
- admin-service

---

## 6. Service Discovery

Kubernetes Services provide internal DNS.

Example:

```txt
http://beat-service:8080
http://order-service:8080
```

Services should not use pod IPs directly.

---

## 7. Liveness and Readiness

### Liveness Probe

Checks if process is alive.

Endpoint:

```txt
GET /healthz
```

### Readiness Probe

Checks if service can receive traffic.

Endpoint:

```txt
GET /readyz
```

Readiness can check critical dependencies.

---

## 8. Resource Management

Every service should define:

```yaml
resources:
  requests:
    cpu: "100m"
    memory: "128Mi"
  limits:
    cpu: "500m"
    memory: "512Mi"
```

Goal:

- avoid one service consuming all node resources
- enable scheduler decisions
- support autoscaling

---

## 9. Autoscaling

Horizontal Pod Autoscaler can scale based on:

- CPU
- memory
- custom metrics

Good candidates:

- marketplace-service
- search/indexing workers
- file-service
- order-service

Less likely to need high scaling:

- admin-service

---

## 10. Rolling Updates

Kubernetes should roll out new versions gradually.

Requirements:

- readiness probe must pass before traffic
- old pods should not terminate before new pods are ready
- graceful shutdown must handle termination

---

## 11. Security

Kubernetes security practices:

- do not run as root
- read-only root filesystem when possible
- least privilege service accounts
- use Secrets for sensitive values
- use ConfigMaps for non-sensitive config
- avoid kubectl exec into production containers unless necessary
- minimal container images

---

## 12. K9s

K9s can be used for local Kubernetes inspection.

Useful for:

- pods
- logs
- services
- deployments
- restarts
- events

---

## 13. ArgoCD

ArgoCD will be used later for GitOps.

Kubernetes manifests will be stored in Git.

ArgoCD will sync cluster state from Git.

---

## 14. Plane Work Items

```txt
kubernetes: create namespace manifest
kubernetes: create auth-service deployment and service
kubernetes: add configmap and secret examples
kubernetes: add liveness and readiness probes
kubernetes: add resource requests and limits
kubernetes: add ingress manifest
kubernetes: add HPA example
kubernetes: document kind local cluster setup
kubernetes: document K9s usage
```
