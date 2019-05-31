package log

import (
	"bytes"
	"fmt"
	"path/filepath"
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
func (logger *Logger) FormatBytes(calldepth int, prefix string, data []byte) []string {
	var lines []string

	for _, s := range bytes.Split(data, []byte{'\n'}) {
		s = bytes.TrimRight(s, "\n\r\t ")
		lines = append(lines, string(s))
	}

	return logger.FormatLines(deeper(calldepth), prefix, lines)
}

// Format formats a multiline string
func (logger *Logger) Format(calldepth int, prefix string, data string) []string {
	var lines []string

	for _, s := range strings.Split(data, "\n") {
		s = strings.TrimRight(s, "\n\r\t ")
		lines = append(lines, s)
	}

	return logger.FormatLines(deeper(calldepth), prefix, lines)
}

// FormatLines formats an array of string lines
func (logger *Logger) FormatLines(calldepth int, prefix string, lines []string) []string {
	logger.ctx.Lock()
	defer logger.ctx.Unlock()

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
	flags := logger.ctx.Flags()
	prefix = formatPrefix(deeper(calldepth), flags, logger.prefix, prefix)
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

func formatPrefix(calldepth int, flags uint, p0, p1 string) string {
	var b strings.Builder

	b.WriteString(p0) // logger prefix
	b.WriteString(p1) // context prefix

	// file and function
	if calldepth > 0 && flags&(Lshortfile|Llongfile|Lfileline|Lpackage|Lfunc) != 0 {
		if pc, fileName, fileLine, ok := runtime.Caller(calldepth + 1); ok {
			var fnName string

			// reduce fileName
			if flags&Lshortfile != 0 {
				fileName = filepath.Base(fileName)
			} else if flags&Llongfile == 0 {
				fileName = ""
			}

			// reduce fileLine
			if flags&Lfileline == 0 {
				fileLine = 0
			}

			// function
			if flags&(Lpackage|Lfunc) != 0 {
				if fn := runtime.FuncForPC(pc); fn != nil {
					fnName = fn.Name()

					// reduce fnName
					if flags&Lpackage == 0 {
						// no Lpackage, only include the function name
						fnName = filepath.Ext(fnName)[1:]
					} else if flags&Lfunc == 0 {
						// no Lfunc, remove the function name
						if ext := filepath.Ext(fnName); len(ext) > 0 {
							l := len(fnName) - len(ext)
							fnName = fnName[:l]
						}
					}
				}
			}

			// capture buffer position to detect changes
			pos := b.Len()

			if len(fnName) > 0 {
				b.WriteString(fnName)

				if len(fileName) > 0 {
					if fileLine > 0 {
						b.WriteString(fmt.Sprintf(" (%s:%v)", fileName, fileLine))
					} else {
						b.WriteString(fmt.Sprintf(" (%s)", fileName))
					}
				} else if fileLine > 0 {
					b.WriteString(fmt.Sprintf(":%v", fileLine))
				}
			} else if len(fileName) > 0 {
				if fileLine > 0 {
					b.WriteString(fmt.Sprintf("%s:%v", fileName, fileLine))
				} else {
					b.WriteString(fmt.Sprintf("%s", fileName))
				}
			} else if fileLine > 0 {
				b.WriteString(fmt.Sprintf(":%v", fileLine))
			}

			// and if we wrote anything, add a delimiter.
			if b.Len() > pos {
				b.WriteRune(':')
			}
		}
	}

	return strings.TrimRight(b.String(), "\t ")
}
