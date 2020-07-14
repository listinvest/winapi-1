package tools

import (
	"syscall"
	"unsafe"
)

func StringPointer(text string) *uint16 {
	ptr, _ := syscall.UTF16PtrFromString(text)
	return ptr
}

func StringUIntPointer(text string) uintptr {
	return uintptr(unsafe.Pointer(StringPointer(text)))
}

func BoolPointer(value bool) uintptr {
	var number int32 = 0

	if value {
		number = 1
	}

	return uintptr(number)
}
