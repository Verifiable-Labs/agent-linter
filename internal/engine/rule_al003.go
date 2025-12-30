package engine

import "fmt"

type RuleAL003DisallowPermissiveParamTypes struct{}

func (r RuleAL003DisallowPermissiveParamTypes) ID() string { return "AL003" }
func (r RuleAL003DisallowPermissiveParamTypes) DefaultSeverity() Severity {
	return SeverityWarning
}

func (r RuleAL003DisallowPermissiveParamTypes) Apply(in Inputs) []Finding {
	var findings []Finding

	for _, d := range in.Definitions {
		for name, p := range d.Parameters {
			if p.Type == "any" || p.Type == "object" {
				findings = append(findings, Finding{
					RuleID:      r.ID(),
					Severity:    r.DefaultSeverity(),
					Message:     fmt.Sprintf("action %q parameter %q has overly permissive type %q", d.Name, name, p.Type),
					Remediation: "Use explicit parameter types or define structured object properties.",
				})
			}
		}
	}

	return findings
}
