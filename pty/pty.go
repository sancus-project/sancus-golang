package pty

import (
	"C"
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func Open() (master, slave *os.File, err error) {
	ptm, err := os.OpenFile("/dev/ptmx", os.O_RDWR, 0)
	if err != nil {
		return nil, nil, err
	}

	ptsName, err := ptsname(ptm)
	if err != nil {
		ptm.Close()
		return nil, nil, err
	}

	err = unlockpt(ptm)
	if err != nil {
		ptm.Close()
		return nil, nil, err
	}

	pts, err := os.OpenFile(ptsName, os.O_RDWR|syscall.O_NOCTTY, 0)
	if err != nil {
		ptm.Close()
		return nil, nil, err
	}

	return ptm, pts, nil
}

func ptsname(ptm *os.File) (string, error) {
	var n C.uint
	err := ioctl(ptm.Fd(), syscall.TIOCGPTN, uintptr(unsafe.Pointer(&n)))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("/dev/pts/%d", n), nil
}

func unlockpt(ptm *os.File) error {
	var n C.uint
	return ioctl(ptm.Fd(), syscall.TIOCSPTLCK, uintptr(unsafe.Pointer(&n)))
}
