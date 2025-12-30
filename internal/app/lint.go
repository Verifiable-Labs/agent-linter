package app

import (
	"fmt"
	"os"
	"path/filepath"

	"encoding/json"
	"time"

	"github.com/spf13/cobra"
	"github.com/verifiable-labs/agent-linter/internal/config"
	"github.com/verifiable-labs/agent-linter/internal/engine"
)

func newLintCmd() *cobra.Command {
	var format string
	var configPath string

	cmd := &cobra.Command{
		Use:   "lint [paths...]",
		Short: "Lint an agent project (definitions, invocations, wiring)",
		Args:  cobra.ArbitraryArgs,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if format != "human" && format != "json" && format != "sarif" {
				return fmt.Errorf("unsupported format %q (use human, json, sarif)", format)
			}

			cfgFile := configPath
			if cfgFile == "" {
				cfgFile = "agent-linter.yaml"
			}

			cfgFile, err := filepath.Abs(cfgFile)
			if err != nil {
				return err
			}

			cfg, err := config.Load(cfgFile)
			if err != nil {
				return err
			}

			enabled := map[string]bool{}
			settings := map[string]engine.RuleSetting{}

			for ruleID, r := range cfg.Rules {
				enabled[ruleID] = r.Enabled

				var sev engine.Severity
				if r.Severity == "error" {
					sev = engine.SeverityError
				} else if r.Severity == "warning" {
					sev = engine.SeverityWarning
				}

				settings[ruleID] = engine.RuleSetting{
					Enabled:  r.Enabled,
					Severity: sev,
					HasValue: true,
				}
			}

			inputs := engine.Inputs{}

			for _, p := range cfg.Inputs.Definitions {
				defs, err := engine.LoadActionDefinitions(p)
				if err != nil {
					return err
				}
				inputs.Definitions = append(inputs.Definitions, defs...)
			}

			for _, p := range cfg.Inputs.Invocations {
				inv, err := engine.LoadActionInvocations(p)
				if err != nil {
					return err
				}
				inputs.Invocations = append(inputs.Invocations, inv...)
			}

			for _, p := range cfg.Inputs.Wiring {
				w, err := engine.LoadWiring(p)
				if err != nil {
					return err
				}
				inputs.Wiring = append(inputs.Wiring, w...)
			}

			rules := engine.DefaultRules()
			result := engine.Run(inputs, rules, enabled)
			result.Findings = engine.ApplyRuleSettings(result.Findings, settings)
			result.Findings = applySuppressions(result.Findings, cfg.Suppress, time.Now())

			if format == "sarif" {
				log := toSarif(result.Findings, Version)
				enc := json.NewEncoder(os.Stdout)
				enc.SetIndent("", "  ")
				if err := enc.Encode(log); err != nil {
					return err
				}
			} else if format == "json" {
				enc := json.NewEncoder(os.Stdout)
				enc.SetIndent("", "  ")

				out := Output{
					Version:  1,
					Findings: result.Findings,
				}
				if err := enc.Encode(out); err != nil {
					return err
				}
			} else {
				if len(result.Findings) == 0 {
					fmt.Fprintln(os.Stdout, "No findings.")
				} else {
					for _, f := range result.Findings {
						fmt.Fprintf(os.Stdout, "[%s] %s: %s\n", f.Severity, f.RuleID, f.Message)
						fmt.Fprintf(os.Stdout, "Remediation: %s\n", f.Remediation)
					}
				}
			}

			if result.HasErrors() {
				return fmt.Errorf("%w", ErrLintFailed)
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&format, "format", "human", "Output format: human, json, sarif")
	cmd.Flags().StringVar(&configPath, "config", "", "Path to agent-linter config file")

	return cmd
}
