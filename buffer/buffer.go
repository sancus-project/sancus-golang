package buffer

import (
	"sync"
)

type Buffer struct {
	mu     sync.Mutex
	buf    []byte
	base   int
	length int
}

const (
	InitialBufferSize = 32
)

func New(size uint) *Buffer {
	if size == 0 {
		size = InitialBufferSize
	}

	return &Buffer{
		buf: make([]byte, size),
	}
}

func (b *Buffer) Len() int {
	return b.length
}

func (b *Buffer) Size() int {
	return len(b.buf)
}

func (b *Buffer) Cap() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	return len(b.buf) - b.base
}

func (b *Buffer) reset() {
	b.length = 0
	b.base = 0
}

func (b *Buffer) Reset() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.reset()
}

func (b *Buffer) bytes() []byte {
	return b.buf[b.base:b.length]
}

func (b *Buffer) Bytes() []byte {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.bytes()
}

func (b *Buffer) grow(n int) int {
	var n0, n1, n2 int

	// find new size
	if n0 = len(b.buf); n0 == 0 {
		n2 = InitialBufferSize
	} else {
		n2 = n0
	}

	n1 = (b.base + b.length + n)
	for n1 > n2 {
		n2 *= 2
	}

	if n2 != n0 {
		// resize and rebase
		b1 := make([]byte, n2)
		if n0 > 0 {
			copy(b1, b.bytes())
		}
		b.buf = b1
		b.base = 0
	}

	return n2
}

func (b *Buffer) Grow(n int) int {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.grow(n)
}

func (b *Buffer) skip(n int) int {
	if n >= b.length {
		b.reset()
	} else {
		b.base += n
		b.length -= n

		if b.base >= b.length {
			// rebase
			copy(b.buf[0:b.length], b.bytes())
			b.base = 0
		}
	}

	return b.length
}

func (b *Buffer) Skip(n int) int {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.skip(n)
}
