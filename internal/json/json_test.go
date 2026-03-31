package json_test

import (
	stdjson "encoding/json"
	"os"
	"path/filepath"
	"testing"

	j "github.com/switchover/eGovFrameChecker/internal/json"
)

func TestAddThreeFileTypes(t *testing.T) {
	// 준비: 임시 디렉터리와 출력 파일
	d := t.TempDir()
	out := filepath.Join(d, "out.json")

	// 스트리머 생성
	streamer, err := j.NewJsonStreamer(out)
	if err != nil {
		t.Fatalf("NewJsonStreamer failed: %v", err)
	}

	// Summary 작성
	summary := j.Summary{
		Target: "testpath",
		Files: j.FileCounts{
			Total:        5,
			Controllers:  1,
			Services:     2,
			Repositories: 2,
		},
	}
	if err := streamer.WriteSummary(summary); err != nil {
		_ = streamer.Close()
		t.Fatalf("WriteSummary failed: %v", err)
	}

	// 각각의 타입에 대해 AddViolation 호출 (Controller -> Service -> Repository)
	if err := streamer.AddViolation(j.Controller, j.Violation{FilePath: "Controller.java", PackageName: "pkg", ClassName: "Ctl", Description: "desc"}); err != nil {
		_ = streamer.Close()
		t.Fatalf("AddViolation Controller failed: %v", err)
	}
	if err := streamer.AddViolation(j.Service, j.Violation{FilePath: "Service.java", PackageName: "pkg", ClassName: "Svc1", Description: "desc"}); err != nil {
		_ = streamer.Close()
		t.Fatalf("AddViolation Service failed: %v", err)
	}
	if err := streamer.AddViolation(j.Service, j.Violation{FilePath: "Service.java", PackageName: "pkg", ClassName: "Svc2", Description: "desc"}); err != nil {
		_ = streamer.Close()
		t.Fatalf("AddViolation Service failed: %v", err)
	}
	if err := streamer.AddViolation(j.Repository, j.Violation{FilePath: "Repository.java", PackageName: "pkg", ClassName: "Repo1", Description: "desc"}); err != nil {
		_ = streamer.Close()
		t.Fatalf("AddViolation Repository failed: %v", err)
	}
	if err := streamer.AddViolation(j.Repository, j.Violation{FilePath: "Repository.java", PackageName: "pkg", ClassName: "Repo2", Description: "desc"}); err != nil {
		_ = streamer.Close()
		t.Fatalf("AddViolation Repository failed: %v", err)
	}

	// 파일 닫기
	if err := streamer.Close(); err != nil {
		t.Fatalf("Close failed: %v", err)
	}

	// 출력 파일 읽기 및 파싱
	b, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}

	var got struct {
		Summary    j.Summary     `json:"summary"`
		Controller []j.Violation `json:"controller"`
		Service    []j.Violation `json:"service"`
		Repository []j.Violation `json:"repository"`
	}
	if err := stdjson.Unmarshal(b, &got); err != nil {
		t.Fatalf("failed to unmarshal output JSON: %v\ncontent:\n%s", err, string(b))
	}

	// 검증: 각 배열이 1개씩 포함되어 있는지
	t.Run("counts", func(t *testing.T) {
		if len(got.Controller) != 1 {
			t.Fatalf("expected 1 controller, got %d", len(got.Controller))
		}
		if len(got.Service) != 2 {
			t.Fatalf("expected 2 service, got %d", len(got.Service))
		}
		if len(got.Repository) != 2 {
			t.Fatalf("expected 2 repository, got %d", len(got.Repository))
		}
	})

	// 항목 내용 검증
	t.Run("contents", func(t *testing.T) {
		if got.Controller[0].ClassName != "Ctl" {
			t.Fatalf("controller class name mismatch: %s", got.Controller[0].ClassName)
		}
		if got.Service[0].ClassName != "Svc1" {
			t.Fatalf("service class name mismatch: %s", got.Service[0].ClassName)
		}
		if got.Repository[0].ClassName != "Repo1" {
			t.Fatalf("repository class name mismatch: %s", got.Repository[0].ClassName)
		}
	})
}
