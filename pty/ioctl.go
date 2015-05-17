package pty

import "syscall"

func ioctl(fd, request, argp uintptr) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, request, argp)
	if err != 0 {
		return err
	}
	return nil
}
