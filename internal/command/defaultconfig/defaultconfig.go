package defaultconfig

import (
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/switchover/eGovFrameChecker/internal/config"
)

func Write() {
	if fileExists("config.ini") {
		if !viper.GetBool("defaultconfig.overwrite") {
			log.Println("'config.ini' file exists. Use --overwrite flag to overwrite.")
			return
		}
	}

	err := config.Write("config.ini")
	if err != nil {
		log.Fatalf("Failed to write default config file: %v\n", err)
	}
	log.Println("Default config file 'config.ini' has been written.")
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
