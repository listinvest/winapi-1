package platform

import (
	"syscall"
)

var (
	libraries = make(map[string]library)
)

func findLibrary(libName string) (lib *library, err error) {
	if libItem, ok := libraries[libName]; ok {
		return &libItem, nil
	}

	dll, err := syscall.LoadDLL(libName)
	if err != nil {
		return nil, err
	}

	lib = newLibrary(dll)
	libraries[libName] = *lib
	return lib, nil
}

func Call(libName, procName string, args ...uintptr) (uintptr, uintptr, error) {
	lib, err := findLibrary(libName)
	if err != nil {
		return 0, 0, err
	}

	return lib.call(procName, args...)
}
