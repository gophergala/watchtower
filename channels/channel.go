package channels

type Channel struct {
	id           uint32
	subscribers  map[uint32]struct{}
	messageQueue chan Message
}
