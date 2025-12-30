package app

import "testing"

func TestRootCmdBuilds(t *testing.T) {
	cmd := NewRootCmd()
	if cmd == nil {
		t.Fatalf("expected root command, got nil")
	}
}
