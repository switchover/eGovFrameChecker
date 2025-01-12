package inspect

import (
	"log"
	"strings"

	"github.com/spf13/viper"
	"github.com/switchover/eGovFrameChecker/internal/examine/controller"
	"github.com/switchover/eGovFrameChecker/internal/examine/repository"
	"github.com/switchover/eGovFrameChecker/internal/examine/service"
	"github.com/switchover/eGovFrameChecker/internal/target"
)

func Inspect() {
	targetDir := viper.GetString("inspect.target")
	packages := strings.Split(viper.GetString("inspect.packages"), ",")

	countOfFiles, err := target.GatherSourceFiles(targetDir, packages)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Total java files:", countOfFiles)

	err = controller.Examine(target.GetControllerFiles())
	if err != nil {
		log.Fatalln(err)
	}

	err = service.Examine(target.GetServiceFiles())
	if err != nil {
		log.Fatalln(err)
	}

	err = repository.Examine(target.GetRepositoryFiles())
	if err != nil {
		log.Fatalln(err)
	}
}
