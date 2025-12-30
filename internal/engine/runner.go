package engine

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
)


type Result struct {
	Findings []Finding
}

func fingerprintFinding(f Finding) string {
	h := sha256.New()

	// Keep fingerprint stable across formatting changes.
	// Use only semantic fields.
	h.Write([]byte(f.RuleID))
	h.Write([]byte{0})
	h.Write([]byte(string(f.Severity)))
	h.Write([]byte{0})
	h.Write([]byte(f.File))
	h.Write([]byte{0})
	h.Write([]byte(f.Message))

	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}


func Run(in Inputs, rules []Rule, enabled map[string]bool) Result {
	var out []Finding

	for _, rule := range rules {
		if enabled != nil {
			if ok, exists := enabled[rule.ID()]; exists && !ok {
				continue
			}
		}
		out = append(out, rule.Apply(in)...)
	}

	for i := range out {
		out[i].Fingerprint = fingerprintFinding(out[i])
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
