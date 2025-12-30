package engine

import "testing"

func TestAL001DuplicateActionNames(t *testing.T) {
	in := Inputs{
		Definitions: []ActionDefinition{
			{Name: "x"},
			{Name: "x"},
		},
	}

	res := Run(in, []Rule{RuleAL001DuplicateActionNames{}})
	if len(res.Findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(res.Findings))
	}

	f := res.Findings[0]
	if f.RuleID != "AL001" {
		t.Fatalf("expected AL001, got %s", f.RuleID)
	}
	if f.Severity != SeverityError {
		t.Fatalf("expected error severity, got %s", f.Severity)
	}
}
