package cli

import (
	"encoding/json"
	"testing"
)

func TestDoctorJSON(t *testing.T) {
	t.Parallel()

	rt := newTestRuntime(t)
	cmd := newRootCmd(rt.runtime)
	cmd.SetArgs([]string{"doctor"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute returned error: %v", err)
	}

	var response doctorResponse
	if err := json.Unmarshal(rt.stdout.Bytes(), &response); err != nil {
		t.Fatalf("failed to parse JSON output: %v", err)
	}

	if response.ConfigExists {
		t.Fatal("expected config file to be absent")
	}
}
