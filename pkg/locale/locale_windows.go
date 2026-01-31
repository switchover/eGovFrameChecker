//go:build windows

package locale

import (
	"syscall"
	"unsafe"

	"github.com/spf13/viper"
)

var (
	kernel32                     = syscall.NewLazyDLL("kernel32.dll")
	procGetUserDefaultLocaleName = kernel32.NewProc("GetUserDefaultLocaleName")
)

const localNameMaxLength = 85

// GetLanguage returns the language code from the system locale.
// For example, if the locale is "ko-KR", it returns "ko".
func GetLanguage() string {
	// First, check "locale" parameter value
	locale := viper.GetString("inspect.locale")
	if locale != "" {
		return extractLanguage(locale)
	}

	// Fall back to Windows API
	if lang := getWindowsLocale(); lang != "" {
		return extractLanguage(lang)
	}

	return "ko" // default to Korean
}

// getWindowsLocale gets the locale using Windows API.
func getWindowsLocale() string {
	buf := make([]uint16, localNameMaxLength)
	ret, _, _ := procGetUserDefaultLocaleName.Call(
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(localNameMaxLength),
	)
	if ret == 0 {
		return ""
	}
	return syscall.UTF16ToString(buf)
}
