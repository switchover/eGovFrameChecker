//go:build windows

package terminal

import (
	"os"
	"syscall"
)

func EnableANSI() {
	handle := syscall.Handle(os.Stdout.Fd())
	var mode uint32
	syscall.GetConsoleMode(handle, &mode)
	mode |= 0x0004 // ENABLE_VIRTUAL_TERMINAL_PROCESSING
	syscall.SetConsoleMode(handle, mode)
}
