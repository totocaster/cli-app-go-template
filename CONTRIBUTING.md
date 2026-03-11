# Contributing

## Development Loop

1. Read `AGENTS.md` and the files under `docs/`.
2. Keep changes small and focused.
3. Run `make fmt test build`.
4. Update docs if the CLI contract or repo conventions changed.

## Code Standards

- Keep commands modular in `internal/cli/`.
- Prefer JSON output by default and `--text` for human-readable summaries.
- Write clear, stable help text with at least one example for non-trivial commands.
- Use stderr for failures and non-zero exit codes for actionable problems.
- Use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for commit subjects.
- Every commit should include a short body of 1-3 concise sentences explaining why the change was made or summarizing the change.

## Tests

- Add unit tests for new command behavior and helper packages.
- Keep commands dependency-injected through `internal/app.Runtime` so tests do not need real user state.
- Use table-driven tests where they improve clarity.

## Commit Format

Use this shape for commits:

```text
<type>[optional scope]: <description>

<1-3 short sentences describing why the change exists or what changed>
```

Examples:

```text
feat(config): add config init command

Add a starter config writer so new projects can expose a real config path immediately.
This keeps the template useful before application-specific config parsing exists.
```

```text
fix(cli): return exit code 2 for existing config file

Make repeated config initialization fail with a specific non-zero exit code.
This gives scripts a predictable way to detect the already-exists case.
```

## Releases

Release instructions live in `RELEASE_SETUP.md`.
