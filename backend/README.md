# Shrink! — Backend

The API and business logic backing the Shrink! weight tracking app.

This is currently a skeleton. Project initialization (`go mod init`
and everything that follows) is separate future work, tracked as its
own feature per the project's TDD process — see
[`docs/tdr/002_tech_stack_design.md`](../docs/tdr/002_tech_stack_design.md)
for the stack decision behind this.

- Language/framework: Go + [huma](https://huma.rocks/) (generates the
  OpenAPI spec directly from code)
- Database: PostgreSQL
- API style: REST + OpenAPI
- Auth: verifies Firebase Auth ID tokens
