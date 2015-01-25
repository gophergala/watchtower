package users

import (
	"github.com/gophergala/watchtower/messages"
)

// NewTCPIPUser creates a new user connected over a TCP/IP socket
// TODO: clean up by closing the connection when the user is deleteds
func NewTCPIPUser() User {
	return &tcpIPUser{
		messageQueue: make(chan messages.Message, 10),
	}
}

type tcpIPUser struct {
	id           uint32
	messageQueue chan messages.Message
}

func (t *tcpIPUser) ID() uint32 {
	return t.id
}

func (t *tcpIPUser) Send(m messages.Message, channelID uint32) error {
	t.messageQueue <- m
	return nil
}

func (t *tcpIPUser) setID(id uint32) {
	t.id = id
}
