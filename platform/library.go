package platform

import (
	"fmt"
	"syscall"
	"unsafe"
)

type library struct {
	dll        syscall.Handle
	procedures map[string]uintptr
}

func (l *library) loadProcedure(procName string) (uintptr, error) {
	proc, err := syscall.GetProcAddress(l.dll, procName)
	if err != nil {
		return -1, err
	}

	l.procedures[procName] = proc
	return l.procedures[procName], nil
}

func (l *library) procedure(name string) (proc uintptr, err error) {
	proc, ok := l.procedures[name]
	if !ok {
		proc, err = l.loadProcedure(name)
		if err != nil {
			return 0, err
		}
	}
	return
}

func pointer(index int, args []interface{}) uintptr {
	if len(args) <= index {
		return uintptr(unsafe.Pointer(nil))
	}

	if ptr, ok := args[0].(uintptr); ok {
		return ptr
	}

	return uintptr(unsafe.Pointer(&args[index]))
}

func (l *library) call(procName string, args ...interface{}) (uintptr, uintptr, error, syscall.Errno) {
	proc, err := l.procedure(procName)
	if err != nil {
		return 0, 0, err, 0
	}

	var result1, result2 uintptr
	var errno syscall.Errno

	argsCount := len(args)
	if argsCount < 4 {
		result1, result2, errno = syscall.Syscall(proc, uintptr(argsCount),
			pointer(0, args), pointer(1, args), pointer(2, args))
	} else if argsCount < 7 {
		result1, result2, errno = syscall.Syscall6(proc, uintptr(argsCount),
			pointer(0, args), pointer(1, args), pointer(2, args),
			pointer(3, args), pointer(4, args), pointer(5, args))
	} else if argsCount < 10 {
		result1, result2, errno = syscall.Syscall9(proc, uintptr(argsCount),
			pointer(0, args), pointer(1, args), pointer(2, args),
			pointer(3, args), pointer(4, args), pointer(5, args),
			pointer(6, args), pointer(7, args), pointer(8, args))
	} else if argsCount < 12 {
		result1, result2, errno = syscall.Syscall12(proc, uintptr(argsCount),
			pointer(0, args), pointer(1, args), pointer(2, args),
			pointer(3, args), pointer(4, args), pointer(5, args),
			pointer(6, args), pointer(7, args), pointer(8, args),
			pointer(9, args), pointer(10, args), pointer(11, args))
	} else if argsCount < 15 {
		result1, result2, errno = syscall.Syscall15(proc, uintptr(argsCount),
			pointer(0, args), pointer(1, args), pointer(2, args),
			pointer(3, args), pointer(4, args), pointer(5, args),
			pointer(6, args), pointer(7, args), pointer(8, args),
			pointer(9, args), pointer(10, args), pointer(11, args),
			pointer(12, args), pointer(13, args), pointer(14, args))
	} else if argsCount < 18 {
		result1, result2, errno = syscall.Syscall18(proc, uintptr(argsCount),
			pointer(0, args), pointer(1, args), pointer(2, args),
			pointer(3, args), pointer(4, args), pointer(5, args),
			pointer(6, args), pointer(7, args), pointer(8, args),
			pointer(9, args), pointer(10, args), pointer(11, args),
			pointer(12, args), pointer(13, args), pointer(14, args),
			pointer(15, args), pointer(16, args), pointer(17, args))
	} else {
		return 0, 0, fmt.Errorf("procedures with %d arguments does not supported", argsCount), errno
	}

	return result1, result2, err, errno
}
