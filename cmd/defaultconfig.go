package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/switchover/eGovFrameChecker/internal/command/defaultconfig"
)

func init() {
	var overwrite bool

	defaultConfigCmd.Flags().BoolVarP(&overwrite, "overwrite", "o", false, "Overwrite config.ini when it exists")

	_ = viper.BindPFlag("defaultconfig.overwrite", defaultConfigCmd.Flags().Lookup("overwrite"))

	rootCmd.AddCommand(defaultConfigCmd)
}

var defaultConfigCmd = &cobra.Command{
	Use:   "defaultconfig",
	Short: "Write default 'config.ini' file",
	Long:  `Write default 'config.ini' file to current directory for custom configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		defaultconfig.Write()
	},
}
