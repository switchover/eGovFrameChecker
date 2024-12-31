package checklist

import (
	"fmt"

	"github.com/switchover/eGovFrameChecker/internal/command/checklist/rules"
	c "github.com/switchover/eGovFrameChecker/internal/constant"
	"github.com/switchover/eGovFrameChecker/internal/criteria"
	"github.com/switchover/eGovFrameChecker/internal/result"
)

func Check() {
	rules.InputVersionRule()

	violations := criteria.Check()
	if len(violations) == 0 {
		fmt.Printf("%s Version Rule: %sCompatible%s\n", c.IconOkay, c.LightBlue, c.Reset)
		if result.CheckResult.CompatibilityLevel == 2 {
			fmt.Printf("  %s Compatibility Level: %s%d%s\n",
				c.IconWarn, c.LightBlue, result.CheckResult.CompatibilityLevel, c.Reset)
		} else {
			fmt.Printf("  %s Compatibility Level: %s%d%s\n",
				c.IconWarn, c.Green, result.CheckResult.CompatibilityLevel, c.Reset)
		}
	} else {
		fmt.Printf("%s Version Rule: %sNot compatible%s\n", c.IconNotOkay, c.LightBlue, c.Reset)
		for _, violation := range violations {
			fmt.Printf("  %s %s%s%s\n", c.IconWarn, c.Grey, violation, c.Reset)
		}
	}

	fmt.Println("--------------------------------------------------------------------------------")

	rules.InputConfigurationRule()
	if result.CheckResult.ConfigurationRule {
		fmt.Printf("%s Configuration Rule: %sCompatible%s\n", c.IconOkay, c.LightBlue, c.Reset)
	} else {
		fmt.Printf("%s Configuration Rule: %sNot compatible%s\n", c.IconNotOkay, c.LightBlue, c.Reset)
	}
	fmt.Println("--------------------------------------------------------------------------------")

	rules.InputArchitectureRule()

	printResult()
}

func printResult() {
	fmt.Println()
	fmt.Println("------------------------------------------------------------")
	fmt.Println(c.ResultBanner)
	fmt.Println("------------------------------------------------------------")
	fmt.Print("eGovFrame Compatibility: ")
	if result.CheckResult.IsCompatible() {
		fmt.Printf("%sCompatible%s ", c.LightBlue, c.Reset)
		if result.CheckResult.CompatibilityLevel == 2 {
			fmt.Printf("(level: %s%d%s)\n", c.LightBlue, result.CheckResult.CompatibilityLevel, c.Reset)
		} else {
			fmt.Printf("(level: %s%d%s)\n", c.Green, result.CheckResult.CompatibilityLevel, c.Reset)
		}
	} else {
		fmt.Printf("%sNot compatible%s\n", c.Magenta, c.Reset)
	}

	fmt.Print("  - Version criteria: ")
	if result.CheckResult.VersionRule {
		fmt.Printf("%sSatisfied%s\n", c.LightBlue, c.Reset)
	} else {
		fmt.Printf("%sNot satisfied%s\n", c.Magenta, c.Reset)
	}

	fmt.Print("  - Configuration criteria: ")
	if result.CheckResult.ConfigurationRule {
		fmt.Printf("%sSatisfied%s\n", c.LightBlue, c.Reset)
	} else {
		fmt.Printf("%sNot satisfied%s\n", c.Magenta, c.Reset)
	}

	fmt.Print("  - MVC criteria: ")
	if result.CheckResult.PresentationLayerRule {
		fmt.Printf("%sSatisfied%s\n", c.LightBlue, c.Reset)
	} else {
		fmt.Printf("%sNot satisfied%s\n", c.Magenta, c.Reset)
	}

	fmt.Print("  - Service criteria: ")
	if result.CheckResult.ServiceLayerRule {
		fmt.Printf("%sSatisfied%s\n", c.LightBlue, c.Reset)
	} else {
		fmt.Printf("%sNot satisfied%s\n", c.Magenta, c.Reset)
	}

	fmt.Print("  - Data access criteria: ")
	if result.CheckResult.DataAccessLayerRule {
		fmt.Printf("%sSatisfied%s\n", c.LightBlue, c.Reset)
	} else {
		fmt.Printf("%sNot satisfied%s\n", c.Magenta, c.Reset)
	}
	fmt.Println("------------------------------------------------------------")
}
