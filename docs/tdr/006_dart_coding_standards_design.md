# 006 — Coding standards (Dart) design

## Context & requirements

004 scoped Go-only coding standards deliberately, flagging Dart/
TypeScript sections as future work "when those components init,
grounded in real code rather than written speculatively." 005
initialized the client, so that condition is now met. This feature
closes the gap the same way 004 did for Go: tooling for the
mechanical stuff, a written doc for what tooling can't check.

Unlike 003→004, the tooling half already exists here: `flutter
create` (005) generated `client/analysis_options.yaml` with `include:
package:flutter_lints/flutter.yaml` — Flutter's own recommended lint
set, the direct Dart analogue of golangci-lint's `default` set. 005
also already confirmed `flutter analyze` exits clean. So this feature
is scoped to just the written-conventions half.

## Alternatives considered

### Add new/stricter lint config vs. keep the `flutter_lints` default

**Add a stricter package (e.g. `very_good_analysis`) now**
- Cons: same reasoning as 004's linter-set decision — a single-file,
  single-screen client has nothing to calibrate stricter rules
  against yet; premature.

**Keep `flutter_lints` (already in place)**
- Pros: it's Google's own recommended baseline, proportionate to the
  current codebase, and already enforced with zero extra setup.

**Decision: keep `flutter_lints`.** Revisit once there's enough real
client code to justify more, mirroring 004's Go linter reasoning
exactly.

### Doc-only feature vs. folding into a direct edit

Discussed directly with the project owner: does a feature this small
(no new tooling, one doc section) still warrant its own backlog/TDR
pair? **Decision: yes**, full feature treatment, same as 004 — the
process doesn't get waived for a small diff, and it keeps "each
language's standards are their own reviewable unit" consistent rather
than special-casing this one.

## Structural decision (summary)

- `CODING_STANDARDS.md` gains a "Dart" section: widget structure
  (prefer `StatelessWidget`; reach for `StatefulWidget` only when
  local mutable state is genuinely needed), public-API doc comments
  (`///` dartdoc convention), naming (already largely covered by
  `flutter_lints`, noted for completeness), and testing conventions
  (`flutter_test`, matching 005's `pumpWidget` + `find`/`expect`
  pattern).
- No changes to `client/analysis_options.yaml` — it already carries
  the enforced half via `flutter_lints`.

## Cross-cutting implications

- **Testing**: no test changes; `flutter analyze` re-confirmed clean
  as this feature's verification step (already was, per 005).
- **Future features implied**: a TypeScript/Astro section in
  `CODING_STANDARDS.md` once `marketing/` initializes, same pattern.
- **Schema/contract changes**: none.
