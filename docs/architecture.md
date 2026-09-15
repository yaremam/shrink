# Shrink! — Architecture

Living overview of Shrink!'s system structure. Update this whenever a
feature changes a component boundary, adds a data model, or reverses
an earlier architectural decision.

## Components

| Component | Path         | Stack | Status |
|-----------|--------------|-------|--------|
| Client    | `client/`    | Flutter (mobile + web), Firebase Auth | Initialized — placeholder screen implemented and tested; no auth/real screens yet |
| Backend   | `backend/`   | Go + [huma](https://huma.rocks/), PostgreSQL | Initialized — `GET /health` implemented and tested; no DB/auth/containerization yet |
| Marketing | `marketing/` | Astro (static site) | Skeleton — stack chosen, not yet initialized |

No monorepo tooling links these — they're three independent projects
communicating over REST (client ↔ backend) or not at all (marketing
is a standalone static site).

Coding conventions are tracked in [`CODING_STANDARDS.md`](../CODING_STANDARDS.md),
added per-language as each component initializes.

## Data flow

Not yet implemented. Planned shape: the Flutter client talks to the
Go backend over REST, authenticated with Firebase Auth ID tokens; the
backend is the source of truth in PostgreSQL. The marketing site has
no backend dependency. Offline entries queue locally on the client
and sync once connectivity returns — the sync/conflict strategy is
not yet designed.

## Anticipated future work

Not designed yet, but implied by decisions made so far, each to get
its own backlog/TDR pair when it starts:
- Initializing the Astro marketing project (backend done, see 003;
  client done, see 005)
- PostgreSQL connectivity + migrations, Firebase ID token
  verification, and Docker/CI packaging for the backend
- Firebase Auth integration and the first real product screen
  (with a mockup) for the client
- Entitlements/feature-gating for general vs. premium tiers, and
  billing (Stripe web, App Store/Play Store IAP mobile)
- Vertex AI integration for future AI features
- Migration path from Docker Compose self-hosting to AWS/GCP
- Offline sync/conflict-resolution design

## Decisions

Structural decisions are recorded individually in [`docs/tdr/`](tdr/):
- [001 — Repo skeleton](tdr/001_repo_skeleton_design.md)
- [002 — Tech stack selection](tdr/002_tech_stack_design.md)
- [003 — Backend initialization](tdr/003_backend_init_design.md)
- [004 — Coding standards (Go)](tdr/004_coding_standards_design.md)
- [005 — Client initialization](tdr/005_client_init_design.md)

This document reflects only the current, settled state.
