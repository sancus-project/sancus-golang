package log

import (
	"fmt"
)

type ErrorString struct {
	str string
}

func (e *ErrorString) Error() string {
	return e.str
}

func NewError(str string, a ...interface{}) error {
	if len(a) > 0 {
		str = fmt.Sprintf(str, a...)
	}

	return &ErrorString{str}
}

func (l *Logger) NewError(str string, a ...interface{}) error {
	if l.tag != "" {
		str = "%s: " + str
		a = append([]interface{}{l.tag}, a...)
	}

	if len(a) > 0 {
		str = fmt.Sprintf(str, a...)
	}

	return &ErrorString{str}
}
