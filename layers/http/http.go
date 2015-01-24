// package http contains all the
// HTTP handlers for Watchtower
package http

import (
	"net/http"

	"github.com/gophergala/watchtower/users"
)

var (
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

// RegisterHandler handles registration of new users (senders)
func RegisterHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response, statusCode, err := registerHandler(r, params)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.WriteHeader(statusCode)
	fmt.Fprintln(w, response)
}

func listChannelsHandler(r *http.Request, params httprouter.Params) (string, int, error) {
	// Grab user ID (error if not exists)
	// Check that user is registered (error if not)
	// Return list of active channels
	return "", http.StatusNotImplemented, nil
}

// ListChannelsHandler returns a list of active channels
func ListChannelsHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response, statusCode, err := listChannelsHandler(r, params)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.WriteHeader(statusCode)
	fmt.Fprintln(w, response)
}
