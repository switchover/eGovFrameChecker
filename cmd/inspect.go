package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/switchover/eGovFrameChecker/internal/command/inspect"
)

func init() {
	var target string
	var packages string

	inspectCmd.Flags().StringVarP(&target, "target", "t", "", "Target directory to inspect")
	inspectCmd.Flags().StringVarP(&packages, "packages", "p", "", "Packages to inspect with comma separated")

	_ = inspectCmd.MarkFlagRequired("target")
	_ = inspectCmd.MarkFlagRequired("packages")

	_ = viper.BindPFlag("inspect.target", inspectCmd.Flags().Lookup("target"))
	_ = viper.BindPFlag("inspect.packages", inspectCmd.Flags().Lookup("packages"))

	rootCmd.AddCommand(inspectCmd)
}

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect eGovFrame compatibility architecture criteria",
	Long:  `Inspect eGovFrame compatibility architecture criteria and save related data.`,
	Run: func(cmd *cobra.Command, args []string) {
		inspect.Inspect()
	},
}
