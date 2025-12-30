package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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

			_ = configPath
			_ = args

			fmt.Fprintln(os.Stdout, "lint stub: Step 1 complete. Next step will load config, discover inputs, and run rules.")
			return nil
		},
	}

	cmd.Flags().StringVar(&format, "format", "human", "Output format: human, json, sarif")
	cmd.Flags().StringVar(&configPath, "config", "", "Path to agent-linter config file (default: discover agent-linter.yaml)")

	return cmd
}
