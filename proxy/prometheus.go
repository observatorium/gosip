package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/Go-SIP/gosip/context"
)

type TenantPrometheus interface {
	PrometheusURL(username string) (*url.URL, error)
}

func NewPrometheus(tenants TenantPrometheus) http.HandlerFunc {
	director := func(r *http.Request) {
		prometheusURL, err := context.PrometheusURLFromContext(r.Context())
		if err != nil {
			return
		}

		r.URL.Scheme = prometheusURL.Scheme
		r.URL.Host = prometheusURL.Host
		//normalize request so Prometheus doesn't know that request was proxied
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/prometheus")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		username, err := context.UsernameFromContext(r.Context())
		if err != nil {
			http.Error(w, `{"message": "failed to identify username"}`, http.StatusUnauthorized)
			return
		}

		prometheusURL, err := tenants.PrometheusURL(username)
		if err != nil {
			http.Error(w, `{"message": "failed to get Prometheus URL for tenant"}`, http.StatusBadRequest)
			return
		}

		proxy := &httputil.ReverseProxy{
			Director: director,
		}
		proxy.ServeHTTP(w, r.WithContext(context.WithPrometheusURL(r.Context(), prometheusURL)))
	}
}
