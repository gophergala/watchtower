package channels

// A Message is sent on a channel, either as
// a broadcast or directly to one or more users
// on the channel
type Message interface {
	Content() string
	Receivers() map[uint32]struct{} // nil if broadcast
}

// A BroadcastMessage is broadcasted to one or more channels
type BroadcastMessage struct {
	content string
}

func (b *BroadcastMessage) Content() string {
	return b.content
}

func (b *BroadcastMessage) Receivers() map[uint32]struct{} {
	return nil
}

// A PrivateMessage is sent to one or more subscribers in a channel
type PrivateMessage struct {
	content   string
	receivers map[uint32]struct{}
}

func (p *PrivateMessage) Content() string {
	return p.content
}

func (p *PrivateMessage) Receivers() map[uint32]struct{} {
	return p.receivers
}
