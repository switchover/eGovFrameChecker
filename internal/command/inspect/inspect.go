package inspect

import (
	"log"
	"strings"

	"github.com/spf13/viper"
	c "github.com/switchover/eGovFrameChecker/internal/constant"
	"github.com/switchover/eGovFrameChecker/internal/examine/common"
	"github.com/switchover/eGovFrameChecker/internal/examine/controller"
	"github.com/switchover/eGovFrameChecker/internal/examine/repository"
	"github.com/switchover/eGovFrameChecker/internal/examine/service"
	"github.com/switchover/eGovFrameChecker/internal/json"
	"github.com/switchover/eGovFrameChecker/internal/target"
)

func Inspect() {
	targetDir := viper.GetString("inspect.target")
	packages := strings.Split(viper.GetString("inspect.packages"), ",")

	jsonStreamer, err := getJsonStreamer()
	if err != nil {
		log.Fatalln(err)
	}
	if jsonStreamer != nil {
		log.Println("JSON output file:", viper.GetString("inspect.json"))
		defer jsonStreamer.Close()
	}

	countOfFiles, err := target.GatherSourceFiles(targetDir, packages)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Total java files:", countOfFiles)

	if jsonStreamer != nil {
		err = writeSummary(jsonStreamer, targetDir, strings.Join(packages, ","), countOfFiles)
		if err != nil {
			jsonStreamer.Delete()
			log.Fatalln(err)
		}
	}

	err = controller.Examine(target.GetControllerFiles(), jsonStreamer)
	if err != nil {
		if jsonStreamer != nil {
			jsonStreamer.Delete()
		}
		log.Fatalln(err)
	}

	err = service.Examine(target.GetServiceFiles(), jsonStreamer)
	if err != nil {
		if jsonStreamer != nil {
			jsonStreamer.Delete()
		}
		log.Fatalln(err)
	}

	err = repository.Examine(target.GetRepositoryFiles(), jsonStreamer)
	if err != nil {
		if jsonStreamer != nil {
			jsonStreamer.Delete()
		}
		log.Fatalln(err)
	}

	toBeCheckedSuperClasses := common.GetToBeCheckedSuperClasses()
	if len(toBeCheckedSuperClasses) > 1 {
		log.Printf("%s%s Super classes to be checked: %s%s%s\n", c.IconCaution, c.Reset,
			c.MagentaUnderline, strings.Join(toBeCheckedSuperClasses, c.Reset+", "+c.MagentaUnderline), c.Reset)
	} else if len(toBeCheckedSuperClasses) == 1 {
		log.Printf("%s%s Super class to be checked: %s%s%s\n", c.IconCaution, c.Reset, c.MagentaUnderline, toBeCheckedSuperClasses[0], c.Reset)
	}
}

func getJsonStreamer() (streamer *json.Streamer, err error) {
	jsonFile := viper.GetString("inspect.json")

	if jsonFile != "" {
		streamer, err = json.NewJsonStreamer(jsonFile)
	}

	return
}

func writeSummary(streamer *json.Streamer, targetDir, packages string, totalFiles int) error {
	return streamer.WriteSummary(json.Summary{
		Target:   targetDir,
		Packages: packages,
		Files: json.FileCounts{
			Total:        totalFiles,
			Controllers:  len(target.GetControllerFiles()),
			Services:     len(target.GetServiceFiles()),
			Repositories: len(target.GetRepositoryFiles()),
		},
	})
}
