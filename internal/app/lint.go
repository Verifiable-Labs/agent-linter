package app

import (
	"fmt"
	"os"
	"path/filepath"

	"encoding/json"

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
			result := engine.Run(inputs, rules)

			if format == "json" {
				enc := json.NewEncoder(os.Stdout)
				enc.SetIndent("", "  ")
				if err := enc.Encode(result); err != nil {
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
