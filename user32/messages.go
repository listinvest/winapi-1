package user32

import (
	"github.com/Magic-Library/winapi/platform"
	"github.com/Magic-Library/winapi/tools"
	"unsafe"
)

func PostQuitMessage(code int) {
	_, _, _ = platform.Call(tools.User32, "PostQuitMessage", uintptr(code))
}

func GetMessage(message *Message, window uintptr, filterMin, filterMax uint32) int {
	r, _, _ := platform.Call(tools.User32, "GetMessageW",
		uintptr(unsafe.Pointer(message)), window, uintptr(filterMin), uintptr(filterMax))

	return int(r)
}

func TranslateMessage(message *Message) bool {
	r, _, _ := platform.Call(tools.User32, "TranslateMessage", uintptr(unsafe.Pointer(message)))
	return r != 0
}

func DispatchMessage(message *Message) uintptr {
	r, _, _ := platform.Call(tools.User32, "DispatchMessageW", uintptr(unsafe.Pointer(message)))
	return r
}

func SendMessage(window uintptr, message uint32, wParam, lParam uintptr) uintptr {
	r, _, _ := platform.Call(tools.User32, "SendMessageW", window, uintptr(message), wParam, lParam)
	return r
}

func SendMessageTimeout(window uintptr, message uint32, wParam, lParam uintptr, fuFlags, uTimeout uint32) uintptr {
	r, _, _ := platform.Call(tools.User32, "SendMessageTimeout",
		window,
		uintptr(message),
		wParam,
		lParam,
		uintptr(fuFlags),
		uintptr(uTimeout))
	return r
}

func PostMessage(window uintptr, message uint32, wParam, lParam uintptr) bool {
	r, _, _ := platform.Call(tools.User32, "PostMessageW",
		window,
		uintptr(message),
		wParam,
		lParam)
	return r != 0
}

func WaitMessage() bool {
	r, _, _ := platform.Call(tools.User32, "WaitMessage")
	return r != 0
}
