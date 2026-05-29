# BeatFlow — Product Requirements

## 1. Purpose

This document defines the product requirements for BeatFlow.

BeatFlow is a production-grade, microservice-based beat marketplace backend where producers can upload beats, define license packages, and sell them to buyers.

This document includes product behavior, user roles, business rules, MVP scope, V2 scope, and production-readiness requirements.

---

## 2. Product Summary

BeatFlow provides the backend for a beat marketplace platform.

The platform supports:

- User registration and login
- Producer profile creation
- Beat upload and management
- License package creation
- Marketplace search and filtering
- Cart and order flow
- Payment confirmation through webhook
- Secure download delivery
- Messaging
- Notifications
- Analytics
- Admin moderation
- Production-ready service behavior

---

## 3. User Roles

### 3.1 Guest

Guests can:

- Browse published beats
- Search beats
- Filter beats
- View beat details
- View producer profiles
- Listen to preview audio
- Register
- Login

Guests cannot:

- Add items to cart
- Create orders
- Purchase licenses
- Download private files
- Send messages
- Upload beats

### 3.2 Buyer

Buyers can:

- Browse and search beats
- Listen to preview audio
- Add license packages to cart
- Create orders
- Start payment
- Download purchased files
- View purchase history
- View license details
- Message producers

### 3.3 Producer

Producers can:

- Do everything buyers can do
- Create a producer profile
- Update producer profile
- Create beats
- Update beats
- Archive beats
- Upload beat file metadata
- Create license packages
- Submit beats for review
- View analytics
- View sales and revenue basics
- Reply to buyer messages

### 3.4 Admin

Admins can:

- View pending beat review queue
- Approve beats
- Reject beats with a reason
- Suspend beats
- Inspect users
- Inspect orders
- Inspect payments
- View admin action logs
- View platform analytics

### 3.5 Super Admin

Super admins can:

- Manage admin users
- Manage platform-wide settings
- Manage commission settings
- Manage moderation rules
- Access global operational views

---

## 4. Functional Requirements

### 4.1 Authentication

- Register
- Login
- Refresh token
- Logout
- Password hashing
- Session tracking
- Token revocation
- JWT validation

### 4.2 Producer Profile

- Create producer profile
- Update producer profile
- View public producer profile
- Store avatar/banner metadata
- Store social links
- Support future producer verification

### 4.3 Beat Management

- Create beat draft
- Update beat metadata
- Archive beat
- Submit beat for review
- Admin approve beat
- Admin reject beat
- Admin suspend beat
- Track beat status history

Beat statuses:

- draft
- pending_review
- published
- rejected
- suspended
- archived
- deleted

### 4.4 License Management

License types:

- basic
- premium
- unlimited
- exclusive

Requirements:

- Producers can define license packages.
- Buyers can purchase active licenses.
- Exclusive license can be sold only once.
- Exclusive sale closes future purchases for that beat.
- Existing non-exclusive buyers keep their rights.

### 4.5 Marketplace and Search

- Public beat listing
- Full-text search
- Genre filter
- BPM range filter
- Key filter
- Mood filter
- Price range filter
- Producer filter
- Sorting by newest, popular, trending, best-selling, price
- Elasticsearch-backed search read model
- Event-based index updates

### 4.6 Cart

- Add license to cart
- Remove license from cart
- View cart
- Revalidate cart at checkout

### 4.7 Order

- Create order from cart
- Store price snapshot
- Track order status
- Support idempotency key
- Create purchase entitlements after payment

Order statuses:

- pending
- paid
- failed
- cancelled
- refunded
- partially_refunded

### 4.8 Payment

- Mock payment provider for MVP
- Payment session creation
- Payment webhook
- Webhook signature verification
- Webhook idempotency
- Payment completed event
- Payment failed event

Critical rule:

- Frontend success callback must not mark order as paid.
- Only verified payment webhook can mark payment completed.

### 4.9 File Delivery

- Presigned S3 upload URL
- Presigned S3 download URL
- Private file access
- Download authorization
- Download limits
- Download logs
- File metadata validation

### 4.10 Messaging

- Create conversation
- Send message
- List messages
- Track read status
- Publish message.sent event

### 4.11 Notifications

- In-app notification records
- Email delivery mock
- Async notification workers
- Retry failed notifications

### 4.12 Analytics

- Track beat views
- Track beat plays
- Track sales
- Producer dashboard basics
- Trending score calculation

### 4.13 Admin

- Beat review queue
- Approve beat
- Reject beat with reason
- Suspend beat
- Inspect user/order/payment
- Admin action log

---

## 5. Production-Ready Requirements

Every service should eventually support:

- Config management
- Structured logging
- Request ID / correlation ID
- Graceful shutdown
- Health check endpoint
- Readiness endpoint
- Metrics endpoint
- Timeout management
- Retry strategy
- Circuit breaker for critical external calls
- Rate limiter where needed
- Context propagation
- Trace propagation
- Static analysis
- Security scanning
- Docker containerization
- Kubernetes deployment
- CI/CD pipeline
- GitOps deployment with ArgoCD

