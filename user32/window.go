package user32

import (
	"github.com/Magic-Library/winapi/platform"
	"github.com/Magic-Library/winapi/tools"
	"unsafe"
)

func CreateWindow(extendedStyle uint, className, windowName *uint16,
	style uint, x, y, width, height int, parentWindow uintptr, menu uintptr,
	instance uintptr, param unsafe.Pointer) uintptr {

	r, _, _ := platform.Call(tools.User32, "CreateWindowExW",
		uintptr(extendedStyle),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)),
		uintptr(style),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		parentWindow,
		menu,
		instance,
		uintptr(param))
	return r
}

func WindowProcedure(hWnd uintptr, msg uint32, firstParam, secondParam uintptr) uintptr {
	switch msg {
	case 2:
		PostQuitMessage(0)
	default:
		return CallDefaultWindowProcedure(hWnd, msg, firstParam, secondParam)
	}
	return 0
}

func CallDefaultWindowProcedure(window uintptr, message uint32, firstParam, secondParam uintptr) uintptr {
	r, _, _ := platform.Call(tools.User32, "DefWindowProcW", window, uintptr(message), firstParam, secondParam)
	return r
}

func ShowWindow(window uintptr, command int) bool {
	r, _, _ := platform.Call(tools.User32, "ShowWindow", window, uintptr(command))
	return r != 0
}

func UpdateWindow(window uintptr) bool {
	r, _, _ := platform.Call(tools.User32, "UpdateWindow", window)
	return r != 0
}

func MoveWindow(window uintptr, x, y, width, height int, repaint bool) bool {
	r, _, _ := platform.Call(tools.User32, "MoveWindow",
		window,
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		tools.BoolPointer(repaint))
	return r != 0
}
