//go:build darwin

package locale

import (
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

// GetLanguage returns the language code from the system locale.
// For example, if the locale is "ko_KR", it returns "ko".
func GetLanguage() string {
	// First, check "locale" parameter value
	locale := viper.GetString("inspect.locale")
	if locale != "" {
		return extractLanguage(locale)
	}

	// Fall back to macOS defaults command
	if lang := getAppleLocale(); lang != "" {
		return lang
	}

	return "ko" // default to Korean
}

// getAppleLocale gets the locale from macOS system preferences.
func getAppleLocale() string {
	cmd := exec.Command("defaults", "read", "-g", "AppleLocale")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	locale := strings.TrimSpace(string(output))
	return extractLanguage(locale)
}
