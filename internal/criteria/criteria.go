package criteria

import (
	_ "embed"
	"encoding/json"
	"log"

	"github.com/Masterminds/semver/v3"
	"github.com/switchover/eGovFrameChecker/internal/result"
)

//go:embed assets/criteria.json
var criteriaJSON []byte

type Version map[string]any

type Criteria struct {
	Versions map[string]Version `json:"versions"`
}

var criteria Criteria

func init() {
	if err := json.Unmarshal(criteriaJSON, &criteria); err != nil {
		log.Fatalf("Failed to load criteria configurations: %v\n", err)
	}
}

func Check() (violations []string) {
	baseVersion, _ := semver.NewVersion(result.CheckResult.EGovFrameVersion)

	eGovFrame := criteria.Versions[baseVersion.String()]
	if eGovFrame == nil {
		log.Fatalf("eGovFrame version %s is not supported\n", baseVersion)
	}

	violations = append(violations, jdkCheck(eGovFrame)...)
	if result.CheckResult.BootVersion != "" {
		violations = append(violations, bootCheck(eGovFrame)...)
		adjustCompatibilityLevel(eGovFrame)
	}
	violations = append(violations, springCheck(eGovFrame)...)

	return
}

func jdkCheck(eGovFrame Version) (violations []string) {
	jdk := eGovFrame["jdk"].(string)
	constraint, err := semver.NewConstraint(jdk)
	if err != nil {
		log.Fatalf("Failed to parse jdk constraint %v\n", err)
	}
	jdkVersion, _ := semver.NewVersion(result.CheckResult.JdkVersion)
	ok, _ := constraint.Validate(jdkVersion)
	if !ok {
		result.CheckResult.VersionRule = false
		violations = append(violations, eGovFrame["jdk-message"].(string))
	}
	return
}

func springCheck(eGovFrame Version) (violations []string) {
	spring := eGovFrame["spring"].(string)
	constraint, err := semver.NewConstraint(spring)
	if err != nil {
		log.Fatalf("Failed to parse spring constraint %v\n", err)
	}
	springVersion, _ := semver.NewVersion(result.CheckResult.SpringVersion)
	ok, _ := constraint.Validate(springVersion)
	if !ok {
		result.CheckResult.VersionRule = false
		violations = append(violations, eGovFrame["spring-message"].(string))
	}
	return
}

func bootCheck(eGovFrame Version) (violations []string) {
	boot := eGovFrame["boot"].(string)

	if boot == "" {
		violations = append(violations, eGovFrame["boot-message"].(string))
		return
	}
	constraint, err := semver.NewConstraint(boot)
	if err != nil {
		log.Fatalf("Failed to parse boot constraint %v\n", err)
	}
	bootVersion, _ := semver.NewVersion(result.CheckResult.BootVersion)
	ok, _ := constraint.Validate(bootVersion)
	if !ok {
		result.CheckResult.VersionRule = false
		violations = append(violations, eGovFrame["boot-message"].(string))
	}
	return
}

func adjustCompatibilityLevel(eGovFrame Version) {
	bootBase := eGovFrame["boot-base"].(string)
	bootBaseVersion, _ := semver.NewVersion(bootBase)

	bootVersion, _ := semver.NewVersion(result.CheckResult.BootVersion)

	if bootBaseVersion.Minor() != bootVersion.Minor() {
		result.CheckResult.CompatibilityLevel = 1
	}
}
