# 003 — Backend initialization design

## Context & requirements

TDR 002 chose the backend stack (Go + huma + PostgreSQL) but explicitly
deferred actual project initialization (AC-4 of 002: "no project
initialization... happens as part of this feature — this feature is
the decision and structure only"). This feature does that
initialization for the backend, scoped down to the smallest slice that
gives the project real, tested behavior — a health check — rather than
building out the full stack (DB, auth, containerization) in one pass.

The environment this feature was built in lacked Go, Flutter, and a
working Docker daemon at the start. Go 1.27.1 and Flutter 3.47.4 were
installed (see session history); Docker was installed but its daemon
cannot start in this sandbox regardless of storage driver (`vfs`) or
`--iptables=false` — it fails on network-namespace/NAT setup
(`operation not permitted` creating the bridge network), which points
to missing capabilities in the sandbox itself, not a fixable
config issue. That unblocked Go and Flutter work but ruled out any
feature that needs to actually run `docker build`/`docker compose up`
here.

Requirements gathered before scoping:
- Backend, client, and marketing are each their own feature/backlog/
  TDR unit, done one at a time — not one combined "initialize
  everything" feature.
- Backend goes first, since the client's API bindings are meant to be
  code-generated from the backend's OpenAPI spec (huma decision in
  002) — the client has more to gain from a real backend existing
  first than vice versa.
- Whatever ships must have a real TDD red-green cycle, not a bare
  scaffold with nothing to test.

## Alternatives considered

### Scope: bare scaffold vs. scaffold + smoke test

**Bare scaffold** (`go mod init` + empty `cmd/server/main.go`)
- Pros: smallest possible change.
- Cons: nothing to write a test against — violates the project's
  TDD-always rule, since there's no behavior for red-green-refactor to
  attach to.

**Scaffold + `/health` endpoint with a failing-then-passing test**
- Pros: gives this feature an actual red-green cycle; huma's
  OpenAPI-generation gets exercised for real (a typed response struct,
  not a bare `http.StatusOK` write) instead of only on paper.
- Cons: slightly more than "just init," but still small.

**Decision: scaffold + smoke test.** A do-nothing scaffold isn't
compatible with TDD; the health check is the smallest unit of real
behavior available.

### How much of the stack to wire up now: health-only vs. health + PostgreSQL

**Health-only, no DB**
- Pros: keeps this feature's blast radius to routing/handler code;
  no connection config, migration tooling, or docker-compose
  dependency pulled in.
- Cons: doesn't yet prove the PostgreSQL piece of the stack works.

**Health + PostgreSQL connectivity**
- Pros: validates the DB half of the stack now.
- Cons: pulls in migration tooling choice, connection config/secrets
  handling, and — practically — a way to run Postgres locally, which
  in this sandbox means Docker Compose, which cannot run here (see
  Context above). Even outside this sandbox's limitation, there's no
  schema yet to migrate, so a DB connection now would have nothing
  real to do.

**Decision: health-only.** PostgreSQL wiring becomes its own next
feature once there's an actual table to store (e.g. the first data
model). This also sidesteps the sandbox's Docker limitation rather
than working around it.

### Go version to pin in `go.mod`: match opusflow's `1.24.0` vs. use the installed `1.27.1`

**Match opusflow (`1.24.0`)**
- Pros: version consistency across sibling projects.
- Cons: arbitrary downgrade from what's actually installed and
  tested; 002 already established shrink has its own independent
  infra (Postgres/storage) decoupled from the siblings — the same
  reasoning applies to toolchain versions.

**Use installed (`1.27.1`)**
- Pros: matches what was actually installed, built, and tested in
  this session; `GOTOOLCHAIN=auto` means a lower pin would just
  self-upgrade on first build anyway.

**Decision: `1.27.1`.**

## Structural decision (summary)

- `backend/go.mod`: module `github.com/yaremam/shrink/backend`, `go
  1.27.1`.
- `backend/cmd/server/main.go`: entry point, starts the huma/http
  server.
- `backend/internal/httpserver/`: router setup and the `/health`
  handler (huma operation, typed response struct →
  `{"status":"ok"}`).
- `backend/internal/httpserver/health_test.go`: the red-green test,
  using `httptest` + `testify/require`.
- No PostgreSQL, no Docker/docker-compose, no CI workflow, no Firebase
  token verification — each is separate future work per
  `docs/architecture.md`'s anticipated-work list.

## Cross-cutting implications

- **Testing**: this is the feature where TDD starts for the backend,
  per 002's note. Established defaults: stdlib `testing` +
  `net/http/httptest` + `testify/require`.
- **Future features implied**: PostgreSQL connectivity + migrations,
  Firebase ID token verification middleware, Docker/docker-compose
  packaging for the backend, CI workflow, and — once the backend has
  a real OpenAPI spec — client API-binding codegen.
- **Schema/contract changes**: the first real API contract now exists
  (`GET /health`), the seed for the OpenAPI spec the client will
  eventually codegen against. No data model yet.
- **Environment note**: Go 1.27.1 and Flutter 3.47.4 are now installed
  in this sandbox (system-wide, via `/etc/profile.d/go.sh` and
  `~/.bashrc`); Docker is installed but its daemon does not run here.
  This is sandbox-specific, not a project decision — a real
  deployment/CI environment is expected to have a working Docker
  daemon.
