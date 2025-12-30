package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/verifiable-labs/agent-linter/internal/config"
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

			fmt.Fprintf(os.Stdout, "Loaded config: %s\n", cfgFile)
			fmt.Fprintf(os.Stdout, "Definitions: %d\n", len(cfg.Inputs.Definitions))
			fmt.Fprintf(os.Stdout, "Invocations: %d\n", len(cfg.Inputs.Invocations))
			fmt.Fprintf(os.Stdout, "Wiring: %d\n", len(cfg.Inputs.Wiring))

			return nil
		},
	}

	cmd.Flags().StringVar(&format, "format", "human", "Output format: human, json, sarif")
	cmd.Flags().StringVar(&configPath, "config", "", "Path to agent-linter config file")

	return cmd
}
