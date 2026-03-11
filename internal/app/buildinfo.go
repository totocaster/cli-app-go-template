package app

// BuildInfo is injected at build time so commands can expose reproducible version output.
type BuildInfo struct {
	Binary  string
	Version string
	Commit  string
	Date    string
}
