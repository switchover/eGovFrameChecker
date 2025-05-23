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

	toBeCheckedSuperClasses := common.GetToBeCheckedSuperClasses()
	if len(toBeCheckedSuperClasses) > 1 {
		log.Printf("%s%s Super classes to be checked: %v\n", c.IconCaution, c.Reset, toBeCheckedSuperClasses)
	} else if len(toBeCheckedSuperClasses) == 1 {
		log.Printf("%s%s Super class to be checked: %v\n", c.IconCaution, c.Reset, toBeCheckedSuperClasses[0])
	}
}
