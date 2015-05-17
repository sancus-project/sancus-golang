package pty

import (
	"os"
	"os/exec"
	"syscall"
)

func StartCommand(name string, arg ...string) (*os.File, error) {
	c := exec.Command(name, arg...)

	return StartCmd(c)
}

func StartCmd(c *exec.Cmd) (*os.File, error) {
	pty, pts, err := Open(nil, nil)
	if err != nil {
		return nil, err
	}

	// on success c will kept it held
	defer pts.Close()

	c.Stdin = pts
	c.Stdout = pts
	c.Stderr = pts

	if c.SysProcAttr == nil {
		c.SysProcAttr = &syscall.SysProcAttr{}
	}
	c.SysProcAttr.Setctty = true // pts is controlling terminal
	c.SysProcAttr.Setsid = true  // new session

	err = c.Start()
	if err != nil {
		pty.Close()
		return nil, err
	}

	return pty, nil
}
