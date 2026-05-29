# ADR-003 — Use RabbitMQ and Document Kafka

## Status

Accepted

## Context

BeatFlow needs asynchronous processing for notifications, search indexing, analytics, and payment/order events.

## Decision

Use RabbitMQ for MVP async workflows.

Document Kafka and optionally add it later for analytics/event-streaming.

## Consequences

Positive:

- RabbitMQ is simpler for task/event workflows.
- Kafka comparison teaches stream processing concepts.

Negative:

- Two broker concepts may confuse scope if added too early.
