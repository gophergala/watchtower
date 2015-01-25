package users

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/gophergala/watchtower/messages"
)

var (
	// ErrFlushingToResponseWriter is thrown if the response writer can't be cast to a HTTP flusher
	ErrFlushingToResponseWriter = errors.New("error flushing to response writer")
)

type User interface {
	ID() uint32
	Send(messages.Message) error
	setID(uint32)
}

func NewHTTPStreamUser(w http.ResponseWriter) User {
	return &httpStreamUser{
		w: w,
	}
}

type httpStreamUser struct {
	id uint32
	w  http.ResponseWriter
}

func (h *httpStreamUser) ID() uint32 {
	return h.id
}

func (h *httpStreamUser) Send(m messages.Message) error {
	h.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(h.w, m.JSON())
	if f, ok := h.w.(http.Flusher); ok {
		f.Flush()
	} else {
		return ErrFlushingToResponseWriter
	}

	return nil
}

func (h *httpStreamUser) setID(id uint32) {
	h.id = id
}

func NewHTTPAsyncUser(callbackURL string) User {
	return &httpAsyncUser{
		callbackUrl: callbackURL,
	}
}

type httpAsyncUser struct {
	id          uint32
	callbackUrl string
}

func (h *httpAsyncUser) ID() uint32 {
	return h.id
}

func (h *httpAsyncUser) Send(m messages.Message) error {
	req, err := http.NewRequest("POST", h.callbackUrl, bytes.NewBuffer([]byte(m.JSON())))
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
