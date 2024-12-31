package utils

import (
	"strconv"
	"strings"
)

func GetVersion(versions ...uint64) string {
	var buildVersion strings.Builder
	//buildVersion.WriteString("v")
	for index, version := range versions {
		if index != 0 {
			buildVersion.WriteString(".")
		}
		buildVersion.WriteString(strconv.FormatUint(version, 10))
	}
	return buildVersion.String()
}
