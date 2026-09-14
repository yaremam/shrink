# 001 — Repo skeleton design

## Context & requirements

Shrink! is a new weight tracking app that will eventually ship as a
backend API, a web frontend, an Android app, and an iOS app. Before
any of those get a tech stack, the repo needs an agreed structure so
future feature work has a home to land in. This pass is scoped to
structure only: no language, framework, or dependency is chosen for
any component here.

The repo already exists on GitHub (`yaremam/shrink`) with an MIT
`LICENSE` and a minimal `README.md`, so the skeleton builds on that
history rather than starting fresh.

Requirements:
- A place for each of the four components that doesn't presuppose
  their eventual tech stack.
- A convention for capturing future feature specs (user story +
  design decision record), per the project's global process.
- A living architecture doc reflecting current (skeleton) state.
- No mixing with unrelated projects already present alongside this
  repo's checkout location.

## Alternatives considered

### Layout: flat top-level dirs vs. `apps/` grouping

**A — Flat top-level directories** (`backend/`, `web/`, `android/`,
`ios/` at repo root)
- Pros: shortest paths; nothing today needs an extra grouping layer;
  matches the mental model of "four components" directly.
- Cons: if non-component concerns accumulate at root (tooling,
  infra-as-code, etc.) later, root can get cluttered.

**B — Grouped under `apps/`** (`apps/backend/`, `apps/web/`, etc.)
- Pros: keeps root reserved for repo-wide concerns (docs, CI config,
  tooling) as the project grows.
- Cons: adds a layer of indirection with no current payoff; nothing
  in the repo today competes with the component dirs for root space.

**Decision: A.** Revisit if/when root-level tooling or infra
directories start competing for space with the four components.

### Repo initialization: fresh `git init` vs. clone existing remote

**A — `git init` fresh, add remote later**
- Pros: simpler, no dependency on the remote's existing state.
- Cons: orphans the commits already on `yaremam/shrink` (LICENSE,
  README); a later `git push` would conflict with that history.

**B — Clone the existing remote, build on top**
- Pros: preserves the existing LICENSE/README history instead of
  discarding it; the eventual push is a fast-forward, not a rewrite.
- Cons: skeleton work has to account for whatever's already there
  (in this case, just a LICENSE and a two-line README).

**Decision: B.** Cloned `yaremam/shrink` into `/workspace/shrink/`
and built the skeleton as new commits on `feature/repo-skeleton`,
branched from `main`.

## Structural decision

- Monorepo at `/workspace/shrink/` (its own repo root, separate from
  unrelated sibling projects already present in `/workspace/`).
- Flat top-level component directories: `backend/`, `web/`,
  `android/`, `ios/` — each holding only a `README.md` for now.
- `docs/backlog/` and `docs/tdr/` established as the spec-doc
  convention, numbered from a shared counter starting at `001`.
- `docs/architecture.md` as the living architecture doc.
- Existing `LICENSE` (MIT) kept as-is; no `CONTRIBUTING.md` (none
  existed on the remote); a generic `.gitignore` added since none
  existed yet.
- Root `README.md` expanded in place to describe the monorepo
  structure and each component.

## Cross-cutting implications

- **Testing**: TDD does not apply to this pass — there is no
  executable behavior, only directory structure and documentation.
  The first component to gain real code is where red-green-refactor
  starts.
- **Future features**: any tech-stack choice for a component (e.g.
  picking the backend language/framework) is its own feature and
  needs its own `docs/backlog/NNN_*.md` + `docs/tdr/NNN_*_design.md`
  pair, continuing this counter.
- **Schema/contract changes**: none — no API or data model exists
  yet.
- **Git history**: this feature's commits sit on `feature/repo-skeleton`,
  local only; pushing to `origin` is a separate, explicit step.
