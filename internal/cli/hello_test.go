package cli

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestHelloJSON(t *testing.T) {
	t.Parallel()

	rt := newTestRuntime(t)
	cmd := newRootCmd(rt.runtime)
	cmd.SetArgs([]string{"hello", "Toto"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute returned error: %v", err)
	}

	var response helloResponse
	if err := json.Unmarshal(rt.stdout.Bytes(), &response); err != nil {
		t.Fatalf("failed to parse JSON output: %v", err)
	}

	if response.Message != "hello, Toto" {
		t.Fatalf("unexpected message %q", response.Message)
	}
}

func TestHelloText(t *testing.T) {
	t.Parallel()

	rt := newTestRuntime(t)
	cmd := newRootCmd(rt.runtime)
	cmd.SetArgs([]string{"--text", "hello", "Toto", "--uppercase"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute returned error: %v", err)
	}

	if got := strings.TrimSpace(rt.stdout.String()); got != "HELLO, TOTO" {
		t.Fatalf("unexpected text output %q", got)
	}
}
