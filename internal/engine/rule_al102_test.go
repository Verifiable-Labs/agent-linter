package engine

import "testing"

func TestAL102MissingRequiredParams(t *testing.T) {
	in := Inputs{
		Definitions: []ActionDefinition{
			{
				Name: "create_user",
				Parameters: map[string]Parameter{
					"email": {Type: "string", Required: true},
					"admin": {Type: "boolean", Required: false},
				},
			},
		},
		Invocations: []ActionInvocation{
			{Action: "create_user", Args: map[string]any{"admin": true}, Source: "inv.jsonl"},
		},
	}

	res := Run(in, []Rule{RuleAL102InvocationMissingRequiredParams{}}, nil)
	if len(res.Findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(res.Findings))
	}
	if res.Findings[0].RuleID != "AL102" {
		t.Fatalf("expected AL102, got %s", res.Findings[0].RuleID)
	}
}
