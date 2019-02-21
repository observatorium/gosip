package proxy

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/Go-SIP/gosip/auth"
)

type TenantPrometheus interface {
	PrometheusURL(username string) (*url.URL, error)
}

func NewPrometheus(tenants TenantPrometheus) http.HandlerFunc {
	type promURL string
	var promKey promURL

	director := func(r *http.Request) {
		prometheusURL := r.Context().Value(promKey).(*url.URL)

		r.URL.Scheme = prometheusURL.Scheme
		r.URL.Host = prometheusURL.Host
		//normalize request so Prometheus doesn't know that request was proxied
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/prometheus")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		username := auth.Username(r.Context())
		prometheusURL, err := tenants.PrometheusURL(username)
		if err != nil {

			http.Error(w, `{"message": "failed to get Prometheus URL for tenant"}`, http.StatusBadRequest)
			return
		}

		// Put the upstream PrometheusURL into the request's context for the director
		r = r.WithContext(context.WithValue(r.Context(), promKey, prometheusURL))

		proxy := &httputil.ReverseProxy{
			Director: director,
		}
		proxy.ServeHTTP(w, r)
	}
}
