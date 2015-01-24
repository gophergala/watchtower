package channels

type Message struct {
	content   string
	broadcast bool
	receivers map[uint32]struct{}
}
