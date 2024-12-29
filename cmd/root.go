package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "egovchecker",
	Short: "eGovFrameChecker is a tool for checking eGovFrame compatibility.",
	Long: "eGovFrameChecker is a cli tool for checking eGovFrame compatibility. " +
		"It checks whether the architecture criteria are met.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
		os.Exit(1)
	}
}
