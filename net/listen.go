package net

import (
	stdnet "net"
	"os"
	"strings"
)

type Listener stdnet.Listener

func Listen(addr string) (Listener, error) {
	dom := "tcp"

	// let the standard Listen return any error, not us
	if len(addr) > 0 {
		if addr[1] == '/' || addr[1:2] == "./" || addr[1:3] == "../" {
			dom = "unix"
		} else if strings.HasPrefix(addr, "unix:") || strings.HasPrefix(addr, "file:") {
			dom = "unix"
			addr = addr[5:]
		}
	}

	if dom == "unix" {
		_ = os.Remove(addr)
	}

	return stdnet.Listen(dom, addr)
}
