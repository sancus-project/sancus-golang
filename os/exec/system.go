package exec

import (
	"bytes"
	"os/exec"
)

func Exec(arg0 string, args ...string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(arg0, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
