package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gophergala/watchtower/users"
	"github.com/julienschmidt/httprouter"
)

var (
	// ErrInvalidSecretKey is thrown if a user tries to register without or with a bad secret key
	ErrInvalidSecretKey = errors.New("invalid secret key")
)

func registerHandler(r *http.Request, params httprouter.Params) (string, int, error) {
	// TODO: Check secret key
	if false {
		return "", http.StatusUnauthorized, ErrInvalidSecretKey
	}

	id, err := users.Register()
	if err != nil {
		return "", http.StatusInternalServerError, users.ErrRegisteringNewUser
	}

	return fmt.Sprintf("{\"id\": %d", id), http.StatusOK, nil
}
