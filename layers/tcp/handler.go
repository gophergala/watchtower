package tcp

import (
	"net"
	"sync"
	"time"
)

type state int

const (
	stateNotConnected  state = 0
	stateRegistered    state = 1
	stateJoinedChannel state = 2
	stateDisconnected  state = 3

	registerMessageType    = byte('R')
	joinChannelMessageType = byte('J')
	sendMessageType        = byte('M')
)

var (
	bufferPool = sync.Pool{
		New: func() interface{} { return make([]byte, 1024) },
	}
)

func Handle(c net.Conn) error {
	var s state
	for {
		buffer := bufferPool.Get().([]byte) // Borrow a buffer from the pool
		// Read the incoming connection into the buffer.
		_, err := c.Read(buffer)
		if err != nil {
			return err
		}

		switch s {
		case stateNotConnected:
			// Waiting for the sender to register
			if buffer[0] == registerMessageType {
				// Register the user
				// Parse the message

				// Reply with an ack
				c.Write([]byte(""))

				// Join channel
				s = stateJoinedChannel
			}

		case stateRegistered:
			// Waiting for the user to join a channel
			if buffer[0] == joinChannelMessageType {
				// Join the channel

				// Reply with an ack
				c.Write([]byte(""))

				s = stateJoinedChannel
			}

			time.Sleep(time.Second)

		case stateJoinedChannel:
			// Waiting for sender to send or receive messages (or join more channels)
			if buffer[0] == joinChannelMessageType {
				// Join the channel

				// Reply with an ack
				c.Write([]byte(""))
			}

			if buffer[0] == sendMessageType {
				// Send message to channel
			}

			time.Sleep(time.Second)

		case stateDisconnected:
			// Clean up connection and exit
			return c.Close()

		default:
			panic("invalid state of tcp connection, bailing out")
		}

		bufferPool.Put(buffer)            // Put the buffer back in the pool
		time.Sleep(time.Millisecond * 10) // no need to check in constantly
	}
}
