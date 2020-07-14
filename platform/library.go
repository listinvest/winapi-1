package platform

import (
	"syscall"
)

type library struct {
	dll        *syscall.DLL
	procedures map[string]*syscall.Proc
}

func newLibrary(dll *syscall.DLL) *library {
	return &library{
		dll:        dll,
		procedures: make(map[string]*syscall.Proc),
	}
}

func (l *library) loadProcedure(procName string) (*syscall.Proc, error) {
	proc, err := l.dll.FindProc(procName)
	if err != nil {
		return nil, err
	}

	l.procedures[procName] = proc
	return l.procedures[procName], nil
}

func (l *library) procedure(name string) (proc *syscall.Proc, err error) {
	proc, ok := l.procedures[name]
	if !ok {
		proc, err = l.loadProcedure(name)
		if err != nil {
			return nil, err
		}
	}
	return
}

func (l *library) call(procName string, args ...uintptr) (uintptr, uintptr, error) {
	proc, err := l.procedure(procName)
	if err != nil {
		return 0, 0, err
	}

	return proc.Call(args...)
}
