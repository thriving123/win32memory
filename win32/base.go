package win32

import "syscall"

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procCloseHandle = modkernel32.NewProc("CloseHandle")
	procOpenProcess = modkernel32.NewProc("OpenProcess")
)

const (
	INVALID_HANDLE_VALUE    = ^syscall.Handle(0)
	INVALID_FILE_SIZE       = 0xFFFFFFFF
	INVALID_FILE_ATTRIBUTES = 0xFFFFFFFF
)

func CloseHandle(object syscall.Handle) {
	syscall.SyscallN(procCloseHandle.Addr(), uintptr(object), 0, 0)
}
func boolToUintptr(value bool) uintptr {
	var valueRaw int32
	if value {
		valueRaw = 1
	} else {
		valueRaw = 0
	}

	return uintptr(valueRaw)
}
