package engine

import "testing"

func TestAL101UnknownAction(t *testing.T) {
	in := Inputs{
		Definitions: []ActionDefinition{
			{Name: "create_user"},
		},
		Invocations: []ActionInvocation{
			{Action: "delete_user", Args: map[string]any{}, Source: "inv.jsonl"},
		},
	}

	res := Run(in, []Rule{RuleAL101InvocationMustReferenceKnownAction{}}, nil)
	if len(res.Findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(res.Findings))
	}
	if res.Findings[0].RuleID != "AL101" {
		t.Fatalf("expected AL101, got %s", res.Findings[0].RuleID)
	}
}
