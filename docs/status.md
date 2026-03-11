# Status

Last updated: 2026-03-11

## Goal

Provide a reusable GitHub template for starting Go CLI applications without redoing command wiring, docs scaffolding, CI, or release automation.

## Implemented

- Modular Cobra CLI scaffold under `cmd/` and `internal/`.
- Build metadata injection via ldflags.
- Example user-facing commands:
  - `hello`
  - `doctor`
  - `config path|template|init`
  - `completion`
  - `version`
- XDG-style config path resolution and config template initialization.
- Unit tests for command behavior and path helpers.
- `Makefile`, `.golangci.yml`, CI workflow, release workflow, and GoReleaser config.
- Repo docs:
  - `AGENTS.md`
  - `README.md`
  - `docs/guidelines.md`
  - `docs/architecture.md`
  - `docs/roadmap.md`
  - `CHANGELOG.md`
  - `RELEASE_SETUP.md`

## Current Command Set

```text
gocli hello [name]
gocli doctor
gocli config path|template|init
gocli completion bash|zsh|fish|powershell
gocli version
```

## Verification Commands

```bash
go mod tidy
make fmt
make test
make build
./bin/gocli doctor --text
./bin/gocli hello Toto
```

## Known Follow-Up Work For Derived Projects

- Rename `gocli` everywhere.
- Replace the sample `hello` command with the first real domain command.
- Add real config parsing once the product schema exists.
- Add package-manager publishing only after the repo identity is stable.
