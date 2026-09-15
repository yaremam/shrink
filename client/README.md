# Shrink! — Client

The Flutter client for Shrink!, covering mobile (iOS/Android) and web
from a single codebase.

Initialized as a Flutter project (package `shrink`, app id
`com.shrink.app`, platforms iOS/Android/web) with a tested placeholder
screen — see
[`docs/tdr/005_client_init_design.md`](../docs/tdr/005_client_init_design.md).
Firebase Auth integration and real product screens are separate
future features; see
[`docs/tdr/002_tech_stack_design.md`](../docs/tdr/002_tech_stack_design.md)
for the original stack decision.

- Auth: Firebase Auth (email/password + Google/Apple sign-in) — not
  yet integrated
- Data entry (e.g. weight logging) is offline-first with sync — not
  yet implemented

Run `flutter test` from this directory. Note: the `flutter_tester`
subprocess is intermittently flaky in some sandboxed dev
environments (segfaults inside flutter_tools itself, not application
code) — re-run on an unexplained failure before assuming a real
regression; see 005's design doc for details.
