package pty

import (
	"os"
	"syscall"
	"unsafe"
)

func GetTermios(f *os.File) (*syscall.Termios, error) {
	termp := &syscall.Termios{}

	e := ioctl(uintptr(f.Fd()), syscall.TCGETS, uintptr(unsafe.Pointer(termp)))
	if e != nil {
		return nil, e
	}
	return termp, nil
}
