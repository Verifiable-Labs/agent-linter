package engine

import "fmt"

type RuleAL101InvocationMustReferenceKnownAction struct{}

func (r RuleAL101InvocationMustReferenceKnownAction) ID() string { return "AL101" }
func (r RuleAL101InvocationMustReferenceKnownAction) DefaultSeverity() Severity {
	return SeverityError
}

func (r RuleAL101InvocationMustReferenceKnownAction) Description() string {
	return "Invocation must reference a known action"
}

func (r RuleAL101InvocationMustReferenceKnownAction) Apply(in Inputs) []Finding {
	known := make(map[string]bool)
	for _, d := range in.Definitions {
		known[d.Name] = true
	}

	var findings []Finding
	for _, inv := range in.Invocations {
		if inv.Action == "" {
			continue
		}
		if !known[inv.Action] {
			findings = append(findings, Finding{
				RuleID:      r.ID(),
				Severity:    r.DefaultSeverity(),
				Message:     fmt.Sprintf("invocation references unknown action %q", inv.Action),
				Remediation: "Define the action or fix the invocation to reference an existing action definition.",
				File:        inv.Source,
			})
		}
	}

	return findings
}
