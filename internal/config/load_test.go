package config

import "testing"

func TestLoadValidConfig(t *testing.T) {
	cfg, err := Load("testdata/valid.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Version != 1 {
		t.Fatalf("expected version 1, got %d", cfg.Version)
	}
}
