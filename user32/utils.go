package user32

import (
	"github.com/Magic-Library/winapi/platform"
	"github.com/Magic-Library/winapi/tools"
	"syscall"
	"unsafe"
)

func RegisterClassEx(ex *WindowClassEx) uint16 {
	r1, _, _ := platform.Call(tools.User32, "RegisterClassExW", uintptr(unsafe.Pointer(ex)))
	return uint16(r1)
}

func MakeIntResource(res uint16) *uint16 {
	return (*uint16)(unsafe.Pointer(uintptr(res)))
}

func LoadCursor(instance uintptr, cursor uint16) uintptr {
	r1, _, _ := platform.Call(tools.User32, "LoadCursorW", instance, uintptr(unsafe.Pointer(MakeIntResource(cursor))))
	return r1
}

func LoadIcon(instance uintptr, icon uint16) uintptr {
	r1, _, _ := platform.Call(tools.User32, "LoadIconW", instance, uintptr(unsafe.Pointer(MakeIntResource(icon))))
	return r1
}

func GetClassNameToBuffer(window uintptr, buffer []uint16, max int) int {
	r1, _, _ := platform.Call(tools.User32, "GetClassNameW",
		window,
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(max))

	return int(r1)
}

func GetClassName(window uintptr) string {
	const max = 255
	buf := make([]uint16, max)
	GetClassNameToBuffer(window, buf, max)
	return syscall.UTF16ToString(buf)
}
