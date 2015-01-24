package messages

import (
	"testing"
)

var (
	_ Message = (*BroadcastMessage)(nil)
	_ Message = (*PrivateMessage)(nil)
)

func TestPrivateMessage(t *testing.T) {
	p := PrivateMessage{
		sender:  0,
		content: "hello",
		receivers: map[uint32]struct{}{
			1: struct{}{},
		},
	}

	if p.sender != 0 || p.Sender() != 0 {
		t.Error("sender or Sender() did not match")
	}

	if len(p.Receivers()) != 1 || len(p.receivers) != 1 {
		t.Error("receivers or Receivers() did not match")
	}

	if p.content != "hello" || p.Content() != "hello" {
		t.Error("content or Content() did not match")
	}
}

func TestBroadcastMessage(t *testing.T) {
	b := BroadcastMessage{
		sender:  0,
		content: "hello",
	}

	if b.sender != 0 || b.Sender() != 0 {
		t.Error("sender or Sender() did not match")
	}

	if b.Receivers() != nil {
		t.Error("receivers or Receivers() did not match")
	}

	if b.content != "hello" || b.Content() != "hello" {
		t.Error("content or Content() did not match")
	}
}
