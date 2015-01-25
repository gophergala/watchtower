package messages

import (
	"encoding/json"
)

// A Message is sent on a channel, either as
// a broadcast or directly to one or more users
// on the channel
type Message interface {
	Sender() uint32
	Content() string
	Receivers() map[uint32]struct{} // nil if broadcast
	JSON() string
}

// A BroadcastMessage is broadcasted to one or more channels
type BroadcastMessage struct {
	sender  uint32 `json:"id"`
	content string `json:"content"`
}

func NewBroadcastMessage(sender uint32, content string) *BroadcastMessage {
	return &BroadcastMessage{
		sender:  sender,
		content: content,
	}
}

// Sender returns the message's sender ID
func (b *BroadcastMessage) Sender() uint32 {
	return b.sender
}

// Content returns the message content
func (b *BroadcastMessage) Content() string {
	return b.content
}

// JSON returns a JSON-encoded version of the message
func (b *BroadcastMessage) JSON() string {
	encoded, _ := json.Marshal(b)
	return string(encoded)
}

// Receivers returns nil for a Broadcast message
func (b *BroadcastMessage) Receivers() map[uint32]struct{} {
	return nil
}

// A PrivateMessage is sent to one or more subscribers in a channel
type PrivateMessage struct {
	sender    uint32              `json:"sender"`
	content   string              `json:"content"`
	receivers map[uint32]struct{} `json:"-"`
}

// Sender returns the message's sender ID
func (p *PrivateMessage) Sender() uint32 {
	return p.sender
}

// Content returns the message content
func (p *PrivateMessage) Content() string {
	return p.content
}

// Receivers returns a map of the receiver user IDs
func (p *PrivateMessage) Receivers() map[uint32]struct{} {
	return p.receivers
}

// JSON returns a JSON-encoded version of the message
func (p *PrivateMessage) JSON() string {
	encoded, _ := json.Marshal(p)
	return string(encoded)
}
