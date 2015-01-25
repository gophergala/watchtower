// Package http contains all the
// HTTP handlers for Watchtower
package http

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RegisterHandler handles registration of new users (senders)
func RegisterHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response, statusCode, err := registerStreamHandler(w, r, params)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.WriteHeader(statusCode)
	fmt.Fprintln(w, response)
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

// JoinChannelsStreamHandler joins a channel (opening a HTTP stream)
func JoinChannelsStreamHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_, statusCode, err := joinChannelStreamHandler(w, r, params)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
}

// JoinChannelsAsyncHandler joins a channel (with URL callbacks)
func JoinChannelsAsyncHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response, statusCode, err := listChannelsHandler(r, params)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.WriteHeader(statusCode)
	fmt.Fprintln(w, response)
}

// BroadcastHandler broadcasts a message across one or more channels
func BroadcastHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response, statusCode, err := listChannelsHandler(r, params)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.WriteHeader(statusCode)
	fmt.Fprintln(w, response)
}

// SendMessageHandler sends a message to specifical recipients in a channel
func SendMessageHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	/*response, statusCode, err := listChannelsHandler(r, params)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.WriteHeader(statusCode)
	fmt.Fprintln(w, response)*/
	w.WriteHeader(http.StatusNotImplemented)
}
