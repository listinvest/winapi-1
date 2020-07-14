package user32

import (
	"github.com/Magic-Library/winapi/platform"
	"github.com/Magic-Library/winapi/tools"
)

func GetDC(window uintptr) uintptr {
	r, _, _ := platform.Call(tools.User32, "GetDC", window)
	return r
}

func ReleaseDC(window, dc uintptr) bool {
	r, _, _ := platform.Call(tools.User32, "ReleaseDC", window, dc)
	return r != 0
}
