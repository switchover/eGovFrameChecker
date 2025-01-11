package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/switchover/eGovFrameChecker/internal/command/inspect"
)

func init() {
	var target string
	var packages string
	var verbose bool
	var output bool
	var skipFileError bool

	inspectCmd.Flags().StringVarP(&target, "target", "t", "", "Target directory to inspect")
	inspectCmd.Flags().StringVarP(&packages, "packages", "p", "", "Packages to inspect with comma separated")
	inspectCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	inspectCmd.Flags().BoolVarP(&output, "output", "o", false, "Output result to CSV file")
	inspectCmd.Flags().BoolVarP(&skipFileError, "skip", "s", false, "Skip file error")

	_ = inspectCmd.MarkFlagRequired("target")
	_ = inspectCmd.MarkFlagRequired("packages")

	_ = viper.BindPFlag("inspect.target", inspectCmd.Flags().Lookup("target"))
	_ = viper.BindPFlag("inspect.packages", inspectCmd.Flags().Lookup("packages"))
	_ = viper.BindPFlag("inspect.verbose", inspectCmd.Flags().Lookup("verbose"))
	_ = viper.BindPFlag("inspect.output", inspectCmd.Flags().Lookup("output"))
	_ = viper.BindPFlag("inspect.skip", inspectCmd.Flags().Lookup("skip"))

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
