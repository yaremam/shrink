# Shrink! — Architecture

Living overview of Shrink!'s system structure. Update this whenever a
feature changes a component boundary, adds a data model, or reverses
an earlier architectural decision.

## Components

| Component | Path       | Status   |
|-----------|------------|----------|
| Backend   | `backend/` | Skeleton — no stack chosen |
| Web       | `web/`     | Skeleton — no stack chosen |
| Android   | `android/` | Skeleton — no stack chosen |
| iOS       | `ios/`     | Skeleton — no stack chosen |

## Data flow

Not yet defined. The backend, web, Android, and iOS components are
expected to communicate over an API to be designed once the backend
stack is chosen.

## Decisions

Structural decisions are recorded individually in [`docs/tdr/`](tdr/).
This document reflects only the current, settled state.
