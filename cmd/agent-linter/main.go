package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/verifiable-labs/agent-linter/internal/app"
)

func main() {
	err := app.NewRootCmd().Execute()
	if err == nil {
		os.Exit(0)
	}

	if errors.Is(err, app.ErrLintFailed) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}
