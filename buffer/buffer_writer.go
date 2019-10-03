package buffer

import (
	"unicode/utf8"
)

func (b *Buffer) write(p []byte, l int) (int, error) {
	// make space and copy
	b.Grow(l)
	copy(b.buf[b.base+b.length:l], p)

	b.length += l

	return l, nil
}

func (b *Buffer) Write(p []byte) (int, error) {
	if l := len(p); l > 0 {
		return b.write(p, l)
	} else {
		return 0, nil
	}
}

func (b *Buffer) WriteString(s string) (int, error) {
	if len(s) > 0 {
		p := []byte(s)
		return b.write(p, len(p))
	}
	return 0, nil
}

func (b *Buffer) WriteRune(r rune) (int, error) {
	b.Grow(utf8.UTFMax)

	l := utf8.EncodeRune(b.buf[b.base+b.length:], r)
	b.length += l

	return l, nil
}
