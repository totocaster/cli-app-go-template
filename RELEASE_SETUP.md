# Release Setup

This template includes GitHub Actions plus GoReleaser so a derived project can ship tagged releases quickly.

## Before First Release

1. Rename the binary and update:
   - `cmd/gocli/`
   - `Makefile`
   - `.goreleaser.yml`
   - `README.md`
2. Update the module path in `go.mod`.
3. Replace placeholder repository URLs and badges.
4. Decide whether to add downstream packaging such as Homebrew, Scoop, or an `asdf` plugin.

## Cut a Release

```bash
git status
make fmt test build
git tag -a v0.1.0 -m "Initial release"
git push origin v0.1.0
```

The release workflow at `.github/workflows/release.yml` publishes artifacts automatically.

## Optional Packaging

The included `.goreleaser.yml` handles GitHub release assets only. If the derived project needs Homebrew or another package manager, add that block after the project identity is finalized.
