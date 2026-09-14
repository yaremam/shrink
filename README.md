# Shrink!

Weight loss tracking.

Shrink! is made up of three components. The stack is chosen (see
[docs/tdr/002](docs/tdr/002_tech_stack_design.md)) but not yet
initialized:

| Component | Path | Stack | Description |
|-----------|------|-------|-------------|
| Client | [`client/`](client/) | Flutter | Mobile (iOS/Android) + web app |
| Backend | [`backend/`](backend/) | Go + huma, PostgreSQL | API and business logic |
| Marketing | [`marketing/`](marketing/) | Astro | Public marketing/landing site |

## Docs

- [`docs/architecture.md`](docs/architecture.md) — living architecture overview
- [`docs/backlog/`](docs/backlog/) — user stories, one per feature
- [`docs/tdr/`](docs/tdr/) — design decision records, one per feature
