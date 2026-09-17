# Shrink! — Marketing

The public marketing/landing site for Shrink! (pricing, signup,
SEO-crawlable content), kept separate from the Flutter client because
Flutter Web isn't SEO-friendly.

Initialized as an Astro project (package `shrink-marketing`, static
output, TypeScript strict) with a tested placeholder page — see
[`docs/tdr/007_marketing_init_design.md`](../docs/tdr/007_marketing_init_design.md).
Real landing-page content (pricing, signup, SEO copy) is separate
future work; see
[`docs/tdr/002_tech_stack_design.md`](../docs/tdr/002_tech_stack_design.md)
for the original stack decision.

Run `npm test` for the Vitest suite (uses Astro's Container API, no
browser needed), `npm run dev` for the local dev server, and
`npm run build` to produce the static `dist/` output.
