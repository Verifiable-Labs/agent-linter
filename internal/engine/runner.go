package engine

import "sort"

type Result struct {
	Findings []Finding
}

func Run(in Inputs, rules []Rule) Result {
	var out []Finding

	for _, rule := range rules {
		out = append(out, rule.Apply(in)...)
	}

	sort.Slice(out, func(i, j int) bool {
		if out[i].RuleID != out[j].RuleID {
			return out[i].RuleID < out[j].RuleID
		}
		if out[i].Severity != out[j].Severity {
			return out[i].Severity < out[j].Severity
		}
		return out[i].Message < out[j].Message
	})

	return Result{Findings: out}
}

func (r Result) HasErrors() bool {
	for _, f := range r.Findings {
		if f.Severity == SeverityError {
			return true
		}
	}
	return false
}
