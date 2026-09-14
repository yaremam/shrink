# 002 — Tech stack selection design

## Context & requirements

Feature 001 scaffolded the repo with placeholder `backend/`, `web/`,
`android/`, `ios/` directories and no tech stack chosen for any of
them. This feature picks the stack.

Constraints gathered before evaluating alternatives:

- **Audience**: public, multi-tenant. Data lives in the backend, not
  only on-device. Export and delete are required from day one
  (data-portability / right-to-erasure).
- **Offline**: logging a weight entry must work fully offline and
  sync later — this is a hard requirement, not a nice-to-have.
- **Ambition**: a real product to keep building on, not a throwaway
  MVP — favors boring, well-supported technology.
- **Tiers**: the product will distinguish general vs. premium
  features. This implies an eventual entitlements/feature-gating
  layer and billing (Stripe on web, App Store/Play Store IAP on
  mobile) — anticipated, not designed, in this feature.
- **Hosting**: Docker Compose self-hosted to start, with an intended
  path to AWS/GCP later. Vertex AI is planned for future AI features
  — a data point for backend language choice, not a hard requirement
  yet.
- **Infra independence**: Shrink! gets its own Postgres/storage,
  fully decoupled from the sibling projects (docuflow, opusflow),
  though it follows their proven *patterns* (Docker Compose,
  docs/backlog + docs/tdr convention).
- **Prior-stack evidence** (found by inspecting the sibling projects,
  not assumed): docuflow uses Rust/axum/Postgres/self-hosted Docker
  Compose; opusflow uses a Go backend and a single Expo/React Native
  `mobile/` codebase in a pnpm workspace. Neither uses Python or
  Flutter for a backend/client, though the project owner has prior
  Flutter experience.
- **Language comfort**: Rust, Go, TypeScript/Node (from prior
  projects), Python 3.13+ (standard/strict typing), C#/.NET
  (strongest), Kotlin, and Flutter/Dart are all viable; C++ is
  explicitly out of scope ("for fun" only).

## Alternatives considered

### Mobile client: React Native vs. Flutter

**React Native (Expo)**
- Pros: proven pattern already in opusflow (`app.json`, `eas.json`,
  Expo managed workflow); TypeScript, same language as a
  TypeScript web frontend would use, enabling shared types/validation
  across client and web via a workspace.
- Cons: bridges to native components rather than owning rendering;
  no native web target (a web frontend would need a separate
  framework/codebase).

**Flutter**
- Pros: the project owner has prior Flutter experience; single
  codebase natively targets iOS, Android, *and* web (`flutter create`
  generates all platform targets from one project — no extra
  multi-package tooling needed for a single app); own rendering
  engine (Impeller/Skia) gives consistent, pixel-accurate UI across
  platforms, relevant for weight-trend charts; mature offline-storage
  options (drift, sqflite, Isar); RevenueCat supports Flutter for
  future IAP.
- Cons: Dart, not shared with a TypeScript backend or (if kept
  separate) a TypeScript web frontend — no cross-client code sharing
  the way opusflow's pnpm workspace allows; diverges from opusflow's
  established React Native pattern.

**Decision: Flutter**, specifically as a **single `client/` codebase
covering mobile and web** — collapsing the original `android/`,
`ios/`, and `web/` skeleton directories into one. This was chosen
over keeping React Native + a separate web framework because a
single codebase covering three platforms outweighed React Native's
opusflow precedent, given the project owner already has Flutter
experience.

### Public-facing web presence: Flutter Web vs. a separate marketing site

Once the client became Flutter, Flutter Web could theoretically serve
public marketing/signup pages too.

**Flutter Web for everything (no separate site)**
- Pros: nothing extra to build or maintain; one codebase, one
  deployment.
- Cons: Flutter Web has known SEO weaknesses — no true SSR, a
  JS/Wasm-heavy initial load, and poor indexability. For a public,
  multi-tenant product, organic discoverability of marketing/signup
  pages matters.

**Separate static marketing site (`marketing/`), Flutter Web for the
authenticated app only**
- Pros: SEO-crawlable, fast-loading marketing pages using tooling
  built for that job, without dragging SEO constraints into the app
  itself.
