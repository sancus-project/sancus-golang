package exec

import (
	"github.com/kballard/go-shellquote"
)

// shell argument escaping
type ShellArgs []string

func (args ShellArgs) String() string {
	return shellquote.Join([]string(args)...)
}

func NewShellArgs(args ...string) ShellArgs {
	return args[:]
}
