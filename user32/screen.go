package user32

import (
	"github.com/Magic-Library/winapi/platform"
	"github.com/Magic-Library/winapi/tools"
	"unsafe"
)

func ScreenToClient(window uintptr, x, y int) (int, int, bool) {
	return ScreenPointToClient(window, Point{X: int32(x), Y: int32(y)})
}

func ScreenPointToClient(window uintptr, point Point) (int, int, bool) {
	r, _, _ := platform.Call(tools.User32, "ScreenToClient", window, uintptr(unsafe.Pointer(&point)))
	return int(point.X), int(point.Y), r != 0
}
