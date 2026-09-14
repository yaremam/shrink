# 002 — Tech stack selection

## User story

As the Shrink! project owner,
I want a chosen, documented tech stack for every component,
So that future features have an agreed foundation to build on instead
of re-litigating language/framework choices feature by feature.

## Acceptance criteria

- **AC-1**: A design decision record exists documenting at least two
  alternatives considered for each of: mobile/web client, marketing
  site, backend language/framework, database, API style, and auth,
  with pros/cons and a stated decision for each.
- **AC-2**: The repo structure reflects the chosen components: a
  single `client/` (Flutter, mobile + web), a `backend/` (Go), and a
  `marketing/` (static site) — replacing the generic `android/`,
  `ios/`, `web/` split from the repo skeleton (001).
- **AC-3**: `docs/architecture.md` is updated to reflect the chosen
  stack and component list.
- **AC-4**: No project initialization (`flutter create`, `go mod
  init`, framework scaffolding, dependencies) happens as part of this
  feature — this feature is the decision and structure only. Actual
  project setup is separate future work, each following the TDD
  process.
- **AC-5**: The decision record notes premium/general tiering and
  Vertex AI integration as anticipated future concerns without
  designing them now.
