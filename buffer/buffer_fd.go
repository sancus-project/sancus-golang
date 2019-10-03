package buffer

import (
	"syscall"
)

func (b *Buffer) ReadFrom(fd uintptr) (int, error) {
	tail := b.Grow(0) - b.base - b.length

	for {
		if rc, err := syscall.Read(int(fd), b.buf[b.base+b.length:tail]); err == nil {
			b.length += rc
			return rc, nil
		} else if err != syscall.EINTR {
			return -1, err
		}
	}
}

func (b *Buffer) WriteTo(fd uintptr) (int, error) {
	if b.length > 0 {
		for {
			if wc, err := syscall.Write(int(fd), b.buf[b.base:b.length]); err == nil {
				b.Skip(wc)
				return wc, nil
			} else if err != syscall.EINTR {
				return -1, err
			}
		}
	}
	return 0, nil
}
