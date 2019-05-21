package log

import (
	"log"
)

// Turn *Logger into the output backend of the standard Go logger
func (logger *Logger) SetStandard() *Logger {
	log.SetOutput(logger)
	log.SetFlags(0)

	return logger
}
