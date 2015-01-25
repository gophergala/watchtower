package tcp

import (
	"github.com/gophergala/watchtower/users"
	"testing"
)

var (
	_ users.User = (*user)(nil)
)

func TestUser(t *testing.T) {
	// TODO: Test Send()
	u := newUser()

	if u.ID() != 0 {
		t.Error("invalid id")
	}

	u.SetID(1)

	if u.ID() != 1 {
		t.Error("invalid id after set")
	}
}
