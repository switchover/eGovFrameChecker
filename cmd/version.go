package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/switchover/eGovFrameChecker/cmd/ver"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of eGovFrameChecker",
	Long:  `Print the version and information of eGovFrameChecker.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("eGovFrame Compatibility Checker Version :", ver.CheckerVersion)
		ver.PrintBuildFlags()
	},
}
