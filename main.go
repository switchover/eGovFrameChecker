package main

import (
	"fmt"
	"os"

	"github.com/switchover/eGovFrameChecker/cmd"
	"github.com/switchover/eGovFrameChecker/pkg/terminal"
)

func main() {
	terminal.EnableANSI()

	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
		os.Exit(1)
	}
}
