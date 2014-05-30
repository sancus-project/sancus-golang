package log

import (
	"bytes"
	"fmt"
	"os"
)

var StderrBackend = FileBackend{f: os.Stderr}

// FileBackend
type FileBackend struct {
	f *os.File
}

func (l *FileBackend) LogWrite(level LogLevel, tag string, fmt string, a ...interface{}) (int, error) {
	return l.f.Write(l.Encode(level, tag, fmt, a...))
}

func (l *FileBackend) Encode(level LogLevel, tag string, format string, a ...interface{}) []byte {
	var b bytes.Buffer
	var c rune

	switch level {
	case DEBUG:
		c = 'D'
	case TRACE:
		c = 'T'
	case VERBOSE:
		c = 'V'
	case INFO:
		c = 'I'
	case WARN:
		c = 'W'
	case ERROR:
		c = 'E'
	case WTF:
		c = 'F'
	case ASSERT:
		c = 'A'
	default:
		c = '\000'
	}

	if tag != "" {
		if c != '\000' {
			b.WriteRune(c)
			b.WriteRune('/')
		}
		b.WriteString(tag)
		b.WriteString(": ")
	} else if c != '\000' {
		b.WriteRune(c)
		b.WriteString(": ")
	}

	if len(a) > 0 {
		b.WriteString(fmt.Sprintf(format, a...))
	} else {
		b.WriteString(format)
	}

	b.WriteRune('\n')
	return b.Bytes()
}
