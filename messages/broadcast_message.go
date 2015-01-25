package messages

import (
	"encoding/binary"
	"encoding/json"
)

// A BroadcastMessage is broadcasted to one or more channels
type BroadcastMessage struct {
	sender  uint32 `json:"id"`
	content string `json:"content"`
}

// NewBroadcastMessage creates a new broadcast message
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
func (b *BroadcastMessage) JSON(channelID uint32) string {
	m := make(map[string]interface{})
	m["sender"] = b.sender
	m["content"] = b.content
	m["channel"] = channelID
	encoded, _ := json.Marshal(m)

	return string(encoded)
}

// Bytes returns a version of the message fit
// for sending over a TCP/IP or UDP pipe. The
// format is defined in the documentation
func (b *BroadcastMessage) Bytes(channelID uint32) []byte {
	bytes := make([]byte, (messageHeaderSize + len([]byte(b.content))))

	binary.LittleEndian.PutUint32(bytes[0:4], (b.sender))
	binary.LittleEndian.PutUint32(bytes[4:8], channelID)
	binary.LittleEndian.PutUint16(bytes[8:10], 0) // Message type - reserved for later use
	binary.LittleEndian.PutUint32(bytes[10:14], uint32(len(b.content)))
	copy(bytes[14:], []byte(b.content))
	return bytes
}

// Receivers returns nil for a Broadcast message
func (b *BroadcastMessage) Receivers() map[uint32]struct{} {
	return nil
}
