package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gophergala/watchtower/config"
	"github.com/gophergala/watchtower/users"
	"github.com/julienschmidt/httprouter"
)

var (
	// ErrInvalidSecretKey is thrown if a user tries to register without or with a bad secret key
	ErrInvalidSecretKey = errors.New("invalid or missing secret key")
)

func registerStreamHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) (string, int, error) {
	// Decode the JSON request
	type Request struct {
		Secret string `json:"secret_key"`
	}
	var req Request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		return "", http.StatusUnauthorized, ErrInvalidSecretKey
	}

	// Check that we got the correct secret key
	if req.Secret != config.Secret() {
		return "", http.StatusUnauthorized, ErrInvalidSecretKey
	}

	// Register a new user
	user := users.NewHTTPStreamUser(w)
	id, err := users.Register(user)
	if err != nil {
		return "", http.StatusInternalServerError, users.ErrRegisteringNewUser
	}

	// return the user's new ID
	return fmt.Sprintf("{\"id\": %d", id), http.StatusOK, nil
}
