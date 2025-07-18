package ver

import "fmt"

// CheckerVersion
// v0.1.0 : 신규 버전
// v0.2.0 :
//   - 상위 클래스 정보 기록 추가
//   - 불필요 error 반환 제거
//   - Repository 처리 상 오류 수정
// v0.2.1 : ANTLR4 처리 오류 수정
// v0.3.0 :
//   - eGovFrame 기본 버전 변경 (v4.2 -> v4.3)
//   - 확인이 필요한 중간 추상 클래스 출력
// v0.3.1 : Service/Repository에서 interface 제외 처리
// v0.4.0 :
//   - Repository 파일 패턴 추가 (*Repository.java)
//   - Hibernate 점검 시 상위 클래스 필드 확인 추가
const CheckerVersion = "v0.4.0"

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
