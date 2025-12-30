package app

import "github.com/verifiable-labs/agent-linter/internal/engine"

type Output struct {
	Version  int              `json:"version"`
	Findings []engine.Finding `json:"findings"`
}
