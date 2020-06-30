package platform

import (
	"syscall"
)

type PlatformApi struct {
	libraries map[string]library
}

func (api *PlatformApi) library(libName string) (lib *library, err error) {
	if libItem, ok := api.libraries[libName]; ok {
		return &libItem, nil
	}

	handle, err := syscall.LoadLibrary(libName)
	if err != nil {
		return nil, err
	}

	lib = &library{dll: handle}
	api.libraries[libName] = *lib
	return lib, nil
}

func (api *PlatformApi) Call(libName, procName string, errors map[syscall.Errno]error, args ...interface{}) (uintptr, uintptr, error) {
	lib, err := api.library(libName)
	if err != nil {
		return 0, 0, err
	}

	return lib.call(procName, errors, args)
}
