package pty

import (
	"os"
	"os/exec"
	"syscall"
)

func StartCommand(termp *syscall.Termios, winp *Winsize,
	name string, arg ...string) (*os.File, error) {

	c := exec.Command(name, arg...)

	return StartCmd(termp, winp, c)
}

func StartCmd(termp *syscall.Termios, winp *Winsize, c *exec.Cmd) (*os.File, error) {
	pty, pts, err := Open(termp, winp)
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
