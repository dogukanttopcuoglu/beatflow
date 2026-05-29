# BeatFlow — CQRS, Consistency, and CAP Notes

## 1. Purpose

This document explains CQRS, read/write separation, consistency, availability, and how these concepts apply to BeatFlow.

---

## 2. CQRS

CQRS means Command Query Responsibility Segregation.

It separates:

- write model
- read model

In BeatFlow:

```txt
Write model: PostgreSQL
Read/search model: Elasticsearch
Cache/read optimization: Redis
```

Example:

1. Beat Service writes beat to PostgreSQL.
2. Beat Service publishes beat.published event.
3. Marketplace Service consumes event.
4. Marketplace Service indexes beat into Elasticsearch.
5. Users search through Elasticsearch.

---

## 3. Why Separate Read and Write?

Benefits:

- Write model can stay normalized.
- Search model can be denormalized.
- Reads can be optimized for marketplace queries.
- Search can scale independently.
- Analytics projections can be built asynchronously.

Tradeoff:

- Read model can be stale.
- Eventual consistency must be accepted.

---

## 4. Eventual Consistency

Some data does not need to be immediately consistent.

Eventually consistent domains:

- search results
- analytics
- notifications
- trending score
- recommendation state

Strongly consistent domains:

- payment
- order paid status
- purchase entitlement
- exclusive license sale
- download authorization

---

## 5. CAP Perspective

Distributed systems must reason about:

- consistency
- availability
- partition tolerance

In practice:

- Some domains prioritize correctness.
- Some domains prioritize availability.

Examples:

### Payment

Prefer consistency.

If payment/order status is uncertain, do not incorrectly grant access.

### Search

Prefer availability/eventual consistency.

If a beat appears in search a few seconds late, that is acceptable.

### Analytics

Prefer availability/eventual consistency.

A play count can update later.

### Exclusive License

Prefer consistency.

Two buyers must not both purchase the same exclusive license.

---

## 6. Read Model Sync Problem

Problem:

- Beat is written to PostgreSQL.
- Elasticsearch is not updated yet.
- User searches and does not see beat immediately.

Solutions:

- Accept eventual consistency.
- Use outbox pattern for reliable event publishing.
- Monitor indexing lag.
- Add reindex job.
- Add admin reindex command.

---

## 7. Database Replication vs Application Projection

Read/write separation can happen at different layers.

### Database-level replication

PostgreSQL primary and read replicas.

Pros:

- simpler for application
- database handles replication

Cons:

- still same relational shape
- not optimized for search ranking

### Application-level projection

PostgreSQL → event → Elasticsearch.

Pros:

- custom search document
- optimized for search
- independent scaling

Cons:

- eventual consistency
- reindexing needed
- event handling complexity

BeatFlow uses application-level projection for search.

---

## 8. BeatFlow Consistency Decisions

| Domain | Consistency Model |
|---|---|
| Auth | Strong |
| Payment | Strong |
| Order | Strong |
| Download Authorization | Strong |
| Exclusive License Sale | Strong |
| Search | Eventual |
| Analytics | Eventual |
| Notifications | Eventual |
| Trending | Eventual |
| Recommendations | Eventual |
