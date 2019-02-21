package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/Go-SIP/gosip/auth"
	"github.com/Go-SIP/gosip/config"
	"github.com/Go-SIP/gosip/tenant"
	"github.com/Go-SIP/gosip/ui"
	"github.com/Go-SIP/gosip/users"
)

func main() {
	os.Exit(Main())
}

func Main() int {
	c, err := config.LoadFile("config.yaml")
	if err != nil {
		fmt.Println("Failed to load config:", err)
		return 1
	}

	tenantsIndex, err := config.Tenants(c)
	if err != nil {
		fmt.Println("failed to parse tenants:", err)
		return 1
	}
	tenants := tenant.NewStatic(tenantsIndex)

	users := users.NewStaticUsersDatabase(c.Users)
	auth := auth.NewHandler(users)

	mux := http.NewServeMux()
	mux.Handle("/prometheus/", auth.Token(NewPrometheusReverseProxy(tenants)))
	//mux.Handle("/jaeger", auth.Token(httputil.NewSingleHostReverseProxy(jaegerURL)))
	mux.Handle("/", auth.Basic(ui.New()))

	fmt.Println("Running server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Failed to run server:", err)
		return 2
	}

	return 0
}