- Cons: a second, independent codebase/deployment to maintain.

**Decision: separate static marketing site**, built with **Astro**
(over plain static HTML or a full app framework like Next.js) because
a public product with a premium tier tends to grow a pricing page,
blog/changelog, and SEO content over time, and Astro handles that
without pulling in a full application framework the marketing site
doesn't need. The in-app experience (post-signup) stays entirely
Flutter.

### Backend language: Go vs. Python vs. C#/.NET vs. Rust

All four are viable per the project owner's language comfort. Node
was ruled out early: its main advantage (sharing TypeScript with a
web frontend) disappeared once the client became Flutter/Dart, and
Rust's advantage (docuflow precedent) was outweighed by slower
iteration speed for CRUD/entitlements-heavy work. The real comparison
came down to Go vs. Python:

**Go**
- Pros: proven operational pattern already in opusflow (`go.mod`,
  `cmd/`+`internal/` layout, Dockerfile conventions directly
  reusable); goroutines give real multi-core parallelism for both
  I/O- and CPU-bound work with no extra plumbing; smaller container
  images, faster cold starts (relevant for autoscaling on Cloud Run
  later); official Go SDKs exist for Vertex AI
  (`cloud.google.com/go/aiplatform`) and the newer unified
  `google.golang.org/genai` package — calling Vertex AI is not a
  real gap, just occasionally translating Python-first sample code.
- Cons: no framework in the ecosystem generates an OpenAPI spec as
  automatically as FastAPI does by default (mitigated — see `huma`
  below); more verbose than Python for CRUD-heavy business logic.

**Python 3.13 + FastAPI**
- Pros: FastAPI derives the OpenAPI spec directly from Pydantic
  models and route signatures with zero extra tooling; Vertex AI's
  Python SDK is the most complete and is where Google's own
  docs/samples land first; Pydantic v2 gives strong request/response
  validation for free.
- Cons: brand-new backend stack for this project owner — no existing
  Dockerfile/CI pattern to reuse (docuflow is Rust, opusflow is Go);
  CPython's GIL means true parallelism requires multiple worker
  processes rather than Go's native multi-core goroutines — not a
  real issue for the mostly I/O-bound workload today (async
  `asyncio`/Uvicorn handles concurrent I/O fine, the same model
  Node.js uses), but a genuine structural disadvantage if CPU-bound
  features appear later (image processing, PDF export generation,
  local analytics).

**C#/.NET**
- Pros: the project owner's strongest language, likely fastest
  development velocity; performance comparable to Go; mature
  ecosystem (EF Core, Stripe .NET SDK).
- Cons: no official Vertex AI SDK, only its REST API; no prior
  operational precedent in either sibling project.

**Rust (axum)**
- Pros: proven in docuflow, reusable Postgres/S3/auth/tracing
  patterns.
- Cons: slowest iteration speed of the group for CRUD/entitlements
  work; no official Vertex AI SDK either.

**Decision: Go.** Once the Vertex AI gap was found to be smaller than
initially assumed (Go has official SDKs; the difference is
documentation-first-language, not capability) and the GIL/performance
question was weighed, Go's proven opusflow precedent and native
multi-core parallelism outweighed Python's OpenAPI auto-generation
convenience — which Go closes separately (see below).

### Go web framework: net/http vs. chi vs. huma

