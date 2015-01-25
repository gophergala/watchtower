package users

import (
	"net"

	"github.com/gophergala/watchtower/messages"
)

// NewTCPIPUser creates a new user connected over a TCP/IP socket
// TODO: clean up by closing the connection when the user is deleteds
func NewTCPIPUser(conn net.Conn) User {
	return &tcpIPUser{
		conn: conn,
	}
}

type tcpIPUser struct {
	id   uint32
	conn net.Conn
}

func (t *tcpIPUser) ID() uint32 {
	return t.id
}

func (t *tcpIPUser) Send(m messages.Message) error {
	_, err := t.conn.Write([]byte("Message received."))
	return err
}

func (t *tcpIPUser) setID(id uint32) {
	t.id = id
}
