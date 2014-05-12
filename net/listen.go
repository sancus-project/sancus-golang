package net

import (
	"net"
	"os"
	"strings"
)

type Listener net.Listener

var networks = map[string]string{
	"unix:": "unix",
	"file:": "unix",
}

func Listen(addr string) (Listener, error) {
	var network string

	if len(addr) > 0 {
		if addr[1] == '/' || addr[1:2] == "./" || addr[1:3] == "../" {
			network = "unix"
		} else {
			// check for known network family prefixes
			for p, n := range networks {
				if strings.HasPrefix(addr, p) {
					network = n
					addr = addr[len(p):]
					break
				}
			}
		}
	}

	if network == "unix" {
		_ = os.Remove(addr)
	} else if network == "" {
		network = "tcp"
	}

	return net.Listen(network, addr)
}
