package exec

import (
	"go.sancus.dev/sancus/os"
)

var (
	shell = os.GetEnv2("SHELL", "/bin/sh", func(s string) bool { return len(s) > 0 })
)

// system(cmd) -> sh -c "cmd"
func System(s string) (error, string, string) {
	return Exec(shell, "-c", s)
}