**stdlib `net/http`** (opusflow's pattern) / **chi**
- Pros: minimal, matches existing precedent exactly.
- Cons: the OpenAPI spec has to be hand-maintained or generated via
  separate annotation tooling — can drift from the code.

**huma**
- Pros: built on top of `net/http`/chi; generates the OpenAPI spec
  directly from Go struct/handler definitions, the closest Go
  equivalent to FastAPI's auto-generation. This matters because the
  Flutter client's API bindings will be code-generated from that
  spec — an auto-generated, drift-proof spec was the one concrete
  advantage Python held, and huma closes it.
- Cons: an additional framework dependency beyond stdlib.

**Decision: huma.**

### Database: PostgreSQL vs. MySQL vs. MongoDB

**PostgreSQL**
- Pros: proven in both sibling projects; strong relational fit for
  multi-tenant data (users, weight entries, entitlements,
  subscriptions); scales to managed Cloud SQL later.
- Cons: none material here.

**MySQL/MariaDB**: similarly mature, but no prior use and no
advantage over Postgres for this data shape.

**MongoDB**: flexible schema, but multi-tenant + billing-adjacent
data is relational by nature and benefits from strong transactional
guarantees; no prior use.

**Decision: PostgreSQL.**

### API style: REST vs. GraphQL vs. tRPC

**REST + OpenAPI**
- Pros: language-agnostic — stays valid regardless of backend
  language choice; well-understood offline-sync patterns (queued
  mutations, cursors/ETags); pairs directly with huma's
  spec-generation and Flutter client codegen from that spec.
- Cons: client type-sync depends on codegen tooling (solved via the
  OpenAPI spec).

**GraphQL**: flexible querying, but adds resolver/N+1 operational
complexity not justified by this app's relatively simple data model.

**tRPC**: zero-codegen type safety, but only works with a TypeScript
backend — ruled out once Go was chosen (and would have been ruled
out anyway once the client became Flutter/Dart).

**Decision: REST + OpenAPI.**

### Auth: self-managed vs. managed provider vs. passwordless-only

**Self-managed (own sessions/JWTs, own password storage)**
- Pros: full control, no external dependency.
- Cons: significant security-sensitive surface to build and maintain
  (password hashing, reset flows, email verification) for a public
  multi-tenant product.

**Managed provider (Firebase Auth / Auth0 / Supabase Auth)**
- Pros: removes password storage and reset/verification flows from
  the backend entirely; social login (Google/Apple) included.
- Cons: external dependency; provider lock-in to some degree.

**Passwordless-only (magic link/OTP)**
- Pros: simpler mental model, no passwords to manage.
- Cons: fully untested assumption about user preference for a
  general-audience product; not chosen without evidence it fits.

**Decision: Firebase Auth**, a managed provider. Chosen over Auth0/
Supabase specifically because it's GCP-native (aligned with the
planned AWS/GCP + Vertex AI direction), has first-class Flutter SDKs
covering both mobile and web from one integration, and includes
Google/Apple social login plus email/password out of the box. The Go
backend verifies Firebase ID tokens rather than managing credentials
itself.

### Repo/monorepo tooling

With the client collapsed into one Flutter project and the backend a
separate Go project, there's no shared package graph between them —
they only communicate over REST. `marketing/` is a third, fully
independent static site. **Decision: no monorepo tool** (no pnpm
workspaces, Melos, or Nx) — three plain top-level directories, each
with its own build tooling, mirroring the simplicity established in
001.

## Structural decision (summary)

- `client/` — Flutter (mobile + web), Firebase Auth, offline-first
  local storage with sync.
- `backend/` — Go + huma, PostgreSQL, REST + OpenAPI, verifies
  Firebase ID tokens.
- `marketing/` — Astro static site, SEO-crawlable, separate from the
  Flutter client.
- No monorepo tooling; three independent projects.
- Docker Compose for self-hosting now, with an intended (not yet
  designed) migration path to AWS/GCP.

## Cross-cutting implications

- **Repo structure**: replaces the `android/`/`ios/`/`web/` split
  from 001 with `client/` and adds `marketing/`; `backend/` is
  unchanged in location, now with a stack attached.
- **Testing**: still N/A for this feature — no code is written here,
  only structure and documentation. TDD begins with the first
  feature that adds real backend or client behavior (expected
  defaults: Go's standard `testing` package + `testify` for the
  backend, `flutter_test` for the client — to be confirmed when that
  feature starts).
- **Future features implied by this decision**, each needing its own
  backlog/TDR pair: initializing the Go backend project (`go mod
  init`, huma, Postgres migrations), initializing the Flutter client
  project, initializing the Astro marketing site, designing the
  entitlements/premium-tier model and billing integration, and
  designing the Vertex AI integration point.
- **Schema/contract changes**: none yet — no API contract exists
  until the backend project is initialized.
- **Data portability**: export/delete being a hard requirement means
  the eventual data model and any managed-auth integration (Firebase)
  must support full account data export and deletion, including
  deleting the Firebase Auth identity itself.
