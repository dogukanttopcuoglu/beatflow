# ADR-002 — Use PostgreSQL, Redis, Elasticsearch, and S3

## Status

Accepted

## Context

BeatFlow needs transactional data, cache/temp data, search, and file storage.

## Decision

Use:

- PostgreSQL for transactional source-of-truth data
- Redis for cache, sessions, counters, idempotency, rate limiting
- Elasticsearch for search projection
- AWS S3 for files

## Consequences

Positive:

- Clear storage responsibility
- Realistic backend stack
- Strong learning value

Negative:

- More local dependencies
- More operational complexity
