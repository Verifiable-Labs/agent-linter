package engine

import "fmt"

type RuleAL001DuplicateActionNames struct{}

func (r RuleAL001DuplicateActionNames) ID() string { return "AL001" }

func (r RuleAL001DuplicateActionNames) Apply(in Inputs) []Finding {
	seen := make(map[string]int)
	var findings []Finding

	for _, d := range in.Definitions {
		seen[d.Name]++
	}

	for name, count := range seen {
		if count > 1 {
			findings = append(findings, Finding{
				RuleID:      r.ID(),
				Severity:    SeverityError,
				Message:     fmt.Sprintf("duplicate action name %q appears %d times", name, count),
				Remediation: "Rename or remove duplicates so every action name is unique.",
			})
		}
	}

	return findings
}
