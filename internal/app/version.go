package app

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	Version   = "dev"
	GitCommit = "unknown"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print agent-linter version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("agent-linter %s\n", Version)
			fmt.Printf("commit: %s\n", GitCommit)
			fmt.Printf("go: %s\n", runtime.Version())
			fmt.Printf("os: %s\n", runtime.GOOS)
			fmt.Printf("arch: %s\n", runtime.GOARCH)
		},
	}
}
