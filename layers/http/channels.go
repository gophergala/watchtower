package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gophergala/watchtower/channels"
	"github.com/gophergala/watchtower/users"
	"github.com/julienschmidt/httprouter"
)

var (
	// ErrInvalidSender is thrown if the sender
	// variable is missing or does not exist
	ErrInvalidSender = errors.New("invalid or missing sender")
	// ErrInvalidChannel is thrown if the channel
	// variable is missing or invalid
	ErrInvalidChannel = errors.New("invalid or missing channel")
)

func listChannelsHandler(r *http.Request, params httprouter.Params) (string, int, error) {
	// Grab user ID (error if not exists)
	sender64, err := strconv.ParseUint(r.FormValue("sender"), 0, 0)
	sender := uint32(sender64)
	if err != nil {
		log.Printf("%v", err)
		return "", http.StatusUnauthorized, ErrInvalidSender
	}

	// Check that user is registered (error if not)
	users := users.List()
	_, userRegistered := users[sender]
	if !userRegistered {
		return "", http.StatusForbidden, ErrInvalidSender
	}

	// Grab the list of channels
	list := channels.List()
	response := make(map[string]interface{})
	response["channels"] = list

	// Encode the response as JSON
	encoded, err := json.Marshal(response)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	// Return list of active channels
	return string(encoded), http.StatusOK, nil
}

// joinChannelsStreamHandler joins a channel (opening a HTTP stream)
func joinChannelStreamHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) (string, int, error) {
	// Grab user ID (error if not exists)
	sender64, err := strconv.ParseUint(r.FormValue("sender"), 0, 0)
	sender := uint32(sender64)
	if err != nil {
		log.Printf("%v", err)
		return "", http.StatusUnauthorized, ErrInvalidSender
	}

	// Check that user is registered (error if not)
	userList := users.List()
	_, userRegistered := userList[sender]
	if !userRegistered {
		return "", http.StatusForbidden, ErrInvalidSender
	}

	// Check that we got a valid channel ID
	channel64, err := strconv.ParseUint(params.ByName("id"), 0, 0)
	channelID := uint32(channel64)
	if err != nil {
		return "", http.StatusUnauthorized, ErrInvalidChannel
	}

	log.Printf("wut")
	// Gives the user a "response path" by tying the ResponseWriter to the channel
	err = users.JoinStreamingChannel(sender, channelID, w)
	if err != nil {
		return "", http.StatusBadRequest, err
	}

	log.Printf("but")
	// Join the channel
	channels.Join(sender, channelID)
	return "{\"status\": \"OK\"}", http.StatusOK, nil
}
