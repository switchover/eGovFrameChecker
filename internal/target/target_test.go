package target

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/spf13/viper"
)

func TestConvertFilePathToPackageStyle(t *testing.T) {
	expectedWindows := "com\\example\\MyClass"
	if runtime.GOOS == "windows" {
		expectedWindows = "com.example.MyClass"
	}

	tests := []struct {
		name          string
		inspectTarget string
		filePath      string
		want          string
	}{
		{
			name:          "Unix style path",
			inspectTarget: "",
			filePath:      filepath.Join("com", "example", "MyClass.java"),
			want:          "com.example.MyClass",
		},
		{
			name:          "Windows style path",
			inspectTarget: "",
			filePath:      "com\\example\\MyClass.java",
			want:          expectedWindows,
		},
		{
			name:          "With inspect.target prefix",
			inspectTarget: filepath.Join("root", "dir"),
			filePath:      filepath.Join("root", "dir", "com", "example", "MyClass.java"),
			want:          "com.example.MyClass",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() { viper.Reset() })
			viper.Set("inspect.target", tt.inspectTarget)

			got := convertFilePathToPackageStyle(tt.filePath)
			if got != tt.want {
				t.Errorf("convertFilePathToPackageStyle(%q) = %q; want %q", tt.filePath, got, tt.want)
			}
		})
	}
}