---

## 6. Non-Functional Requirements

### 6.1 Performance

- Search should generally respond under 300ms for typical queries.
- Auth endpoints should generally respond under 200ms.
- API handlers should enforce context timeouts.
- Expensive operations should be async where possible.

### 6.2 Reliability

- Payment webhook must be idempotent.
- Event consumers must be idempotent.
- Background jobs must support retries.
- Dead-letter queues must be used for failed messages.
- Services must shut down gracefully.

### 6.3 Scalability

- Stateless services should scale horizontally.
- Kubernetes HPA can scale services based on CPU or custom metrics.
- Search, marketplace, and file services are expected to need scaling earlier than admin service.

### 6.4 Security

- Passwords must be hashed.
- JWT secrets must not be committed.
- S3 private files must not be public.
- Webhook signatures must be verified.
- Rate limiting must protect sensitive endpoints.
- Containers should run with least privilege.

### 6.5 Observability

- Logs must be structured.
- Metrics must be exposed.
- Tracing must propagate across HTTP and async boundaries.
- Dashboards should show latency, error rate, throughput, and queue lag.

---

## 7. Data Requirements

### 7.1 PostgreSQL

Used for transactional source-of-truth data.

Examples:

- users
- producers
- beats
- licenses
- orders
- payments
- downloads
- messages
- admin actions

### 7.2 Redis

Used for:

- cache
- sessions
- refresh token tracking
- idempotency keys
- rate limiting
- counters
- temporary reservations

### 7.3 Elasticsearch

Used for:

- beat search
- filtering
- sorting
- search read model
- autocomplete
- aggregations

### 7.4 Couchbase

Couchbase will be included as a NoSQL learning track.

Potential use cases:

- producer dashboard preferences
- buyer personalization profile
- recommendation state
- flexible metadata documents
- NoSQL CRUD experiment service

Couchbase is not the primary transactional source of truth in MVP.

### 7.5 AWS S3

Used for:

- cover images
- preview audio
- master WAV files
- stems ZIP files
- license contracts
- invoices

---

## 8. Messaging Requirements

### 8.1 RabbitMQ

Used for operational async workflows:

- notifications
- search indexing
- analytics events
- payment completed handling
- download events

### 8.2 Kafka

Kafka will be documented and optionally added for high-volume event streaming.

Potential future use cases:

- analytics event stream
- beat play stream
- audit log stream
- recommendation event pipeline

### 8.3 Message Broker Concepts

The project should cover:

- exchange
- queue
- routing key
- producer
- consumer
- ack/nack
- retry
- dead-letter queue
- idempotent consumer
- consumer group / offset comparison for Kafka

---

## 9. Business Rules Summary

1. Only producers can create beats.
2. Only beat owner can update beat metadata.
3. Non-published beats are not visible in marketplace.
4. Beat publishing requires admin approval.
5. A beat must have at least one active license before review.
6. Buyers can only purchase active licenses on published beats.
7. Producers cannot buy their own beats in MVP.
8. Exclusive license can be sold only once.
9. Exclusive sale closes future license purchases.
10. Order can be marked paid only by verified payment webhook.
11. Order item price snapshot must be stored.
12. Private files must never be publicly accessible.
13. Download requires auth, ownership, paid order, and valid license file access.
14. Elasticsearch is not source of truth.
15. Redis is not source of truth.
16. RabbitMQ handles async operational workflows.
17. Kafka may be used for stream-oriented analytics later.
18. Admin actions must be audited.
19. Deletion should usually be archive/soft delete.
20. Every production service must support health/readiness, logging, timeouts, and graceful shutdown.

---

## 10. Updated MVP Feature List

- Auth
- Producer profiles
- Beat CRUD
- Beat file metadata
- License CRUD
- Admin beat review
- Marketplace search
- Cart
- Order
- Mock payment
- Payment webhook
- Download authorization
- Messaging basics
- Analytics basics
- Notifications basics
- Admin actions
- Docker Compose
- PostgreSQL
- Redis
- RabbitMQ
- Elasticsearch
- AWS S3 integration
- Couchbase experiment documentation
- Zap logging
- Viper config
- Goose migrations
- pgx database access
- Health/readiness endpoints
- Graceful shutdown
- Timeout strategy
- Retry strategy
- Basic metrics
- Tests
- CI basics

---

## 11. Updated V2 Feature List

- Real Stripe integration
- iyzico support
- Payout system
- Refund workflow
- Wishlist
- Reviews and ratings
- WebSocket messaging
- Custom beat requests
- Advanced recommendation
- Producer verification
- Subscription plans
- Affiliate/referral system
- Fraud detection
- Audio fingerprinting
- Advanced analytics
- Multi-currency
- Kubernetes production deployment
- ArgoCD GitOps deployment
- Prometheus/Grafana dashboards
- OpenTelemetry + Jaeger tracing
- Kafka analytics stream
- Couchbase NoSQL module
- Temporal workflow engine experiment
- Service mesh experiment
- Feature flags
- OPA/Casbin authorization
