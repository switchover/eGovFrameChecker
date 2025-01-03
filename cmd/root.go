package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "egovchecker",
		Short: "eGovFrameChecker is a tool for checking eGovFrame compatibility.",
		Long: "eGovFrameChecker is a cli tool for checking eGovFrame compatibility. " +
			"It checks whether the architecture criteria are met.",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config.ini. Or ./config.ini is used)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home and working directory with name "config" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(getWorkingDirectory())
		viper.AddConfigPath(getDirWithExecutable())

		viper.SetConfigName("config")
		viper.SetConfigType("ini")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}

	setDefaultValues()
}

func setDefaultValues() {
	// eg: viper.SetDefault("projects.packages", "com.example")
}

func getWorkingDirectory() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory, %s", err)
	}
	return path
}

func getDirWithExecutable() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatalf("Error getting executable path, %s", err)
	}
	return filepath.Dir(ex)
}
