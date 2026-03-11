# Go CLI Template Agent Notes

This repository is a starter template for production-minded Go CLIs. Treat the code, docs, and automation here as the baseline contract unless the user explicitly changes it.

## Read First

Before making changes, review:

1. `README.md`
2. `docs/guidelines.md`
3. `docs/status.md`
4. `docs/roadmap.md`
5. `docs/architecture.md`

## External References

Use these resources when shaping behavior and repo decisions:

- [Command Line Interface Guidelines](https://clig.dev/)
- [Conventional Commits 1.0.0](https://www.conventionalcommits.org/en/v1.0.0/)
- [Go module layout](https://go.dev/doc/modules/layout)
- [Go testing tutorial](https://go.dev/doc/tutorial/add-a-test)
- [Cobra command structure](https://cobra.dev/docs/how-to-guides/working-with-commands/)
- [Cobra shell completion](https://cobra.dev/docs/how-to-guides/shell-completion/)
- [GoReleaser GitHub Actions docs](https://goreleaser.com/ci/actions/)

## Standing Rules

1. Follow `clig.dev` defaults: quiet success, clear stderr on failure, useful `--help`, and script-friendly stdout.
2. Keep the CLI JSON-first unless a command intentionally emits raw text or generated shell/config content.
3. Add or update tests for every behavior change. Run `make fmt test build` before handing work back.
4. Keep command construction modular under `internal/cli/`; do not collapse the template back into one huge `main.go`.
5. Preserve `cmd/` for binaries and `internal/` for non-public packages. Only add `pkg/` if the repo intentionally exposes a reusable Go API.
6. Update `docs/status.md`, `docs/roadmap.md`, `README.md`, and `CHANGELOG.md` whenever the user-facing contract changes.
7. Keep release automation working. If you change build paths, binary names, or version wiring, update `Makefile`, `.goreleaser.yml`, and the GitHub workflows in the same change.
8. Treat `hello` as a disposable example. Replace it early in real projects, but keep at least one tested feature command in the template at all times.
9. Prefer XDG-aware config locations and explicit config path discovery over hidden magic.
10. Use Conventional Commits for every commit header, for example `feat(cli): add doctor command` or `fix(config): handle missing config file`.
11. Every commit must include a short body after the subject line. Write 1-3 concise sentences explaining why the change exists or giving a very brief summary of the change.

## Commit Ceremony

Before creating a commit:

1. Review the staged diff and ensure the commit is one logical change.
2. Run `make fmt test build`. If `golangci-lint` is installed, run `make lint` too.
3. Update docs that changed the public contract, workflow, or repo rules.

When writing the commit:

1. Use a Conventional Commits header: `<type>[optional scope]: <description>`.
2. Add a blank line.
3. Add a short body of 1-3 sentences. The body must explain the reason for the change or give a very brief summary of what changed.
4. Add a footer only when needed, for example `BREAKING CHANGE:` or issue references.

Example:

```text
feat(cli): add doctor command

Add a built-in environment sanity check for new projects created from the template.
This gives derived CLIs a tested diagnostics command from day one.
```

## Release Ritual

When preparing a release:

1. Confirm `make fmt test build` passes locally.
2. Review `docs/status.md` and `CHANGELOG.md`.
3. Tag with `vX.Y.Z`.
4. Push the tag and watch `.github/workflows/release.yml`.
5. Verify GitHub release assets and any downstream packaging you enabled for the derived project.
