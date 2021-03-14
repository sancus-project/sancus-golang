package buffer

import (
	"io"
	"syscall"
	"unicode/utf8"
)

func (b *Buffer) write(p []byte, l int) (int, error) {
	// make space and copy
	b.grow(l)
	copy(b.buf[b.base+b.length:l], p)

	b.length += l

	return l, nil
}

func (b *Buffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if l := len(p); l > 0 {
		return b.write(p, l)
	} else {
		return 0, nil
	}
}

func (b *Buffer) WriteString(s string) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(s) > 0 {
		p := []byte(s)
		return b.write(p, len(p))
	}
	return 0, nil
}

func (b *Buffer) WriteRune(r rune) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.grow(utf8.UTFMax)

	l := utf8.EncodeRune(b.buf[b.base+b.length:], r)
	b.length += l

	return l, nil
}

func (b *Buffer) readOnceFrom(r io.Reader) (int64, error) {
	b.grow(MinimumReadSpace)

	for {
		if rc, err := r.Read(b.buf[b.base+b.length:]); err == nil {
			if rc == 0 {
				return 0, io.EOF
			} else {
				b.length += rc
				return int64(rc), nil
			}
		} else if err != syscall.EINTR {
			return -1, err
		}
	}
}

func (b *Buffer) ReadOnceFrom(r io.Reader) (int64, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.readOnceFrom(r)
}

func (b *Buffer) ReadFrom(r io.Reader) (int64, error) {
	var n int64

	b.mu.Lock()
	defer b.mu.Unlock()

	for {
		if rc, err := b.readOnceFrom(r); err == nil {
			n += rc
		} else if err != syscall.EAGAIN {
			return n, err
		}
	}
}
