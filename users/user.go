package users

import (
	"errors"

	"github.com/gophergala/watchtower/messages"
)

var (
	// ErrFlushingToResponseWriter is thrown if the response writer can't be cast to a HTTP flusher
	ErrFlushingToResponseWriter = errors.New("error flushing to response writer")
)

// A User represents a user that is currently connected
// and listening to one or more channels. The main thing
// tying users together is their ability to receive messages
// via the Send method
type User interface {
	ID() uint32
	Send(messages.Message) error
	setID(uint32)
}
