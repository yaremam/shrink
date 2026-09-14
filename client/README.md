# Shrink! — Client

The Flutter client for Shrink!, covering mobile (iOS/Android) and web
from a single codebase.

This is currently a skeleton. Project initialization (`flutter create`
and everything that follows) is separate future work, tracked as its
own feature per the project's TDD process — see
[`docs/tdr/002_tech_stack_design.md`](../docs/tdr/002_tech_stack_design.md)
for the stack decision behind this.

- Auth: Firebase Auth (email/password + Google/Apple sign-in)
- Data entry (e.g. weight logging) is offline-first with sync
