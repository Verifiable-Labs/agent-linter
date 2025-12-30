package engine

import (
	"fmt"
	"sort"
	"strings"
)

type RuleAL102InvocationMissingRequiredParams struct{}

func (r RuleAL102InvocationMissingRequiredParams) ID() string { return "AL102" }
func (r RuleAL102InvocationMissingRequiredParams) DefaultSeverity() Severity {
	return SeverityError
}

func (r RuleAL102InvocationMissingRequiredParams) Apply(in Inputs) []Finding {
	defByName := make(map[string]ActionDefinition)
	for _, d := range in.Definitions {
		defByName[d.Name] = d
	}

	var findings []Finding
	for _, inv := range in.Invocations {
		def, ok := defByName[inv.Action]
		if !ok {
			continue
		}

		var missing []string
		for paramName, param := range def.Parameters {
			if !param.Required {
				continue
			}
			if inv.Args == nil {
				missing = append(missing, paramName)
				continue
			}
			if _, ok := inv.Args[paramName]; !ok {
				missing = append(missing, paramName)
			}
		}

		if len(missing) > 0 {
			sort.Strings(missing)
			findings = append(findings, Finding{
				RuleID:      r.ID(),
				Severity:    r.DefaultSeverity(),
				Message:     fmt.Sprintf("invocation of %q missing required params: %s", inv.Action, strings.Join(missing, ", ")),
				Remediation: "Add the missing required parameters to the invocation args.",
				File:        inv.Source,
			})
		}
	}

	return findings
}
