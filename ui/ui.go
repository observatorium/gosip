package ui

import (
	"fmt"
	"net/http"

	"github.com/abbot/go-http-auth"
)

type UI struct{}

func New() *UI {
	return &UI{}
}

func (u *UI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	info := auth.FromContext(r.Context())
	fmt.Fprintf(w, "<html><body><h1>Successfully Authenticated %s!</h1></body></html>", info.Username)
}
