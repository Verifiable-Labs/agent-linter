package config

import "fmt"

func validate(cfg *Config) error {
	if cfg.Version != 1 {
		return fmt.Errorf("unsupported config version %d (only version 1 is supported)", cfg.Version)
	}

	if len(cfg.Inputs.Definitions) == 0 &&
		len(cfg.Inputs.Invocations) == 0 &&
		len(cfg.Inputs.Wiring) == 0 {
		return fmt.Errorf("config must specify at least one input (definitions, invocations, or wiring)")
	}

	for ruleID, rule := range cfg.Rules {
		if rule.Severity != "error" && rule.Severity != "warning" {
			return fmt.Errorf("rule %s has invalid severity %q", ruleID, rule.Severity)
		}
	}

	return nil
}
