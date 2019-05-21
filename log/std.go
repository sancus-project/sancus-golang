package log

import (
	"log"
)

// Turn *Logger into the output backend of the standard Go logger
func (self *Logger) SetStandard() *Logger {
	log.SetOutput(self)
	log.SetFlags(0)

	return self
}
