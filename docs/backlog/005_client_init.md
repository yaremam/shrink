# 005 — Client initialization

## User story

As the Shrink! project owner,
I want the Flutter client project actually initialized with a working,
tested placeholder screen,
So that there's a real, buildable app to add real screens to, instead
of just a documented stack decision.

## Acceptance criteria

- **AC-1**: `client/` is a valid Flutter project (package name
  `shrink`, org `com.shrink`, app id `com.shrink.app`), scaffolded for
  platforms `ios`, `android`, `web` only — no desktop targets.
- **AC-2**: The default `flutter create` counter-app template is
  replaced with a minimal placeholder screen displaying "Shrink!".
- **AC-3**: A widget test exists (`flutter_test`) that fails against
  the unmodified default template and passes once the placeholder
  screen is implemented — this feature's red-green pair.
- **AC-4**: No Firebase Auth integration, no real product screens
  (login, weight entry, dashboard, etc.), and no mockup/sign-off step
  are required for this feature — the placeholder is not a designed
  product screen, it's the smallest real thing TDD can anchor to,
  same category as the backend's `/health` endpoint (003). The first
  actual product screen gets a mockup when that feature starts.
- **AC-5**: `client/README.md` and `docs/architecture.md` are updated
  to reflect the client's initialized (not skeleton) status.
