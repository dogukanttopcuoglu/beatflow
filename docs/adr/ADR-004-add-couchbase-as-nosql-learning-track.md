# ADR-004 — Add Couchbase as NoSQL Learning Track

## Status

Accepted

## Context

The project should include NoSQL learning concepts.

Couchbase is useful for document-oriented modeling and key-value access.

## Decision

Add Couchbase as an experimental learning track, not as the primary transactional store.

Potential module:

- Personalization Service
- User preference documents
- Recommendation profile documents

## Consequences

Positive:

- NoSQL hands-on practice
- Document modeling experience
- Comparison with PostgreSQL JSONB

Negative:

- Extra infrastructure
- Extra SDK and testing setup
- More operational complexity
