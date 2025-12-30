package engine

type RuleSetting struct {
	Enabled  bool
	Severity Severity
	HasValue bool
}

func ApplyRuleSettings(findings []Finding, settings map[string]RuleSetting) []Finding {
	out := make([]Finding, 0, len(findings))
	for _, f := range findings {
		s, ok := settings[f.RuleID]
		if ok {
			if !s.Enabled {
				continue
			}
			if s.HasValue && s.Severity != "" {
				f.Severity = s.Severity
			}
		}
		out = append(out, f)
	}
	return out
}
