# Shrink! — Backend

The API and business logic backing the Shrink! weight tracking app.

Initialized as a Go module with a working, tested `GET /health`
endpoint — see
[`docs/tdr/003_backend_init_design.md`](../docs/tdr/003_backend_init_design.md).
PostgreSQL connectivity, Firebase auth verification, and Docker/CI
packaging are separate future features; see
[`docs/tdr/002_tech_stack_design.md`](../docs/tdr/002_tech_stack_design.md)
for the original stack decision.

- Language/framework: Go + [huma](https://huma.rocks/) (generates the
  OpenAPI spec directly from code)
- Database: PostgreSQL
- API style: REST + OpenAPI
- Auth: verifies Firebase Auth ID tokens
