package log

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

// deeper increments the stack position, unless it's zero used to mark it's disabled
func deeper(depth int) int {
	if depth < 1 {
		return 0
	}
	return depth + 1
}

// FormatBytes formats a multiline []byte
func (self *Logger) FormatBytes(calldepth int, prefix string, data []byte) []string {
	var lines []string

	for _, s := range bytes.Split(data, []byte{'\n'}) {
		s = bytes.TrimRight(s, "\n\r\t ")
		lines = append(lines, string(s))
	}

	return self.FormatLines(deeper(calldepth), prefix, lines)
}

// Format formats a multiline string
func (self *Logger) Format(calldepth int, prefix string, data string) []string {
	var lines []string

	for _, s := range strings.Split(data, "\n") {
		s = strings.TrimRight(s, "\n\r\t ")
		lines = append(lines, s)
	}

	return self.FormatLines(deeper(calldepth), prefix, lines)
}

// FormatLines formats an array of string lines
func (self *Logger) FormatLines(calldepth int, prefix string, lines []string) []string {
	self.mu.Lock()
	defer self.mu.Unlock()

	// remove trailing empty lines
	i := len(lines)
	for i > 0 {
		if len(lines[i-1]) > 0 {
			break
		} else {
			i--
		}
	}
	lines = lines[:i]

	// compose prefix
	prefix = formatPrefix(deeper(calldepth), self.prefix, prefix)
	if len(prefix) > 0 {
		if len(lines) > 0 {
			for i, s := range lines {
				if len(s) == 0 {
					lines[i] = prefix
				} else {
					lines[i] = strings.Join([]string{prefix, s}, " ")
				}
			}
		} else {
			// print a marker if there is no message
			lines = []string{prefix}
		}
	}

	return lines
}

func formatPrefix(calldepth int, p0, p1 string) string {
	var b strings.Builder

	b.WriteString(p0) // logger prefix
	b.WriteString(p1) // context prefix

	// file and function
	if calldepth > 0 {
		if pc, fileName, fileLine, ok := runtime.Caller(calldepth + 1); ok {
			if fn := runtime.FuncForPC(pc); fn != nil {
				pos := b.Len()

				b.WriteString(fn.Name())
				if len(fileName) > 0 {
					if fileLine > 0 {
						b.WriteString(fmt.Sprintf(" (%s:%v)", fileName, fileLine))
					} else {
						b.WriteString(fmt.Sprintf(" (%s)", fileName))
					}
				} else if fileLine > 0 {
					b.WriteString(fmt.Sprintf(":%v", fileLine))
				}

				if b.Len() > pos {
					b.WriteRune(':')
				}
			}
		}
	}

	return strings.TrimRight(b.String(), "\t ")
}
