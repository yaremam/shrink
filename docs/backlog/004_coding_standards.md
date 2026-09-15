# 004 — Coding standards (Go)

## User story

As the Shrink! project owner,
I want a documented, partially-enforced coding standard for the Go
backend,
So that future code review can cite a concrete rule instead of
falling back to generic judgement calls, and mechanical issues get
caught by tooling instead of a human.

## Acceptance criteria

- **AC-1**: `backend/.golangci.yml` exists, enabling golangci-lint's
  `default` linter set plus `revive` configured for the
  package-comments rule.
- **AC-2**: `golangci-lint run ./...` (from `backend/`) exits `0`.
- **AC-3**: Root `CODING_STANDARDS.md` exists with a "Go" section
  covering conventions the linter can't check: package/exported-symbol
  doc comments, naming, error-handling patterns, and testing
  conventions (stdlib `testing` + `httptest` + `testify/require`, per
  003).
- **AC-4**: Any issue `golangci-lint run ./...` flags in the existing
  `backend/` code (from feature 003) is fixed as part of this feature
  — including, at minimum, the missing package doc comment on
  `internal/httpserver` found during 003's code review.
- **AC-5**: Scoped to Go only. No Dart or TypeScript/Astro conventions
  are written — `client/` and `marketing/` aren't initialized yet, and
  writing conventions for languages with no real code would be
  speculative. Their sections get added when those components init.
- **AC-6**: `docs/architecture.md` references the new standards doc.
