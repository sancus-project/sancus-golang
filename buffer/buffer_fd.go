package buffer

import (
	"syscall"
)

func (b *Buffer) ReadFrom(fd uintptr) (int, error) {
	tail := b.Grow(0) - b.base - b.length
	rc, err := syscall.Read(int(fd), b.buf[b.base+b.length:tail])

	if err == nil && rc > 0 {
		b.length += rc
	}

	return rc, err
}

func (b *Buffer) WriteTo(fd uintptr) (int, error) {
	if b.length > 0 {
		wc, err := syscall.Write(int(fd), b.buf[b.base:b.length])

		if err == nil && wc > 0 {
			b.Skip(wc)
		}

		return wc, err
	}
	return 0, nil
}
