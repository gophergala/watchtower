package users

import (
	"testing"
)

var (
	_ User = (*tcpIPUser)(nil)
)

func TestTcpIPUser(t *testing.T) {
	// TODO: Test Send() by opening an actual TCP/IP
	// connection, passing it in and reading what
	// Watchtower sends
	u := NewTCPIPUser(nil)

	if u.ID() != 0 {
		t.Error("invalid id")
	}

	u.setID(1)

	if u.ID() != 1 {
		t.Error("invalid id after set")
	}
}
