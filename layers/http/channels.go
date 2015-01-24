package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func listChannelsHandler(r *http.Request, params httprouter.Params) (string, int, error) {
	// Grab user ID (error if not exists)
	// Check that user is registered (error if not)
	// Return list of active channels
	return "", http.StatusNotImplemented, nil
}
