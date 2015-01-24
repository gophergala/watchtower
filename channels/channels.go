package channels

import (
	"sync"
)

const (
	channelMessageBufferSize = 100
)

var (
	channels         = make(map[uint32]Channel)
	channelEditMutex = &sync.RWMutex{}
)

func init() {
	// TODO: Start a goroutine that continously goes through
	// existing channels and sends out the messages
}

// Join a channel. Creates the channel if it doesn't exist
func Join(userID, channelID uint32) {
	channelEditMutex.Lock()

	// Add the new subscriber to an existing channel
	channel, exists := channels[channelID]
	if exists {
		channel.subscribers[userID] = struct{}{}
		channels[channelID] = channel
		return
	}

	// If the channel doesn't exist, create a new
	// one with the caller as the only subscriber
	c := Channel{
		id:           channelID,
		subscribers:  map[uint32]struct{}{userID: struct{}{}},
		messageQueue: make(chan Message, channelMessageBufferSize),
	}
	channels[channelID] = c

	channelEditMutex.Unlock()
}

// List returns a thread-safe list of active channels
func List() []uint32 {
	var list []uint32

	channelEditMutex.RLock()
	for channelID := range channels {
		list = append(list, channelID)
	}
	channelEditMutex.RUnlock()

	return list

}
