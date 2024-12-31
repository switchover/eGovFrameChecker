package rules

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/manifoldco/promptui"
	"github.com/switchover/eGovFrameChecker/internal/result"
	"github.com/switchover/eGovFrameChecker/internal/utils"
)

func InputVersionRule() {
	result.CheckResult.VersionRule = true
	// Default compatibility level is 2
	result.CheckResult.CompatibilityLevel = 2

	ok, v := eGovFrameVersion()
	if !ok {
		return
	}
	ver := utils.GetVersion(v.Major(), v.Minor())
	fmt.Printf("eGovFrame Version: %v\n", ver)
	result.CheckResult.EGovFrameVersion = ver

	ok, v = jdkVersion()
	if !ok {
		return
	}
	ver = utils.GetVersion(v.Major())
	fmt.Printf("Jdk Version: %v\n", ver)
	result.CheckResult.JdkVersion = ver

	ok, v = springBootVersion()
	if !ok {
		return
	}
	if v != nil {
		ver = utils.GetVersion(v.Major(), v.Minor(), v.Patch())
		fmt.Printf("Spring Boot Version: %v\n", ver)
		result.CheckResult.BootVersion = ver
	}

	ok, v = springVersion()
	if !ok {
		return
	}
	ver = utils.GetVersion(v.Major(), v.Minor(), v.Patch())
	fmt.Printf("Spring Version: %v\n", ver)
	result.CheckResult.SpringVersion = ver
}

func eGovFrameVersion() (ok bool, semVer *semver.Version) {
	prompt := promptui.Select{
		Label:     "eGovFrame version used",
		CursorPos: 1,
		Items: []string{
			"v4.3", "v4.2", "v4.1", "v4.0", "v3.10", "v3.9", "v3.8", "v3.7", "v3.6", "v3.5", "v3.1", "v3.0"},
	}
	_, version, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	v, err := semver.NewVersion(version)
	if err != nil {
		fmt.Printf("eGovFrame version error %v\n", err)
		return false, nil
	}
	return true, v
}

func jdkVersion() (ok bool, semVer *semver.Version) {
	prompt := promptui.Select{
		Label:     "Minimum supported or applied JDK version",
		CursorPos: 3,
		Items: []string{
			"5 (1.5)", "6 (1.6)", "7 (1.7)", "8 (LTS)", "11 (LTS)", "17 (LTS)", "21 (LTS)"},
	}
	_, versionString, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	version := versionString[:strings.Index(versionString, " ")]
	v, err := semver.NewVersion(version)
	if err != nil {
		fmt.Printf("Java version error %v\n", err)
		return false, nil
	}
	return true, v
}

func springBootVersion() (ok bool, semVer *semver.Version) {
	checkPrompt := promptui.Select{
		Label:     "Is Spring Boot applied? [Yes/No]",
		CursorPos: 1,
		Items:     []string{"Yes", "No"},
	}
	_, yesOrNo, err := checkPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	if yesOrNo == "No" {
		return true, nil
	}

	prompt := promptui.Prompt{
		Label: "Spring Boot version used",
		Validate: func(s string) error {
			_, err := semver.NewVersion(s)
			if err != nil {
				return errors.New("invalid version")
			}
			return nil
		},
	}
	version, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	v, err := semver.NewVersion(version)
	if err != nil {
		fmt.Printf("Spring Boot version error %v\n", err)
		return false, nil
	}
	return true, v
}

func springVersion() (ok bool, semVer *semver.Version) {
	prompt := promptui.Prompt{
		Label: "Spring Framework version used",
		Validate: func(s string) error {
			_, err := semver.NewVersion(s)
			if err != nil {
				return errors.New("invalid version")
			}
			return nil
		},
	}
	version, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	v, err := semver.NewVersion(version)
	if err != nil {
		fmt.Printf("Spring Framework version error %v\n", err)
		return false, nil
	}
	return true, v
}
