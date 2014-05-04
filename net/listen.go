package net

import (
	stdnet "net"
	"os"
	"strings"
)

type Listener stdnet.Listener

func Listen(addr string) (Listener, error) {
	network := "tcp"

	// let the standard Listen return any error, not us
	if len(addr) > 0 {
		if addr[1] == '/' || addr[1:2] == "./" || addr[1:3] == "../" {
			network = "unix"
		} else if strings.HasPrefix(addr, "unix:") || strings.HasPrefix(addr, "file:") {
			network = "unix"
			addr = addr[5:]
		}
	}

	if network == "unix" {
		_ = os.Remove(addr)
	}

	return stdnet.Listen(network, addr)
}
