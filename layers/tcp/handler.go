package tcp

import (
	"net"
	"sync"
	"time"
)

type state int

const (
	stateNotConnected  state = 0
	stateJoinedChannel state = 1
	stateDisconnected  state = 5

	joinChannelMessageType = byte('J')
)

var (
	bufferPool = sync.Pool{
		New: func() interface{} { return make([]byte, 1024) },
	}
)

func Handle(c net.Conn) error {
	var s state
	for {
		switch s {
		case stateNotConnected:
			// Waiting for the sender to register / pick a channel
			buffer := bufferPool.Get().([]byte) // Borrow a buffer from the pool

			buf := make([]byte, 512)
			// Read the incoming connection into the buffer.
			_, err := c.Read(buf)
			if err != nil {
				return err
			}

			if buf[0] == joinChannelMessageType {
				// Parse the message
				// Reply with an ack
				c.Write([]byte(""))

				// Join channel
				s = stateJoinedChannel
			}

			bufferPool.Put(buffer) // Give the buffer back to the pool

		case stateJoinedChannel:
			// Waiting for sender to send or receive messages
			time.Sleep(time.Second)
		case stateDisconnected:
			// Clean up connection and exit
			return c.Close()

		default:
			panic("invalid state of tcp connection, bailing out")
		}

		time.Sleep(time.Millisecond * 10) // no need to check in constantly
	}
}
