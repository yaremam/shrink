# 001 — Repo skeleton

## User story

As the Shrink! project owner,
I want a monorepo skeleton with placeholders for the backend, web,
Android, and iOS components,
So that each platform has an agreed home in the repo before any
tech-stack decisions are made.

## Acceptance criteria

- **AC-1**: The repo root contains exactly four top-level component
  directories: `backend/`, `web/`, `android/`, `ios/`.
- **AC-2**: Each component directory contains a `README.md` describing
  its purpose and noting that no tech stack has been chosen yet. No
  other files or subdirectories exist under the component directories.
- **AC-3**: No dependency manifests, build files, or source files for
  any specific language or framework exist anywhere in the repo.
- **AC-4**: The root `README.md` describes Shrink! as a weight
  tracking app and lists all four components with a one-line
  description of each.
- **AC-5**: `docs/architecture.md` exists as a living architecture doc
  listing the four components and their (skeleton) status.
- **AC-6**: `docs/backlog/` and `docs/tdr/` exist and this feature's
  own user story and design decision record are the first entries in
  each, sharing the numbering counter (`001`).
