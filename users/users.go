package users

import (
	"errors"
	"math/rand"
	"sync"
)

var (
	users      = make(map[uint32]struct{})
	usersMutex = &sync.RWMutex{}

	// ErrRegisteringNewUser if a new user can't be registered
	ErrRegisteringNewUser = errors.New("error registering new user")
)

// Register registers and stores a new user, returning his ID
func Register() (uint32, error) {
	var userID uint32

	usersMutex.Lock()
	for i := 0; i < 1000; i++ {
		// Try 1000 times to find a random user id, not already in users
		userID = rand.Uint32()
		_, alreadyExists := users[r]
		if userID == 0 || !alreadyExists { // 0 is not a valid user ID
			users[userID] = struct{}{}
		}
		break
	}

	if userID != 0 {
		return userID, nil
	}

	return userID, ErrRegisteringNewUser
}

// List returns a thread-safe copy of the current user list
func List() map[uint32]struct{} {
	// Return a copy of users
	usersCopy := make(map[uint32]struct{})

	usersMutex.RLock()
	for user := range users {
		usersCopy[user] = struct{}{}
	}
	usersMutex.RUnlock()

	return usersCopy

}
