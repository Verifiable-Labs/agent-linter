package app

import (
	"github.com/spf13/cobra"
	
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agent-linter",
		Short: "agent-linter is a deterministic static analysis tool for executable agent actions",
		Long:  "agent-linter checks action definitions, action invocations, and execution wiring before agents run.",
	}

	cmd.AddCommand(newVersionCmd())
	cmd.AddCommand(newLintCmd())
	cmd.AddCommand(newRulesCmd())


	return cmd
}
