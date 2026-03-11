# Guidelines

This file is the behavioral contract for the template.

## CLI Behavior

- Default to machine-readable JSON for structured commands.
- Offer `--text` for human-readable output when it materially improves terminal use.
- Keep success output quiet and predictable.
- Send errors to stderr and return non-zero exit codes.
- Include examples in `--help` for any command that is not obvious.
- Ship shell completion from day one.

## Command Design

- Keep command construction in `internal/cli/`.
- Prefer one file per command or tightly related command group.
- Use `RunE` so errors propagate cleanly.
- Avoid global state outside the injected runtime.
- Add new dependencies through `internal/app.Runtime` or explicit constructors so tests stay easy.

## Config Design

- Use XDG-aware config paths by default.
- Expose a command that tells users where config lives.
- If you persist secrets or tokens later, use restrictive file permissions.
- Document every environment-variable override in the README once they exist.

## Repository Hygiene

- `README.md` explains the current public contract.
- `docs/status.md` tracks what is true right now.
- `docs/roadmap.md` tracks upcoming work.
- `CHANGELOG.md` tracks notable changes.
- `AGENTS.md` tells coding agents how to behave inside the repo.
- Commits use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) headers plus a short explanatory body.

## Quality Bar

- Add tests for user-visible behavior.
- Run `make fmt test build` before merge.
- Keep lint passing in CI.
- Keep release paths current when binary names or entrypoints change.

## Commit Convention

- Use `<type>[optional scope]: <description>` for the subject line.
- Prefer standard types such as `feat`, `fix`, `docs`, `refactor`, `test`, `ci`, `build`, and `chore`.
- Include a blank line after the subject.
- Include a body of 1-3 concise sentences explaining the reason for the change or giving a very brief summary of the change.
- Use `BREAKING CHANGE:` in the footer when a change intentionally breaks compatibility.

## Primary References

- [clig.dev](https://clig.dev/)
- [Conventional Commits 1.0.0](https://www.conventionalcommits.org/en/v1.0.0/)
- [Go module layout](https://go.dev/doc/modules/layout)
- [Go testing tutorial](https://go.dev/doc/tutorial/add-a-test)
- [Cobra command structure](https://cobra.dev/docs/how-to-guides/working-with-commands/)
- [Cobra shell completion](https://cobra.dev/docs/how-to-guides/shell-completion/)
- [GitHub template repositories](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-template-repository)
- [GoReleaser GitHub Actions](https://goreleaser.com/ci/actions/)
