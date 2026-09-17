# 007 — Marketing site initialization

## User story

As the Shrink! project owner,
I want the Astro marketing project actually initialized with a
working, tested placeholder page,
So that there's a real, buildable static site to add real
pricing/signup/SEO content to, instead of just a documented stack
decision.

## Acceptance criteria

- **AC-1**: `marketing/` is a valid Astro project (package name
  `shrink-marketing`, npm as the package manager, `tsconfig.json` on
  the Strict preset), configured for static output — no server
  adapter.
- **AC-2**: The default Astro starter template's homepage is replaced
  with a minimal placeholder page rendering a single `<h1>Shrink!</h1>`
  and nothing else — no nav, styling, or tagline.
- **AC-3**: A test exists (Vitest + Astro's Container API,
  `astro:container`) that fails against the unmodified starter
  template and passes once the placeholder page is implemented — this
  feature's red-green pair.
- **AC-4**: No pricing/signup/blog content, no styling system, no
  hosting/CI/deployment config, and no mockup/sign-off step are
  required for this feature — the placeholder is not a designed
  marketing page, it's the smallest real thing TDD can anchor to, same
  category as the backend's `/health` endpoint (003) and the client's
  placeholder screen (005). The first real landing page gets a mockup
  when that feature starts.
- **AC-5**: `marketing/README.md` and `docs/architecture.md` are
  updated to reflect marketing's initialized (not skeleton) status.
