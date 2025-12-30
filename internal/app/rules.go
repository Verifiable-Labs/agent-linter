package app

import (
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
	"github.com/verifiable-labs/agent-linter/internal/engine"
)

func newRulesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "rules",
		Short: "List all available lint rules",
		Run: func(cmd *cobra.Command, args []string) {
			rules := engine.DefaultRules()
			sort.Slice(rules, func(i, j int) bool {
				return rules[i].ID() < rules[j].ID()
			})

			for _, r := range rules {
				fmt.Fprintf(
					os.Stdout,
					"%s [%s] %s\n",
					r.ID(),
					r.DefaultSeverity(),
					r.Description(),
				)
			}
		},
	}
}
