package ui

import (
	"fmt"
	"net/http"

	auth "github.com/abbot/go-http-auth"
)

// UI serves Go-SIP user interface.
type UI struct{}

// New creates and returns a new UI instance.
func New() *UI {
	return &UI{}
}

// ServeHTTP implements the http.Handler interface.
func (u *UI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	info := auth.FromContext(r.Context())
	if info == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "<html><body><h1>Successfully Authenticated %s!</h1></body></html>", info.Username)
}
