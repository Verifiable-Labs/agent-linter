package app

import (
	"time"

	"github.com/verifiable-labs/agent-linter/internal/config"
	"github.com/verifiable-labs/agent-linter/internal/engine"
)

func applySuppressions(findings []engine.Finding, suppress []config.Suppression, now time.Time) []engine.Finding {
	if len(suppress) == 0 || len(findings) == 0 {
		return findings
	}

	index := make(map[string]config.Suppression)
	for _, s := range suppress {
		key := s.Rule + "|" + s.Fingerprint
		index[key] = s
	}

	out := make([]engine.Finding, 0, len(findings))
	for _, f := range findings {
		key := f.RuleID + "|" + f.Fingerprint
		s, ok := index[key]
		if !ok {
			out = append(out, f)
			continue
		}

		// If expires is empty, suppression is active forever.
		if s.Expires == "" {
			continue
		}

		// Expires format: YYYY-MM-DD
		exp, err := time.Parse("2006-01-02", s.Expires)
		if err != nil {
			// Invalid expiration means do not suppress.
			out = append(out, f)
			continue
		}

		// If now is after expiration date, do not suppress.
		if now.After(exp) {
			out = append(out, f)
			continue
		}

		// Otherwise suppress.
	}

	return out
}
