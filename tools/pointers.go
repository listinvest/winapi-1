package tools

import (
	"syscall"
	"unsafe"
)

func UIntPointer(arg interface{}) uintptr {
	if arg == nil {
		return 0
	}

	return uintptr(unsafe.Pointer(&arg))
}

func StringPointer(text string) *uint16 {
	ptr, _ := syscall.UTF16PtrFromString(text)
	return ptr
}
