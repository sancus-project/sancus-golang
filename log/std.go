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

func (logger *Logger) NewStd(prefix string, v Variant) *log.Logger {
	return log.New(logger.NewWriter(prefix, v), "", 0)
}
