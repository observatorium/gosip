package context

import (
	"context"
	"fmt"
	"net/url"
)

// key is used for private context keys.
type key int

const (
	prometheusURLKey key = iota
	usernameKey
)

// ErrContextMissing is returned when a context is missing a value.
type ErrContextMissing struct {
	// Value is the value that is missing from the context.
	Value string
}

// Error implements the error interface.
func (e *ErrContextMissing) Error() string {
	return fmt.Sprintf("context missing %s", e.Value)
}

// WithPrometheusURL returns a copy of the given context with a Prometheus URL stored.
func WithPrometheusURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, prometheusURLKey, u)
}

// PrometheusURLFromContext returns a Prometheus URL from the context.
func PrometheusURLFromContext(ctx context.Context) (*url.URL, error) {
	u, ok := ctx.Value(prometheusURLKey).(*url.URL)
	if !ok {
		return nil, &ErrContextMissing{"Prometheus URL"}
	}
	return u, nil
}

// WithUsername returns a copy of the given context with a username stored.
func WithUsername(ctx context.Context, u string) context.Context {
	return context.WithValue(ctx, usernameKey, u)
}

// UsernameFromContext returns a username from the context.
func UsernameFromContext(ctx context.Context) (string, error) {
	u, ok := ctx.Value(usernameKey).(string)
	if !ok {
		return "", &ErrContextMissing{"username"}
	}
	return u, nil
}
