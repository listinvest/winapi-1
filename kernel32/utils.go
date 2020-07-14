package kernel32

import (
	"github.com/Magic-Library/winapi/platform"
	"github.com/Magic-Library/winapi/tools"
	"unsafe"
)

func GetModuleHandle(name string) uintptr {
	var namePointer uintptr

	if name == "" {
		namePointer = 0
	} else {
		namePointer = uintptr(unsafe.Pointer(tools.StringPointer(name)))
	}

	hInstance, _, _ := platform.Call(tools.Kernel32, "GetModuleHandleW", namePointer)
	return hInstance
}
