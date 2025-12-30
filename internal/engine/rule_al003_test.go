package engine

import "testing"

func TestAL003PermissiveParamType(t *testing.T) {
	in := Inputs{
		Definitions: []ActionDefinition{
			{
				Name: "x",
				Parameters: map[string]Parameter{
					"payload": {Type: "any", Required: true},
				},
			},
		},
	}

	res := Run(in, []Rule{RuleAL003DisallowPermissiveParamTypes{}}, nil)
	if len(res.Findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(res.Findings))
	}
	if res.Findings[0].RuleID != "AL003" {
		t.Fatalf("expected AL003, got %s", res.Findings[0].RuleID)
	}
}
