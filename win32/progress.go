package win32

import (
	"errors"
	"syscall"
	"unsafe"
)

const (
	DELETE                   = 0x00010000
	READ_CONTROL             = 0x00020000
	WRITE_DAC                = 0x00040000
	WRITE_OWNER              = 0x00080000
	SYNCHRONIZE              = 0x00100000
	STANDARD_RIGHTS_REQUIRED = 0x000F0000
	STANDARD_RIGHTS_READ     = READ_CONTROL
	STANDARD_RIGHTS_WRITE    = READ_CONTROL
	STANDARD_RIGHTS_EXECUTE  = READ_CONTROL
	STANDARD_RIGHTS_ALL      = 0x001F0000
	SPECIFIC_RIGHTS_ALL      = 0x0000FFFF
	ACCESS_SYSTEM_SECURITY   = 0x01000000
	MAXIMUM_ALLOWED          = 0x02000000
)
const (
	PROCESS_TERMINATE                 = 0x0001
	PROCESS_CREATE_THREAD             = 0x0002
	PROCESS_SET_SESSIONID             = 0x0004
	PROCESS_VM_OPERATION              = 0x0008
	PROCESS_VM_READ                   = 0x0010
	PROCESS_VM_WRITE                  = 0x0020
	PROCESS_DUP_HANDLE                = 0x0040
	PROCESS_CREATE_PROCESS            = 0x0080
	PROCESS_SET_QUOTA                 = 0x0100
	PROCESS_SET_INFORMATION           = 0x0200
	PROCESS_QUERY_INFORMATION         = 0x0400
	PROCESS_SUSPEND_RESUME            = 0x0800
	PROCESS_QUERY_LIMITED_INFORMATION = 0x1000
	PROCESS_ALL_ACCESS                = STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0xFFFF
)

// GetModuleAddr 根据模块名获取模块基址
func GetModuleAddr(pid int, moduleName string) (int, error) {
	handle, err := CreateToolhelp32Snapshot(TH32CS_SNAPMODULE, uint32(pid))
	if err != nil {
		return -1, err
	}
	defer CloseHandle(handle)
	var info MODULEENTRY32
	info.Size = uint32(unsafe.Sizeof(info))
	err = Module32First(handle, &info)
	if err != nil {
		return -1, err
	}
	for {
		name := syscall.UTF16ToString(info.ModuleName[:])
		if moduleName == name {
			moduleAddressPtr := (*uintptr)(unsafe.Pointer(&info.ModBaseAddr))
			return int(*moduleAddressPtr), nil
		}
		err = Module32Next(handle, &info)
		if err != nil {
			break
		}
	}
	return -1, errors.New("module not found")
}

// GetPidByName 根据进程名获取进程Pid
func GetPidByName(processName string) (int, error) {
	handle, err := CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return -1, err
	}
	defer CloseHandle(handle)
	var info PROCESSENTRY32
	info.Size = uint32(unsafe.Sizeof(info))
	if err = Process32First(handle, &info); err != nil {
		return -1, err
	}
	for {
		name := syscall.UTF16ToString(info.ExeFile[:])
		if processName == name {
			return int(info.ProcessID), nil
		}
		if err = Process32Next(handle, &info); err != nil {
			break
		}
	}
	return -1, errors.New("process not found")
}

// OpenProcess 打开进程
func OpenProcess(desiredAccess uint32, inheritHandle bool, pid int) (syscall.Handle, error) {
	r1, _, e1 := syscall.SyscallN(
		procOpenProcess.Addr(),
		uintptr(desiredAccess),
		boolToUintptr(inheritHandle),
		uintptr(uint32(pid)))
	if r1 == 0 {
		if !errors.Is(e1, ERROR_SUCCESS) {
			return 0, e1
		} else {
			return 0, syscall.EINVAL
		}
	}
	return syscall.Handle(r1), nil
}
