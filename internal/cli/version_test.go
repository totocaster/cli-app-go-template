package cli

import (
	"encoding/json"
	"testing"
)

func TestVersionJSON(t *testing.T) {
	t.Parallel()

	rt := newTestRuntime(t)
	cmd := newRootCmd(rt.runtime)
	cmd.SetArgs([]string{"version"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute returned error: %v", err)
	}

	var response versionResponse
	if err := json.Unmarshal(rt.stdout.Bytes(), &response); err != nil {
		t.Fatalf("failed to parse JSON output: %v", err)
	}

	if response.Version != "test" {
		t.Fatalf("unexpected version %q", response.Version)
	}
}
