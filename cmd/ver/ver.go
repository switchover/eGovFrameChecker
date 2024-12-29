package ver

import "fmt"

// CheckerVersion
// v0.1.0 : 신규 버전
const CheckerVersion = "v0.1.0"

// build flags
var (
	BuildTime  string
	GoVersion  string
	CommitHash string
)

func PrintBuildFlags() {
	if BuildTime != "" {
		fmt.Println(" - Build time  :", BuildTime)
	}
	if GoVersion != "" {
		fmt.Println(" - Go version  :", GoVersion)
	}
	if CommitHash != "" {
		fmt.Println(" - Commit hash :", CommitHash)
	}
}
