# 004 — Coding standards (Go) design

## Context & requirements

Feature 003's code review (Standards axis) found the repo had zero
documented coding standards — no `CONTRIBUTING.md`,
`CODING_STANDARDS.md`, `.golangci.yml`, or `.editorconfig` anywhere.
The reviewer fell back entirely to a generic Fowler-smell baseline and
flagged Go-idiom nits (a missing package doc comment on
`internal/httpserver`) as judgement calls rather than citable rule
violations. This feature closes that gap for Go, the only language
with real code so far.

Requirements gathered before scoping:
- Go only — `client/` (Dart) and `marketing/` (TypeScript/Astro) are
  still skeletons; writing their conventions now would be guessing
  without real code to calibrate against.
- Both a written doc (for judgement-requiring conventions) and
  enforced tooling (for mechanical issues) — neither alone closes the
  gap the review found.
- Retroactively fix whatever the new linter flags in 003's code,
  rather than landing tooling that immediately fails on `main`.
- A real red-green verification cycle: run the linter before fixes
  (fails), apply fixes, run again (passes) — even though this isn't a
  unit test, it's the same shape the project's TDD-always rule cares
  about.

## Alternatives considered

### Scope: Go-only vs. all three languages now

**All three now**
- Pros: one feature covers the whole project's conventions at once.
- Cons: `client/` and `marketing/` don't exist yet — there's no real
  Dart or TypeScript code to derive conventions from, so anything
  written would be speculative and likely wrong or irrelevant once
  those components actually init.

**Go-only**
- Pros: grounded in code that actually exists and was just reviewed;
  matches the project's established one-component-at-a-time pattern.
- Cons: `CODING_STANDARDS.md` will need Dart/TS sections appended
  later — acceptable, same shape as `docs/architecture.md` already
  growing incrementally per component.

**Decision: Go-only.**

### Doc vs. tooling vs. both

**Written doc only**
- Pros: covers nuance (why a convention exists) tooling can't express.
- Cons: nothing stops a future change from violating it silently —
  review has to catch every instance by eye, the exact gap 003's
  review just hit.

**Enforced tooling only (`.golangci.yml`)**
- Pros: mechanical issues (unused vars, ineffectual assignments,
  missing error checks, import ordering) get caught automatically.
- Cons: linters can flag *that* a package doc comment is missing
  (`revive`'s package-comments rule) but not judge doc-comment
  *quality*, naming conventions, or project-specific patterns like
  which testing stack to use — the review found idiom gaps a linter
  alone wouldn't fully close either.

**Decision: both.** Tooling handles what's mechanically checkable;
the doc handles what requires judgement.

### File locations

**`.golangci.yml` at repo root vs. inside `backend/`**

Repo root has no `go.mod` — golangci-lint operates per Go module, and
the module lives at `backend/go.mod`. **Decision: `backend/.golangci.yml`.**

**`CODING_STANDARDS.md`: one root file vs. per-component files**

One root file with a "Go" section (mirrors how `docs/architecture.md`
is already one living document covering all three components, not
three separate files) vs. `backend/CODING_STANDARDS.md` +
future `client/CODING_STANDARDS.md` / `marketing/CODING_STANDARDS.md`.

**Decision: one root `CODING_STANDARDS.md`**, consistent with the
existing `docs/architecture.md` pattern. "Dart" and "TypeScript"
sections get appended when those components initialize.

### Linter set: default vs. broad/strict

**golangci-lint `default` set + `revive` (package-comments rule)**
- Pros: catches the actual gap found (missing package doc comments)
  plus standard mechanical issues (errcheck, govet, ineffassign,
  staticcheck, unused); proportionate to a single-file, single-endpoint
  codebase.
- Cons: doesn't enforce complexity/length limits (`gocyclo`,
  `funlen`), import-grouping style (`gci`), etc.

**Broader/stricter set now**

Cons: nothing in the current codebase (one handler, one test) is
complex enough to calibrate limits like cyclomatic complexity or
function length against — those thresholds would be guesses.

**Decision: default + `revive` (package-comments) now.** Revisit and
add more linters once there's enough real code for their limits to
mean something.

### TDD applicability

`golangci-lint run ./...` before the missing package doc comment is
fixed = red (flags the gap). Adding the comment = green. Not a unit
test, but the same red-green shape the project's TDD-always rule
requires — unlike 001/002, which were pure documentation with no
checkable artifact at all. **Decision: this counts as this feature's
red-green cycle**, run and confirmed as part of implementation.

## Structural decision (summary)

- `backend/.golangci.yml`: `default` linters + `revive`
  (package-comments rule enabled).
- Root `CODING_STANDARDS.md`: a "Go" section covering package/exported-
  symbol doc comments, naming, error handling, and testing conventions
  (stdlib `testing` + `httptest` + `testify/require`, matching 003).
- `backend/internal/httpserver/httpserver.go` gets its package doc
  comment added (and any other `golangci-lint` finding fixed) as part
  of this feature.
- `docs/architecture.md` references `CODING_STANDARDS.md`.

## Cross-cutting implications

- **Testing**: no unit tests added or changed by this feature; the
  verification is `golangci-lint run ./...` exiting `0`, run as a
  red-green pair against the pre-existing 003 code.
- **Future features implied**: Dart conventions when `client/` inits
  (004's Go pattern is the template — grounded in real code, not
  written speculatively); TypeScript/Astro conventions when
  `marketing/` inits; broader Go linters once there's enough code to
  calibrate complexity/length rules against; CI wiring `golangci-lint
  run` into a pipeline (no CI exists yet, per 003's deferred scope).
- **Schema/contract changes**: none.
