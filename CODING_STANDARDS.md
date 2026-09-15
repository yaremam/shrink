# Shrink! — Coding standards

Per-language sections, added as each component initializes (see
[`docs/architecture.md`](docs/architecture.md)). Mechanical rules are
enforced by tooling where possible; this document covers what tooling
can't check.

## Go (`backend/`)

Enforced by `backend/.golangci.yml` (golangci-lint's `default` set —
errcheck, govet, ineffassign, staticcheck, unused — plus `revive`'s
package-comments rule). Run `golangci-lint run ./...` from `backend/`
before committing.

Not enforceable by tooling, so covered here instead:

- **Doc comments**: every package gets a `// Package x ...` comment
  (enforced by revive); every exported identifier (type, func, const)
  gets a doc comment starting with its own name, per standard Go
  convention (`godoc` relies on this; not currently linter-enforced,
  but expected).
- **Errors**: don't discard errors. If an error genuinely can't occur
  or genuinely doesn't matter, say why in a comment rather than a bare
  `_ =` with no explanation — `errcheck` catches the discard itself,
  not whether discarding it was the right call.
- **Layout**: `cmd/<binary>/main.go` for entry points, `internal/` for
  everything not meant to be imported by other modules — established
  in `cmd/server` + `internal/httpserver` (see
  [`docs/tdr/003_backend_init_design.md`](docs/tdr/003_backend_init_design.md)).
- **Testing**: stdlib `testing` + `net/http/httptest` for HTTP
  handlers, `testify/require` for assertions (fail-fast; use `require`
  over `assert` unless a test intentionally needs to continue after a
  failed check to report multiple independent problems).
- **HTTP handlers**: use huma (typed request/response structs) rather
  than writing to `http.ResponseWriter` directly — keeps the OpenAPI
  spec accurate without separate annotation, per
  [`docs/tdr/002_tech_stack_design.md`](docs/tdr/002_tech_stack_design.md).
