# BeatFlow — Event-Driven Architecture

## 1. Purpose

This document defines BeatFlow's event-driven architecture.

BeatFlow uses asynchronous communication for workflows that do not need to block the user request.

---

## 2. Message Broker Strategy

### 2.1 RabbitMQ

RabbitMQ is the primary message broker for MVP.

Use cases:

- payment completed handling
- order paid processing
- notification delivery
- search indexing
- analytics tracking
- download event processing

RabbitMQ concepts to learn:

- exchange
- queue
- routing key
- producer
- consumer
- ack/nack
- retry
- dead-letter queue
- prefetch
- idempotent consumer

### 2.2 Kafka

Kafka is not MVP primary broker but will be documented and optionally added later.

Kafka is better suited for:

- high-volume event streams
- analytics pipelines
- replayable event history
- consumer-managed offsets
- multi-consumer stream processing

Comparison:

```txt
RabbitMQ:
- broker tracks message delivery
- good for task queues and routing
- easier for command-like async workflows

Kafka:
- consumer tracks offsets
- good for event streams
- strong replay capability
- better for large-scale analytics streams
```

---

## 3. Event Naming Convention

Format:

```txt
<domain>.<event>
```

Examples:

- user.registered
- producer.created
- beat.created
- beat.updated
- beat.submitted_for_review
- beat.published
- beat.rejected
- beat.archived
- license.created
- license.updated
- order.created
- order.paid
- order.failed
- payment.created
- payment.completed
- payment.failed
- download.requested
- download.completed
- beat.viewed
- beat.played
- message.sent
- notification.created

---

## 4. Event Envelope

All events should use a common envelope.

```json
{
  "event_id": "uuid",
  "event_type": "beat.published",
  "occurred_at": "2026-05-29T10:00:00Z",
  "producer": "beat-service",
  "correlation_id": "uuid",
  "traceparent": "00-...",
  "payload": {}
}
```

Fields:

- event_id: unique event id
- event_type: event name
- occurred_at: event creation time
- producer: source service
- correlation_id: request/workflow correlation
- traceparent: OpenTelemetry context
- payload: event-specific data

---

## 5. Core Event Flows

### 5.1 Beat Published Flow

```txt
Beat Service
  → beat.published
  → Marketplace Service indexes beat in Elasticsearch
  → Notification Service notifies producer
  → Analytics Service records publish event
```

### 5.2 Payment Completed Flow

```txt
Payment Service
  → payment.completed
  → Order Service marks order paid
  → Order Service publishes order.paid
  → Notification Service notifies buyer and producer
  → Analytics Service records sale
  → File Service can allow download entitlement
```

### 5.3 Message Sent Flow

```txt
Messaging Service
  → message.sent
  → Notification Service creates in-app notification
  → Email worker may send email
```

### 5.4 Beat Played Flow

```txt
Marketplace Service
  → beat.played
  → Analytics Service increments counters
  → Trending worker updates trending score
```

---

## 6. Retry Strategy

Consumer retry rules:

- Retry transient failures.
- Do not retry invalid payload forever.
- Use retry count header.
- Send poison messages to DLQ.
- Log failure with event_id and consumer name.

---

## 7. DLQ Strategy

Each critical queue should have a DLQ.

Example:

```txt
beat.published.queue
beat.published.dlq
payment.completed.queue
payment.completed.dlq
notification.email.queue
notification.email.dlq
```

DLQ review should be part of operations runbook.

---

## 8. Idempotent Consumers

Every consumer must be idempotent.

Approaches:

- processed_events table
- unique constraints
- idempotent upsert
- Redis keys with TTL for temporary operations

Example:

```txt
Marketplace Service receives beat.published twice.
It upserts the Elasticsearch document with same beat_id.
No duplicate side effect occurs.
```

---

## 9. Transactional Outbox

Critical services should publish events through outbox.

Recommended services:

- Beat Service
- Order Service
- Payment Service
- Messaging Service

Outbox fields:

```txt
id
event_id
event_type
payload
headers
status
created_at
published_at
attempt_count
last_error
```

---

## 10. Event Catalog

### user.registered

Producer:

- Auth Service

Consumers:

- Notification Service
- Analytics Service

### beat.published

Producer:

- Beat Service

Consumers:

- Marketplace Service
- Notification Service
- Analytics Service

### payment.completed

Producer:

- Payment Service

Consumers:

- Order Service
- Notification Service
- Analytics Service

### order.paid

Producer:

- Order Service

Consumers:

- Notification Service
- Analytics Service
- File Service

### beat.played

Producer:

- Marketplace Service

Consumers:

- Analytics Service

### message.sent

Producer:

- Messaging Service

Consumers:

- Notification Service

---

## 11. Async Trace Propagation

Async workflows should include tracing headers.

Headers:

- correlation_id
- traceparent
- tracestate

Goal:

- Follow request from API to RabbitMQ consumer to Elasticsearch/S3/DB calls in Jaeger.

---

## 12. Future Kafka Track

Kafka can be added later for:

- analytics stream
- beat play stream
- audit log stream
- recommendation pipeline

RabbitMQ remains primary MVP broker.
