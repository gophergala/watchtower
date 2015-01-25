package channels

import (
	"errors"
	"sync"

	"github.com/gophergala/watchtower/messages"
	"github.com/gophergala/watchtower/users"
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

// Join a channel. Creates the channel if it doesn't exist
func Join(user users.User, channelID uint32) {
	channelEditMutex.Lock()

	// Add the new subscriber to an existing channel
	channel, exists := channels[channelID]
	if exists {
		channel.subscribers[user.ID()] = user
		channels[channelID] = channel
		return
	}

	// If the channel doesn't exist, create a new
	// one with the caller as the only subscriber
	c := Channel{
		id:           channelID,
		subscribers:  map[uint32]users.User{user.ID(): user},
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

// Send sends a message on a channel. The message can be either
// a BroadcastMessage (in which case it will be broadcasted)
// or a PrivateMessage (in which case it will only be sent to
// it's recipients)
func Send(m messages.Message, channelID uint32) error {
	channelEditMutex.RLock()
	defer channelEditMutex.RUnlock()

	// Check that the channel exists
	channel, exists := channels[channelID]
	if !exists {
		return ErrChannelDoesNotExist
	}

	// Queue the message for sending
	channel.messageQueue <- m
	return nil
}
