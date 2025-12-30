package app

import (
	"fmt"
	"os"
	"path/filepath"

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

			fmt.Fprintf(os.Stdout, "Loaded %d action definitions\n", len(inputs.Definitions))
			fmt.Fprintf(os.Stdout, "Loaded %d action invocations\n", len(inputs.Invocations))
			fmt.Fprintf(os.Stdout, "Loaded %d wiring bindings\n", len(inputs.Wiring))

			return nil
		},
	}

	cmd.Flags().StringVar(&format, "format", "human", "Output format: human, json, sarif")
	cmd.Flags().StringVar(&configPath, "config", "", "Path to agent-linter config file")

	return cmd
}
