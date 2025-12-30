package engine

import "fmt"

type RuleAL201NoProdByDefault struct{}

func (r RuleAL201NoProdByDefault) ID() string { return "AL201" }
func (r RuleAL201NoProdByDefault) DefaultSeverity() Severity {
	return SeverityError
}

func (r RuleAL201NoProdByDefault) Apply(in Inputs) []Finding {
	var findings []Finding

	for _, w := range in.Wiring {
		if w.Environment == "prod" || w.Environment == "production" {
			findings = append(findings, Finding{
				RuleID:      r.ID(),
				Severity:    r.DefaultSeverity(),
				Message:     fmt.Sprintf("action %q wired to production target %q by default", w.Action, w.Target),
				Remediation: "Do not bind production environments by default. Require explicit environment selection.",
			})
		}
	}

	return findings
}
