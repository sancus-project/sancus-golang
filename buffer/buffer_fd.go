package buffer

import (
	"syscall"
)

func (b *Buffer) ReadFrom(fd uintptr) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.grow(MinimumReadSpace)

	for {
		if rc, err := syscall.Read(int(fd), b.buf[b.base+b.length:]); err == nil {
			b.length += rc
			return rc, nil
		} else if err != syscall.EINTR {
			return -1, err
		}
	}
}

func (b *Buffer) WriteTo(fd uintptr) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.length > 0 {
		for {
			if wc, err := syscall.Write(int(fd), b.buf[b.base:b.base+b.length]); err == nil {
				b.skip(wc)
				return wc, nil
			} else if err != syscall.EINTR {
				return -1, err
			}
		}
	}
	return 0, nil
}
