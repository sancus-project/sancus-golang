package buffer

import (
	"syscall"
)

func (b *Buffer) peek(p []byte) (int, error) {
	if b.length == 0 {
		return 0, syscall.EAGAIN
	} else if l := len(p); l == 0 {
		return 0, syscall.ENOBUFS
	} else {
		if l > b.length {
			l = b.length
		}
		return l, nil
	}
}

func (b *Buffer) Read(p []byte) (int, error) {
	if l, err := b.peek(p); err != nil {
		return 0, err
	} else {
		copy(p, b.buf[b.base:l])
		b.Skip(l)
		return l, nil
	}
}

func (b *Buffer) Peek(buf []byte) (int, error) {
	return b.peek(buf)
}
