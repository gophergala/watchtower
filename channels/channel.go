package channels

import (
	"github.com/gophergala/watchtower/messages"
)

// A Channel has one or more subscribers
// and a queue of messages which should
// hopefully be empty most of the time
type Channel struct {
	id           uint32
	subscribers  map[uint32]struct{}
	messageQueue chan messages.Message
}
