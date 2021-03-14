package os

import (
	"os"
)

type StringValidator func(string) bool

func GetEnv(name, fallback string) string {
	if v, ok := os.LookupEnv(name); !ok {
		return fallback
	} else {
		return v
	}
}

func GetEnv2(name, fallback string, valid StringValidator) string {
	if v, ok := os.LookupEnv(name); !ok {
		return fallback
	} else if valid != nil && !valid(v) {
		return fallback
	} else {
		return v
	}
}
