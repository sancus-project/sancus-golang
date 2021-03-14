package log

import (
	"errors"
	"fmt"
)

func NewError(str string, a ...interface{}) error {
	if len(a) > 0 {
		str = fmt.Sprintf(str, a...)
	}

	return errors.New(str)
}

func (l *Logger) NewError(str string, a ...interface{}) error {
	if len(a) > 0 {
		if l.tag != "" {
			str = "%s: " + str
			a = append([]interface{}{l.tag}, a...)
		}

		str = fmt.Sprintf(str, a...)
	} else if l.tag != "" {
		str = l.tag + ": " + str
	}

	return errors.New(str)
}
