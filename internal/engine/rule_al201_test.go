package engine

import "testing"

func TestAL201ProdWiring(t *testing.T) {
	in := Inputs{
		Wiring: []WiringBinding{
			{
				Action:      "create_user",
				Environment: "prod",
				Target:      "user-service",
			},
		},
	}

	res := Run(in, []Rule{RuleAL201NoProdByDefault{}}, nil)
	if len(res.Findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(res.Findings))
	}
	if res.Findings[0].RuleID != "AL201" {
		t.Fatalf("expected AL201, got %s", res.Findings[0].RuleID)
	}
}
