package common

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"github.com/switchover/eGovFrameChecker/pkg/java"
)

func CheckClass(section string, listener *java.Listener) bool {
	annotations := viper.GetString(fmt.Sprintf("%s.%s", section, "classAnnotations"))
	for _, annotation := range strings.Split(annotations, ",") {
		for _, classAnnotation := range listener.ClassAnnotations {
			if classAnnotation == strings.TrimSpace(annotation) {
				return true
			}
		}
	}

	return false
}

func CheckMethods(section string, listener *java.Listener) bool {
	annotations := viper.GetString(fmt.Sprintf("%s.%s", section, "methodAnnotations"))
	for _, annotation := range strings.Split(annotations, ",") {
		if listener.MethodAnnotations[strings.TrimSpace(annotation)] == true {
			return true
		}
	}

	return false
}
