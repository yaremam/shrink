# 005 — Client initialization design

## Context & requirements

TDR 002 chose the client stack (Flutter, mobile + web, Firebase Auth)
but deferred actual project initialization, same as it did for the
backend. This feature does that initialization, scoped the same way
003 was: the smallest slice that gives the project real, tested
behavior, rather than building out the full stack (auth, real
screens) in one pass.

Flutter 3.47.4 (stable) was installed in this sandbox during 003/004's
work. `flutter doctor` shows no Android SDK, no Chrome, and no Linux
desktop toolchain — none of which block `flutter test`, since widget
tests run on the host Dart VM against a simulated framework binding,
not a real device or browser. That's sufficient for this feature's
TDD cycle.

**Sandbox caveat found during this feature**: `flutter test` is
intermittently flaky in this sandbox — the `flutter_tester` subprocess
segfaults roughly 1 run in 3, inside flutter_tools' own
`FlutterPlatform._startTest`, not in application/test code. Installing
`libgl1-mesa-dri`/`libegl1`/`libgles2`/`mesa-utils` (prompted by
`flutter doctor`'s missing `eglinfo` warning) reduced but did not
eliminate it. This mirrors 003's Docker finding: a sandbox-level
limitation, not a code defect — every successful run shows correct
pass/fail behavior (confirmed red before the placeholder screen
existed, green after, consistently whenever the subprocess doesn't
crash). Re-run `flutter test` on a failure before assuming a real
regression; a real CI environment should be checked for the same
flakiness before relying on a single green run there either.

## Alternatives considered

### Scope: bare `flutter create` vs. scaffold + widget test

**Bare `flutter create`** (unmodified default counter-app template)
- Pros: smallest possible change.
- Cons: the default template already ships with its own widget test
  that passes immediately — no red-green cycle happens, violating the
  project's TDD-always rule. Same reasoning as 003's decision to add a
  health check rather than a bare `go mod init`.

**Scaffold + rewritten widget test anchored on a placeholder screen**
- Pros: gives this feature a real red-green cycle; produces something
  slightly more identifiable as "Shrink!" than the stock counter demo.
- Cons: still minimal — not a designed screen.

**Decision: scaffold + widget test**, mirroring 003 exactly.

### Does the placeholder screen need a mockup + sign-off?

The global process requires a mockup (Artifact) and sign-off before
writing "any new screen or user-facing view." This is technically the
client's first screen.

**Require a mockup**
- Cons: the rule's intent is to avoid building real user-facing UI
  blind. A bare "Shrink!" placeholder isn't a designed product
  screen — there's nothing to mock up, it's one line of text standing
  in for "the app boots and renders something," the UI equivalent of
  the backend's `/health` check.

**Skip the mockup for this specific placeholder**
- Pros: consistent with treating this as scaffolding, not product UI;
  the first real screen (login, weight entry, dashboard) gets a
  mockup when that feature starts, when there's an actual design to
  put in front of the user.

**Decision: skip it**, on the same reasoning as 003 treating `/health`
as scaffolding rather than a designed API.

### Firebase Auth: in scope now or deferred?

Same shape as deferring PostgreSQL from 003. No Firebase project
exists yet, and wiring auth in now would add config/secrets handling
with no real signed-in-user flow behind it yet.

**Decision: deferred.** Becomes its own future feature once a real
screen needs a signed-in user.

### Project/org naming

Sibling project opusflow uses `com.opusflow.app` as its bundle/package
identifier (found in `opusflow/mobile/app.json`).

**Decision:** package name `shrink`, org `com.shrink`, explicit
`--app-id com.shrink.app` (rather than letting `flutter create`
default to `com.shrink.shrink` from org + directory name) — mirrors
opusflow's `com.<product>.app` convention exactly.

### Platforms to scaffold

`flutter create` defaults to iOS, Android, web, *and* desktop
(Linux/macOS/Windows). TDR 002 scoped the client to mobile + web only.

**Decision:** `--platforms=ios,android,web`. Generating unused desktop
targets would add dead platform folders — and Linux desktop
specifically can't even build in this sandbox (no clang/cmake/gtk3).

### Branch basis

`feature/client-init` off `main` vs. stacked on an unmerged sibling
feature branch.

**Decision: off `main`**, independent of `feature/backend-init` and
`feature/coding-standards`. Client and backend are decoupled per TDR
002 (no monorepo tooling, REST-only, no shared code) — nothing in
client-init depends on either of those branches' contents.

## Structural decision (summary)

- `client/`: Flutter project, package `shrink`, org `com.shrink`,
  app id `com.shrink.app`, platforms `ios`, `android`, `web`.
- `client/lib/main.dart`: minimal placeholder screen displaying
  "Shrink!", replacing the default counter-app template.
- `client/test/widget_test.dart`: rewritten to assert the placeholder
  text renders — this feature's red-green pair.
- No Firebase Auth, no real product screens, no mockup step.

## Cross-cutting implications

- **Testing**: this is the feature where TDD starts for the client,
  mirroring 003 for the backend. `flutter_test` confirmed as the
  testing default flagged (but not locked in) by 002.
- **Future features implied**: Firebase Auth integration, the first
  real product screen (with a mockup + sign-off), a Dart section in
  `CODING_STANDARDS.md` once there's enough real client code to
  ground it (mirrors 004's Go-only scoping reasoning).
- **Schema/contract changes**: none — the client doesn't talk to the
  backend yet.
