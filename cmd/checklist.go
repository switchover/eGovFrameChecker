package cmd

import (
	"github.com/spf13/cobra"
	"github.com/switchover/eGovFrameChecker/internal/command/checklist"
)

func init() {
	rootCmd.AddCommand(checklistCmd)
}

var checklistCmd = &cobra.Command{
	Use:   "checklist",
	Short: "Check the checklist for eGovFrame compatibility verification",
	Long:  `Check the checklist for eGovFrame compatibility verification using user input information such as eGov Framework version.`,
	Run: func(cmd *cobra.Command, args []string) {
		checklist.Check()
	},
}
