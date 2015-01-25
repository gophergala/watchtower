package users

import (
	"testing"
)

var (
	_ User = (*httpStreamUser)(nil)
	_ User = (*httpAsyncUser)(nil)
)

func TestHTTPStreamUser(t *testing.T) {
	u := NewHTTPStreamUser()
	if u == nil {
		t.Error("failed to create http stream user")
	}

	if u.ID() != 0 {
		t.Error("invalid id")
	}

	u.setID(1)

	if u.ID() != 1 {
		t.Error("invalid id after set")
	}

	// test Send by passing in something that implements Flusher?
}

func TestHTTPAsyncUser(t *testing.T) {
	u := NewHTTPAsyncUser("www.google.com")
	if u == nil {
		t.Error("failed to create http async user")
	}

	if u.ID() != 0 {
		t.Error("invalid id")
	}

	u.setID(1)

	if u.ID() != 1 {
		t.Error("invalid id after set")
	}
	// Start server and test by sending to local URL?
}
