# 007 — Marketing site initialization design

## Context & requirements

TDR 002 chose the marketing stack (Astro, static output, separate
from the Flutter client for SEO reasons) but deferred actual project
initialization, same as it did for the backend and client. This
feature does that initialization, scoped the same way 003 and 005
were: the smallest slice that gives the project real, tested behavior,
rather than building out real landing-page content (pricing, signup,
SEO copy) in one pass.

Node 22.23.2 and npm 10.9.8 are already installed in this sandbox
(pre-existing, not installed for this feature). No JS package-manager
precedent applies cleanly from the sibling projects: opusflow uses
pnpm, but only inside a pnpm workspace (monorepo tooling), which 002
already ruled out for marketing as a standalone project.

## Alternatives considered

### Scope: bare `npm create astro` vs. scaffold + tested placeholder

**Bare scaffold** (unmodified Astro starter template)
- Pros: smallest possible change.
- Cons: nothing to write a test against — violates the project's
  TDD-always rule, same reasoning 003 and 005 used to reject a bare
  scaffold for the backend and client.

**Scaffold + placeholder page with a failing-then-passing test**
- Pros: gives this feature a real red-green cycle, mirroring 003's
  `/health` endpoint and 005's placeholder screen exactly.
- Cons: slightly more than "just init," but still small.

**Decision: scaffold + tested placeholder.**

### Mockup + sign-off: required or skipped?

The global process requires a mockup (Artifact) and sign-off before
writing "any new screen or user-facing view." Marketing's placeholder
is arguably a harder case than the client's: a marketing site's whole
purpose is public-facing pages, unlike an app's internal blank screen.

**Require a mockup**
- Cons: there's no real design to mock up yet — the placeholder is
  one line of throwaway text, not a landing page with pricing/signup
  content. Forcing a mockup for that would be process theater.

**Skip the mockup for this specific placeholder**
- Pros: consistent with 005's reasoning for the client's placeholder
  screen — treats this as scaffolding, not product UI. The first real
  landing page (with pricing/signup/SEO content) gets a mockup when
  that feature starts, when there's an actual design to put in front
  of the user.

**Decision: skip it**, same reasoning as 005.

### Test approach: Vitest + Container API vs. Playwright vs. manual build check

**Vitest + Astro's Container API** (`astro:container`)
- Pros: renders a `.astro` component to a string in Node, no browser
  needed — fast, closest in spirit to the backend's `httptest`-based
  unit test and the client's widget test.
- Cons: newer/less battle-tested API than Playwright; can't test
  client-side JS interactivity (not needed — this page has none).

**Playwright** (real browser against the built/dev site)
- Pros: tests the actual rendered output end-to-end, including any
  future client-side behavior.
- Cons: heavier — browser install, slower runs — overkill for one
  static placeholder page with no interactivity.

**Manual `npm run build` check, no automated test**
- Cons: no red-green cycle — fails the project's TDD-always rule, same
  reasoning that ruled out a bare scaffold above.

**Decision: Vitest + Container API.**

### Package manager: npm vs. pnpm vs. yarn

**npm**
- Pros: already installed in this sandbox; Astro's own default
  tooling (`npm create astro@latest`); no extra install.
- Cons: none material for a single standalone project.

**pnpm**
- Pros: faster installs, stricter dependency resolution.
- Cons: not installed in this sandbox; its main precedent (opusflow)
  is workspace-oriented, which doesn't apply here — marketing has no
  sibling packages to link.

**yarn**
- No precedent, no advantage over npm for this project.

**Decision: npm.**

### TypeScript strictness

`npm create astro@latest` prompts for a `tsconfig.json` preset:
Relaxed, Strict, or Strictest. No prior TS precedent exists elsewhere
in this repo (backend is Go, client is Dart).

**Decision: Strict** — Astro's own recommended default, consistent
with the project's general bias toward boring, well-supported
defaults over exotic configuration (same spirit as 002's toolchain
choices).

### Project/package naming

The backend module is `github.com/yaremam/shrink/backend`; the
client's `pubspec.yaml` name is `shrink`.

**`marketing`**: matches the directory name, plain.
**`shrink`**: matches the client's bare choice, but would collide in
meaning across three differently-scoped `package.json`/
`pubspec.yaml`/`go.mod` names if ever compared side by side.
**`shrink-marketing`**: ties the package to the product by name.

**Decision: `shrink-marketing`** — avoids the bare-`shrink` collision
now that a third package name is in play.

### Placeholder content

Mirroring the client's literal `"Shrink!"` text screen and the
backend's literal `{"status":"ok"}` response.

**Decision:** a single `<h1>Shrink!</h1>` on the index page, nothing
else — no tagline, nav, or styling. Real landing-page content (pricing,
signup, SEO copy) is explicit future work per `docs/architecture.md`.

### Output mode and hosting

TDR 002 already decided marketing is a static site, separate from the
Flutter client, specifically for SEO reasons. This feature doesn't
revisit that — static output, no server adapter, no hosting/CI/
deployment config. Same "defer the parts with nothing real to
configure yet" reasoning as 003 deferring PostgreSQL and 005 deferring
Firebase Auth.

### Branch basis

`feature/marketing-init` off `main` vs. stacked on an unmerged sibling
branch.

**Decision: off `main`**, independent — marketing shares no code or
package graph with the backend or client (per 002's "no monorepo
tooling" decision), so nothing in this feature depends on any other
branch's contents. Same reasoning as 005's client-init branch basis.

## Structural decision (summary)

- `marketing/`: Astro project, package name `shrink-marketing`, npm,
  `tsconfig.json` on Strict, static output (no adapter).
- `marketing/src/pages/index.astro`: minimal placeholder page
  rendering `<h1>Shrink!</h1>`, replacing the default starter
  template's homepage.
- `marketing/src/pages/index.test.ts` (or equivalent): Vitest test
  using `astro:container` asserting the placeholder text renders —
  this feature's red-green pair.
- No pricing/signup/blog content, no styling system, no hosting/CI/
  deployment config, no mockup step.

## Cross-cutting implications

- **Testing**: this is the feature where automated testing starts for
  marketing, mirroring 003 (backend) and 005 (client). Established
  default: Vitest + Astro's Container API.
- **Future features implied**: the first real landing page (pricing,
  signup, SEO content) with a mockup + sign-off, a TypeScript section
  in `CODING_STANDARDS.md` once there's enough real marketing code to
  ground it (mirrors 004 and 006's scoping reasoning), and hosting/CI/
  deployment config for the marketing site.
- **Schema/contract changes**: none — marketing has no backend
  dependency (per 002, it's a standalone static site).
- **Environment note**: Node 22.23.2 and npm 10.9.8 were already
  present in this sandbox; nothing new needed installing for this
  feature, unlike 003 (Go/Flutter) and 005 (Flutter sandbox
  toolchain caveats).
