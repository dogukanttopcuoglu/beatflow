# BeatFlow — Data Storage Strategy

## 1. Purpose

This document defines how BeatFlow uses different storage technologies.

BeatFlow intentionally uses multiple storage systems to learn real backend tradeoffs.

---

## 2. Storage Overview

| Technology | Role |
|---|---|
| PostgreSQL | Transactional source of truth |
| Redis | Cache, sessions, rate limit, counters, idempotency |
| Elasticsearch | Search/read model |
| AWS S3 | Object storage |
| Couchbase | NoSQL learning/experiment track |

---

## 3. PostgreSQL

PostgreSQL is the primary transactional database.

Used for:

- auth identity
- producer profiles
- beat metadata
- licenses
- orders
- payments
- downloads
- messages
- notifications
- admin actions

Learning goals:

- schema design
- transactions
- indexing
- constraints
- row-level locking
- isolation levels
- migrations
- query optimization
- EXPLAIN ANALYZE

Tools:

- pgx
- Goose

---

## 4. Redis

Redis is not the source of truth.

Used for:

- cache
- refresh token/session tracking
- rate limiting
- counters
- idempotency keys
- temporary reservations
- trending cache
- search result cache

Learning goals:

- TTL
- atomic increment
- hashes
- sets
- sorted sets
- Lua basics
- distributed lock concepts
- cache invalidation

---

## 5. Elasticsearch

Elasticsearch is used as a read/search projection.

Used for:

- beat search
- filters
- sorting
- autocomplete
- aggregations
- trending/search ranking

Learning goals:

- index mapping
- document modeling
- analyzers
- query DSL
- fuzzy search
- aggregations
- reindexing
- eventual consistency

Rules:

- Elasticsearch is not source of truth.
- Documents can be rebuilt from PostgreSQL and events.

---

## 6. AWS S3

S3 stores files.

Object types:

- cover images
- preview MP3
- master WAV
- stems ZIP
- license contracts
- invoices

Learning goals:

- bucket design
- object key strategy
- public/private access
- presigned upload URL
- presigned download URL
- IAM permissions
- lifecycle policy
- file validation

---

## 7. Couchbase

Couchbase is added as a NoSQL learning track.

It is not the MVP primary source of truth.

Potential use cases:

- buyer personalization profile
- producer dashboard preferences
- flexible metadata documents
- recommendation state
- NoSQL CRUD experiment
- comparison with PostgreSQL JSONB

Learning goals:

- document database modeling
- key-value access
- N1QL basics
- bucket/scope/collection structure
- TTL
- indexing
- SDK usage in Go
- difference between relational and document modeling

Example experimental module:

```txt
services/personalization-service
```

Possible documents:

```json
{
  "user_id": "uuid",
  "favorite_genres": ["trap", "drill"],
  "preferred_bpm_range": {
    "min": 120,
    "max": 150
  },
  "recent_searches": ["dark trap", "uk drill"],
  "updated_at": "2026-05-29T10:00:00Z"
}
```

---

## 8. Kafka

Kafka is not a storage database but is relevant to event streams.

Potential future use:

- analytics event stream
- audit log stream
- recommendation stream

Kafka is compared with RabbitMQ in event-driven docs.

---

## 9. Storage Decision Rules

Use PostgreSQL when:

- data is transactional
- consistency matters
- relations matter
- financial/audit history matters

Use Redis when:

- data is temporary
- low latency is needed
- TTL is useful
- counters or rate limits are needed

Use Elasticsearch when:

- search/filter/ranking is needed
- denormalized read model is useful

Use S3 when:

- storing files
- files are large
- presigned access is needed

Use Couchbase when:

- flexible JSON documents are useful
- NoSQL modeling is the learning goal
- personalization/recommendation state is document-like

---

## 10. Backup Strategy

Future backup docs should include:

- PostgreSQL backups
- S3 lifecycle and versioning
- Elasticsearch snapshots
- Redis persistence awareness
- Couchbase backup for experimental module
