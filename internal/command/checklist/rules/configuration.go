package rules

import (
	"log"

	"github.com/manifoldco/promptui"
	"github.com/switchover/eGovFrameChecker/internal/result"
)

func InputConfigurationRule() {
	result.CheckResult.ConfigurationRule = true

	checkPrompt := promptui.Select{
		Label: "Has the transaction been configured? (AOP or Annotation based) [Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, yesOrNo, err := checkPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	if yesOrNo == "No" {
		result.CheckResult.ConfigurationRule = false
	}

	checkPrompt = promptui.Select{
		Label: "Has the DB Connection Pool(eg: DBCP2, HikariCP) been configured? [Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, yesOrNo, err = checkPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	if yesOrNo == "No" {
		result.CheckResult.ConfigurationRule = false
	}

	checkPrompt = promptui.Select{
		Label: "Does the setup basically follow the eGovFrame configuration guide? [Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, yesOrNo, err = checkPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	if yesOrNo == "No" {
		result.CheckResult.ConfigurationRule = false
	}
}
