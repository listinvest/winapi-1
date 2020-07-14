package user32

import (
	"github.com/Magic-Library/winapi/platform"
	"github.com/Magic-Library/winapi/tools"
	"syscall"
	"unsafe"
)

const (
	Monitor_DefaultNull    = 0x00000000
	Monitor_DefaultPrimary = 0x00000001
	Monitor_DefaultNearest = 0x00000002
)

type MonitorInfo struct {
	Size    uint32
	Monitor Rect
	Work    Rect
	DwFlags uint32
}

type ExtendedMonitorInfo struct {
	MonitorInfo
	nameArray [tools.DeviceNameLength]uint16
}

func (info *ExtendedMonitorInfo) Name() string {
	utf16array := make([]uint16, 32)
	for i := range info.nameArray {
		utf16array[i] = info.nameArray[i]
	}
	return syscall.UTF16ToString(utf16array)
}

func MonitorFromPoint(x, y int, dwFlags uint32) uintptr {
	point := Point{X: int32(x), Y: int32(y)}
	r, _, _ := platform.Call(tools.User32, "MonitorFromPoint", uintptr(unsafe.Pointer(&point)), uintptr(dwFlags))
	return r
}

func MonitorFromRect(rc *Rect, dwFlags uint32) uintptr {
	r, _, _ := platform.Call(tools.User32, "MonitorFromRect", uintptr(unsafe.Pointer(rc)), uintptr(dwFlags))
	return r
}

func MonitorFromWindow(window uintptr, dwFlags uint32) uintptr {
	r, _, _ := platform.Call(tools.User32, "MonitorFromWindow", window, uintptr(dwFlags))
	return r
}

func GetMonitorInfo(monitor uintptr) *MonitorInfo {
	info := MonitorInfo{}
	info.Size = uint32(unsafe.Sizeof(info))

	if getMonitorInfoW(monitor, uintptr(unsafe.Pointer(&info))) {
		return &info
	}
	return nil
}

func GetExtendedMonitorInfo(monitor uintptr) *ExtendedMonitorInfo {
	info := ExtendedMonitorInfo{}
	info.Size = uint32(unsafe.Sizeof(info))

	if getMonitorInfoW(monitor, uintptr(unsafe.Pointer(&info))) {
		return &info
	}
	return nil
}

func getMonitorInfoW(monitor, result uintptr) bool {
	r, _, _ := platform.Call(tools.User32, "GetMonitorInfoW", monitor, result)
	return r != 0
}

type EnumDisplayFunction func(monitor, hdc uintptr, rect *Rect, param uintptr) bool

func EnumDisplayMonitors(dc uintptr, clip *Rect, fnEnum EnumDisplayFunction, dwData uintptr) bool {
	r, _, _ := platform.Call(tools.User32, "EnumDisplayMonitors",
		dc,
		uintptr(unsafe.Pointer(clip)),
		syscall.NewCallback(func(monitor, hdc uintptr, rect *Rect, param uintptr) uintptr {
			return tools.BoolPointer(fnEnum(monitor, hdc, rect, param))
		}),
		dwData,
	)
	return r != 0
}

func EnumDisplaySettings(deviceName string, iModeNum uint32, dwFlags uint32) *DeviceSettings {
	settings := DeviceSettings{}
	r, _, _ := platform.Call(tools.User32, "EnumDisplaySettingsExW",
		tools.StringUIntPointer(deviceName),
		uintptr(iModeNum),
		uintptr(unsafe.Pointer(&settings)),
		uintptr(dwFlags),
	)

	if r == 0 {
		return nil
	} else {
		return &settings
	}
}

func ChangeDisplaySettings(deviceName string, settings *DeviceSettings, hwnd uintptr, dwFlags uint32, lParam uintptr) int {
	r, _, _ := platform.Call(tools.User32, "ChangeDisplaySettingsExW",
		tools.StringUIntPointer(deviceName),
		uintptr(unsafe.Pointer(settings)),
		hwnd,
		uintptr(dwFlags),
		lParam,
	)
	return int(r)
}
