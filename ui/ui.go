package ui

import (
	"fmt"
	"net/http"

	auth "github.com/abbot/go-http-auth"
)

type ui struct {
}

func New() *ui {
	return &ui{}
}

func (u *ui) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	info := auth.FromContext(r.Context())
	fmt.Fprintf(w, "<html><body><h1>Successfully Authenticated %s!</h1></body></html>", info.Username)
}
