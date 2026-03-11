# Go CLI Template

[![CI](https://img.shields.io/github/actions/workflow/status/toto/cli-go-template/ci.yml?branch=main&label=CI)](https://github.com/toto/cli-go-template/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/toto/cli-go-template)](https://go.dev/)
[![License](https://img.shields.io/github/license/toto/cli-go-template)](LICENSE)

Opinionated starter template for a Go CLI application with:

- Modular Cobra command layout
- JSON-first output with optional `--text`
- XDG-style config path handling
- Tests, linting, CI, and GoReleaser wiring
- `AGENTS.md`, guidelines, status, roadmap, and architecture docs

It is intentionally generic. Rename `gocli`, the module path, and the sample `hello` command when you turn it into a real product.

## Included Command Set

```text
gocli hello [name]
gocli doctor
gocli config path|template|init
gocli completion bash|zsh|fish|powershell
gocli version
```

## Quick Start

```bash
# Install dependencies and generate go.sum
go mod tidy

# Run tests
make test

# Build the binary
make build

# Inspect the local runtime
./bin/gocli doctor --text
```

## Template Customization Checklist

When you create a real CLI from this template:

1. Rename the binary under `cmd/gocli/`, `Makefile`, `.goreleaser.yml`, and docs.
2. Change the module path in `go.mod`.
3. Replace the sample `hello` command with your first real domain command.
4. Update `README.md`, `docs/status.md`, `docs/roadmap.md`, and `CHANGELOG.md`.
5. Adjust release naming, badges, and repository URLs.
6. Decide whether you need package-manager publishing such as Homebrew.

## Development Commands

```bash
make fmt
make test
make test-race
make build
make install
make lint
make release-snapshot
```

## Project Structure

```text
cmd/gocli/          binary entrypoint
internal/app/       runtime dependencies and build info
internal/cli/       Cobra commands
internal/config/    starter config helpers
internal/output/    JSON/text helpers
internal/paths/     XDG-aware path resolution
docs/               repo contract and planning docs
.github/workflows/  CI and release automation
```

## Documentation

- [AGENTS.md](AGENTS.md)
- [Contribution Guide](CONTRIBUTING.md)
- [Architecture](docs/architecture.md)
- [Guidelines](docs/guidelines.md)
- [Status](docs/status.md)
- [Roadmap](docs/roadmap.md)
- [Release Setup](RELEASE_SETUP.md)

## Why This Shape

The layout is derived from the working structure in your `whoopy`, `stamp`, and `withingy` repos, but tightened into a cleaner template:

- `whoopy` and `withingy` contributed the modular `cmd/` + `internal/` split, agent guidance, and status ledger pattern.
- `stamp` contributed the stronger CI/lint/release setup.
- Current official guidance from Go, Cobra, GoReleaser, GitHub template docs, and [clig.dev](https://clig.dev/) informed the final defaults.

## License

MIT
