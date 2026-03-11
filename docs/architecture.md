# Architecture

## Goals

- Keep the binary entrypoint minimal.
- Keep Cobra command setup separate from business logic and runtime dependencies.
- Keep the template safe for automation, scripting, and future growth.

## Layers

### `cmd/gocli`

Owns process startup, build metadata injection, and exit-code handling.

### `internal/app`

Holds runtime dependencies such as stdout, stderr, current time, config directory resolution, and build metadata. This keeps commands testable.

### `internal/cli`

Owns Cobra command construction. Each command lives in its own file and receives dependencies through `app.Runtime`.

### `internal/config`

Contains starter config location and initialization helpers. Replace the scaffolded schema once the real application shape is known.

### `internal/output`

Keeps JSON/text rendering logic trivial and consistent.

### `internal/paths`

Resolves XDG-style config paths without mixing OS path logic into commands.

## Testing Strategy

- Command tests should instantiate `app.Runtime` with in-memory buffers and temporary directories.
- Pure helpers should have focused unit tests in their own package.
- Keep live network or filesystem integration out of the default test suite unless it is fully controlled and reproducible.
