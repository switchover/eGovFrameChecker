//go:build linux

package locale

import (
	"os"

	"github.com/spf13/viper"
)

// GetLanguage returns the language code from the system locale.
// For example, if the locale is "ko_KR.UTF-8", it returns "ko".
func GetLanguage() string {
	// First, check "locale" parameter value
	locale := viper.GetString("inspect.locale")
	if locale != "" {
		return extractLanguage(locale)
	}

	// Check locale environment variables in order of priority
	for _, envVar := range []string{"LC_ALL", "LC_MESSAGES", "LANG"} {
		if locale := os.Getenv(envVar); locale != "" {
			return extractLanguage(locale)
		}
	}
	return "ko" // default to Korean
}
