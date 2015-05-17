package pty

import (
	"os"
	"syscall"
	"unsafe"
)

type Winsize struct {
	Row    uint16
	Col    uint16
	XPixel uint16
	YPixel uint16
}

func SetWinSize(f *os.File, winp *Winsize) error {
	return ioctl(uintptr(f.Fd()), syscall.TIOCSWINSZ, uintptr(unsafe.Pointer(winp)))
}

func SetWinSize2(f *os.File, row, col uint16) error {
	winp := &Winsize{Row: row, Col: col}

	return SetWinSize(f, winp)
}

func GetWinSize(f *os.File) (*Winsize, error) {
	winp := &Winsize{}

	e := ioctl(uintptr(f.Fd()), syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(winp)))
	if e != nil {
		return nil, e
	}
	return winp, nil
}

func GetWinSize2(f *os.File) (row, col uint16, e error) {
	winp, e := GetWinSize(f)
	if e != nil {
		return 0, 0, e
	}
	return winp.Row, winp.Col, nil
}
