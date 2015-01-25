package users

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gophergala/watchtower/messages"
)

// NewHTTPStreamUser creates a user who will get
// his messages sent as part of a streaming response
// TODO: Keep-alive
// TODO: Clean up by closing the streams
func NewHTTPStreamUser() User {
	return &httpStreamUser{
		channelStreams: make(map[uint32]http.ResponseWriter),
	}
}

type httpStreamUser struct {
	id             uint32
	channelStreams map[uint32]http.ResponseWriter
}

func (h *httpStreamUser) ID() uint32 {
	return h.id
}

func (h *httpStreamUser) Send(m messages.Message, channelID uint32) error {
	// TODO: Deal with race conditions when accessing the map from several threads
	h.channelStreams[channelID].Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(h.channelStreams[channelID], m.JSON(channelID))
	if f, ok := h.channelStreams[channelID].(http.Flusher); ok {
		f.Flush()
	} else {
		return ErrFlushingToResponseWriter
	}

	return nil
}

func (h *httpStreamUser) setID(id uint32) {
	h.id = id
}

// NewHTTPAsyncUser creates a user who will get
// his messages sent to a specified URL endpoint
func NewHTTPAsyncUser(url string) User {
	return &httpAsyncUser{
		callbackURL: url,
	}
}

type httpAsyncUser struct {
	id          uint32
	callbackURL string
}

func (h *httpAsyncUser) ID() uint32 {
	return h.id
}

func (h *httpAsyncUser) Send(m messages.Message, channelID uint32) error {
	req, err := http.NewRequest("POST", h.callbackURL, bytes.NewBuffer([]byte(m.JSON(channelID))))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Server-Name", "Watchtower")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func (h *httpAsyncUser) setID(id uint32) {
	h.id = id
}
