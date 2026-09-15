# 003 — Backend initialization

## User story

As the Shrink! project owner,
I want the Go backend project actually initialized with a working,
tested health check,
So that there's a real, buildable project to add API endpoints to,
instead of just a documented stack decision.

## Acceptance criteria

- **AC-1**: `backend/` is a valid Go module at
  `github.com/yaremam/shrink/backend` (`go.mod` present, `go 1.27.1`).
- **AC-2**: The project layout is `cmd/server/main.go` (entry point)
  and `internal/httpserver/` (routing/handlers), matching opusflow's
  established Go layout.
- **AC-3**: The server uses [huma](https://huma.rocks/) for routing
  and exposes `GET /health`, returning `200` with JSON body
  `{"status":"ok"}`.
- **AC-4**: A test exists (stdlib `testing` + `net/http/httptest` +
  `testify/require`) that fails against an empty handler and passes
  once `/health` is implemented — this feature's red-green pair.
- **AC-5**: No PostgreSQL connection, migration tooling, Docker/
  docker-compose files, CI workflow, or Firebase auth verification are
  added as part of this feature — each is separate future work.
- **AC-6**: `backend/README.md` and `docs/architecture.md` are updated
  to reflect the backend's initialized (not skeleton) status.
