package pty

import (
	"C"
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func Open(termp *syscall.Termios, winp *Winsize) (pty, pts *os.File, err error) {
	var ptsName string

	// open pty master
	pty, err = os.OpenFile("/dev/ptmx", os.O_RDWR, 0)
	if err != nil {
		goto fail_open_pty
	}

	// open pty slave
	ptsName, err = ptsname(pty)
	if err != nil {
		goto fail
	}

	err = unlockpt(pty)
	if err != nil {
		goto fail
	}

	pts, err = os.OpenFile(ptsName, os.O_RDWR|syscall.O_NOCTTY, 0)
	if err != nil {
		goto fail
	}

	// configure pts
	if termp != nil {
		err = tcsetattr(pts, termp)
	}

	if err == nil && winp != nil {
		err = SetWinSize(pts, winp)
	}

	if err == nil {
		return pty, pts, nil
	}

	pts.Close()
fail:
	pty.Close()
fail_open_pty:
	return nil, nil, err
}

func ptsname(pty *os.File) (string, error) {
	var n C.uint
	err := ioctl(pty.Fd(), syscall.TIOCGPTN, uintptr(unsafe.Pointer(&n)))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("/dev/pts/%d", n), nil
}

func unlockpt(pty *os.File) error {
	var n C.uint
	return ioctl(pty.Fd(), syscall.TIOCSPTLCK, uintptr(unsafe.Pointer(&n)))
}

func tcsetattr(pts *os.File, termp *syscall.Termios) error {
	return ioctl(pts.Fd(), syscall.TCSETS, uintptr(unsafe.Pointer(termp)))
}
