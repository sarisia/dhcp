//go:build go1.12 && !darwin && !freebsd && !linux && !netbsd && !openbsd

package nclient4

import (
	"errors"
	"net"
)

var ErrNotImplemented = errors.New("not implemented for this architecture")

func NewRawUDPConn(iface string, port int) (net.PacketConn, error) {
	return nil, ErrNotImplemented
}
