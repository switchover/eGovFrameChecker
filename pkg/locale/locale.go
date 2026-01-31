package locale

import "strings"

// extractLanguage extracts the language code from a locale string.
// Examples:
//   - "ko-KR" -> "ko"
//   - "en-US" -> "en"
//   - "ko_KR.UTF-8" -> "ko"
func extractLanguage(locale string) string {
	// Handle special cases
	if strings.HasPrefix(locale, "C.") || strings.HasPrefix(locale, "POSIX.") {
		return "en"
	}

	// Remove encoding suffix (e.g., ".UTF-8")
	if idx := strings.Index(locale, "."); idx != -1 {
		locale = locale[:idx]
	}

	// Extract language part (before underscore or hyphen)
	if idx := strings.IndexAny(locale, "_-"); idx != -1 {
		return locale[:idx]
	}

	return locale
}
