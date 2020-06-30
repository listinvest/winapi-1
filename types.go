package winapi

import "github.com/Magic-Library/winapi/platform"

type Window struct {
	winapi  *platform.PlatformApi
	pointer uintptr
}

func (w *Window) Animate(time int64) {
	w.winapi.Call(User32, "AnimateWindow", nil, w.pointer)
}

func (w *Window) Close() {

}

func (w *Window) Destroy() {

}
