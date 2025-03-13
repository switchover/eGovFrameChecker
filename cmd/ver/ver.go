package ver

import "fmt"

// CheckerVersion
// v0.1.0 : 신규 버전
// v0.2.0 :
//   - 상위 클래스 정보 기록 추가
//   - 불필요 error 반환 제거
//   - Repository 처리 상 오류 수정
const CheckerVersion = "v0.2.0"

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
