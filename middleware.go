package main

import (
	"net/http"
)

// Middleware is s convenience type for constructing
// stacks of middleware functions. E.g.:
// m := Middleware{auth, logger, tracing}
type Middleware []func(http.Handler) http.Handler

// Wrap adds the base handler that should
// leverage the middleware stack. E.g.:
// mux.Handle("/api/user", m.Wrap(userHandler))
func (m Middleware) Wrap(h http.Handler) http.Handler {
	if len(m) == 0 {
		return h
	}
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}
	return h
}
