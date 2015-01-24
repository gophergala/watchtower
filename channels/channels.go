package channels

import (
	"errors"
	"sync"

	"github.com/gophergala/watchtower/messages"
)

const (
	channelMessageBufferSize = 100
)

var (
	channels         = make(map[uint32]*Channel)
	channelEditMutex = &sync.RWMutex{}

	// ErrChannelDoesNotExist is thrown if a user tries to
	// send a message to a channel that does not exist
	ErrChannelDoesNotExist = errors.New("the specified channel does not exist")
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
		messageQueue: make(chan messages.Message, channelMessageBufferSize),
	}
	channels[channelID] = &c

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

// Broadcast broadcasts a message on a channel
func Broadcast(sender, channelID uint32, content string) error {
	channelEditMutex.RLock()
	defer channelEditMutex.RUnlock()

	// Check that the channel exists
	channel, exists := channels[channelID]
	if !exists {
		return ErrChannelDoesNotExist
	}

	// Queue the message
	channel.messageQueue <- messages.NewBroadcastMessage(sender, content)
	return nil
}
