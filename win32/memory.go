package win32

import (
	"errors"
	"syscall"
	"unsafe"
)

var (
	procReadProcessMemory  = modkernel32.NewProc("ReadProcessMemory")
	procWriteProcessMemory = modkernel32.NewProc("WriteProcessMemory")
)

// writeProcessMemory 进程写内存
func writeProcessMemory(process syscall.Handle, baseAddress uintptr, buffer *byte, size uint32, numberOfBytesOut *uint32) error {
	r1, _, e1 := syscall.SyscallN(
		procWriteProcessMemory.Addr(),
		uintptr(process),
		baseAddress,
		uintptr(unsafe.Pointer(buffer)),
		uintptr(size),
		uintptr(unsafe.Pointer(numberOfBytesOut)),
		0)
	if r1 == 0 {
		if !errors.Is(e1, ERROR_SUCCESS) {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

// readProcessMemory 进程读内存
func readProcessMemory(process syscall.Handle, baseAddress uintptr, buffer *byte, size uint32, numberOfBytesRead *uint32) error {
	r1, _, e1 := syscall.SyscallN(
		procReadProcessMemory.Addr(),
		uintptr(process),
		baseAddress,
		uintptr(unsafe.Pointer(buffer)),
		uintptr(size),
		uintptr(unsafe.Pointer(numberOfBytesRead)),
		0)
	if r1 == 0 {
		if !errors.Is(e1, ERROR_SUCCESS) {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

// ReadInt 读内存整数型
func ReadInt(pid int, address int) (int, error) {
	process, err := OpenProcess(PROCESS_ALL_ACCESS, false, pid)
	if err != nil {
		return 0, err
	}
	defer CloseHandle(process)
	var tempData int
	if err := readProcessMemory(process, uintptr(address), (*byte)(unsafe.Pointer(&tempData)), 4, nil); err != nil {
		return 0, err
	}
	return tempData, nil
}

// WriteInt 写内存整数型
func WriteInt(pid int, address int, value int) error {
	process, err := OpenProcess(PROCESS_ALL_ACCESS, false, pid)
	if err != nil {
		return err
	}
	defer CloseHandle(process)
	return writeProcessMemory(process, uintptr(address), (*byte)(unsafe.Pointer(&value)), 4, nil)
}

// ReadInt64 读内存64位整数型
func ReadInt64(pid int, address int) (int64, error) {
	process, err := OpenProcess(PROCESS_ALL_ACCESS, false, pid)
	if err != nil {
		return 0, err
	}
	defer CloseHandle(process)
	var tempData int64
	if err := readProcessMemory(process, uintptr(address), (*byte)(unsafe.Pointer(&tempData)), 8, nil); err != nil {
		return 0, err
	}
	return tempData, nil
}

// WriteInt64 写内存64位整数型
func WriteInt64(pid int, address int, value int64) error {
	process, err := OpenProcess(PROCESS_ALL_ACCESS, false, pid)
	if err != nil {
		return err
	}
	defer CloseHandle(process)
	return writeProcessMemory(process, uintptr(address), (*byte)(unsafe.Pointer(&value)), 8, nil)
}

// ReadByte 读内存字节型
func ReadByte(pid int, address int) (byte, error) {
	process, err := OpenProcess(PROCESS_ALL_ACCESS, false, pid)
	if err != nil {
		return 0, err
	}
	defer CloseHandle(process)
	var tempData byte
	if err := readProcessMemory(process, uintptr(address), (*byte)(&tempData), 1, nil); err != nil {
		return 0, err
	}
	return tempData, nil
}

// WriteByte 写内存字节型
func WriteByte(pid int, address int, value byte) error {
	process, err := OpenProcess(PROCESS_ALL_ACCESS, false, pid)
	if err != nil {
		return err
	}
	defer CloseHandle(process)
	return readProcessMemory(process, uintptr(address), (*byte)(&value), 1, nil)
}

// ReadFloat 读内存浮点型
func ReadFloat(pid int, address int) (float32, error) {
	process, err := OpenProcess(PROCESS_ALL_ACCESS, false, pid)
	if err != nil {
		return 0, err
	}
	defer CloseHandle(process)
	var tempData float32
	if err := readProcessMemory(process, uintptr(address), (*byte)(unsafe.Pointer(&tempData)), 4, nil); err != nil {
		return 0, err
	}
	return tempData, nil
}

// WriteFloat 写内存浮点型
func WriteFloat(pid int, address int, value float32) error {
	process, err := OpenProcess(PROCESS_ALL_ACCESS, false, pid)
	if err != nil {
		return err
	}
	defer CloseHandle(process)
	return writeProcessMemory(process, uintptr(address), (*byte)(unsafe.Pointer(&value)), 4, nil)
}

// ReadFloat64 读内存64位浮点型
func ReadFloat64(pid int, address int) (float64, error) {
	process, err := OpenProcess(PROCESS_ALL_ACCESS, false, pid)
	if err != nil {
		return 0, err
	}
	defer CloseHandle(process)
	var tempData float64
	if err := readProcessMemory(process, uintptr(address), (*byte)(unsafe.Pointer(&tempData)), 8, nil); err != nil {
		return 0, err
	}
	return tempData, nil
}

// WriteFloat64 写内存64位浮点型
func WriteFloat64(pid int, address int, value float64) error {
	process, err := OpenProcess(PROCESS_ALL_ACCESS, false, pid)
	if err != nil {
		return err
	}
	defer CloseHandle(process)
	return writeProcessMemory(process, uintptr(address), (*byte)(unsafe.Pointer(&value)), 8, nil)
}

// ReadString 读内存字符串
func ReadString(pid int, address int, size int) (string, error) {
	process, err := OpenProcess(PROCESS_ALL_ACCESS, false, pid)
	if err != nil {
		return "", err
	}
	defer CloseHandle(process)
	tempData := make([]byte, size)
	if err := readProcessMemory(process, uintptr(address), (*byte)(unsafe.Pointer(&tempData)), uint32(size), nil); err != nil {
		return "", err
	}
	return string(tempData[:]), nil
}

// WriteString 写内存字符串
func WriteString(pid int, address int, value string) error {
	process, err := OpenProcess(PROCESS_ALL_ACCESS, false, pid)
	if err != nil {
		return err
	}
	defer CloseHandle(process)
	var tempData = []byte(value)
	return writeProcessMemory(process, uintptr(address), (*byte)(unsafe.Pointer(&tempData)), uint32(len(tempData)), nil)
}

// ReadBytes 读内存字节集
func ReadBytes(pid int, address int, size int) ([]byte, error) {
	process, err := OpenProcess(PROCESS_ALL_ACCESS, false, pid)
	if err != nil {
		return nil, err
	}
	defer CloseHandle(process)
	tempData := make([]byte, size)
	if err := readProcessMemory(process, uintptr(address), (*byte)(unsafe.Pointer(&tempData)), uint32(size), nil); err != nil {
		return nil, err
	}
	return tempData, nil
}

// WriteBytes 写内存字节集
func WriteBytes(pid int, address int, value []byte) error {
	process, err := OpenProcess(PROCESS_ALL_ACCESS, false, pid)
	if err != nil {
		return err
	}
	defer CloseHandle(process)
	return writeProcessMemory(process, uintptr(address), (*byte)(unsafe.Pointer(&value)), uint32(len(value)), nil)
}
